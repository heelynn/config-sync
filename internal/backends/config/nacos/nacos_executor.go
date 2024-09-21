package nacos

import (
	"config-sync/internal/backends/config"
	"config-sync/internal/properties"
	"config-sync/internal/sync"
	"config-sync/pkg/utils/file_util"
	"config-sync/pkg/zlog"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"strconv"
	"strings"
)

//var nacosConfigExecutors []NacosConfigExecutor

type NacosConfigExecutor struct {
	Id            string
	NacosClient   config_client.IConfigClient
	Group         string
	PropertyNames []string
	FilePath      string //同步到的文件目录，绝对路径/相对路径均可
	Command       string //发生变化时执行的命令
}

func (n *NacosConfigExecutor) RegisterChangedListener() error {
	if n.PropertyNames == nil || len(n.PropertyNames) < 1 {
		return fmt.Errorf("nacosConfigExecutor.PropertyNames is empty")
	}
	for _, propertyName := range n.PropertyNames {

		content, err := n.getConfig(propertyName)
		if err != nil {
			zlog.Logger.Warnf("can not get nacos [%s] config, dataId:%s, err:%v", n.Id, propertyName, err)
			// 忽略此配置
			continue
		} else {
			// Write to file , initialization
			zlog.Logger.Debugf("nacos [%s] config initialized, write to file, dataId:%s", n.Id, propertyName)
			err = sync.CheckFileChangedAndExecuteCommand(file_util.GetFileName(n.FilePath, propertyName), content, n.Command)
			if err != nil {
				return err
			}
			// init nacos config listener
			err = n.NacosClient.ListenConfig(vo.ConfigParam{
				DataId: propertyName,
				Group:  n.Group,
				OnChange: func(namespace, group, dataId, data string) {
					zlog.Logger.Infof("Config changed, dataId: %s, group: %s", dataId, group)
					// Write to file
					err = sync.CheckFileChangedAndExecuteCommand(file_util.GetFileName(n.FilePath, propertyName), content, n.Command)
					if err != nil {
						return
					}
					// Execute command
					if n.Command != "" {
						config.ExecuteCommand(n.Command)
					} else {
						zlog.Logger.Infof("nacos [%s] command is empty, dataId: %s, group: %s", n.Id, dataId, group)
					}
				},
			})

			if err != nil {
				panic(err)
			}

		}

	}
	return nil
}

// getConfig 获取Nacos配置
func (n *NacosConfigExecutor) getConfig(dataId string) (string, error) {
	if n.NacosClient == nil {
		return "", fmt.Errorf("nacosConfigExecutor.nacosClient is nil")
	}
	content, err := n.NacosClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  n.Group,
	})
	if err != nil {
		zlog.Logger.Warnf("fail to get config, dataId: %s, group: %s, err: %v", dataId, n.Group, err)
	}
	return content, err

}

// RegisterNacosConfigListener 注册Nacos配置监听器
func RegisterNacosConfigListener() {
	if properties.Prop == nil || properties.Prop.Config == nil || properties.Prop.Config.Nacos == nil || len(properties.Prop.Config.Nacos) < 1 {
		zlog.Logger.Warn("Nacos config not found")
	}
	for _, nacosConfig := range properties.Prop.Config.Nacos {
		var executor NacosConfigExecutor
		client, err := getNacosClient(nacosConfig)
		if err != nil {
			zlog.Logger.Errorf("fail to create nacos config client, err: %v", err)
		}
		executor.NacosClient = client
		executor.Group = nacosConfig.Group
		executor.PropertyNames = nacosConfig.PropertyNames
		executor.FilePath = nacosConfig.FilePath
		executor.Command = nacosConfig.Command
		executor.Id = nacosConfig.Id
		//nacosConfigExecutors = append(nacosConfigExecutors, executor)

		err = executor.RegisterChangedListener()
		if err != nil {
			zlog.Logger.Errorf("fail to register nacos config listener, err: %v", err)
		}
	}

}

// getNacosClient 创建Nacos客户端
func getNacosClient(nacosConfig *properties.NacosConfig) (config_client.IConfigClient, error) {
	if nacosConfig == nil {
		return nil, fmt.Errorf("nacos properties is nil")
	}
	// Nacos服务器地址
	var serverConfigs []constant.ServerConfig
	for _, server := range strings.Split(nacosConfig.ServerAddr, ",") {
		// Nacos服务器地址
		if server == "" || !strings.Contains(server, ":") {
			continue
		}
		ipPort := strings.Split(server, ":")
		port, err := strconv.ParseUint(ipPort[1], 10, 64)
		if err != nil {
			zlog.Logger.Warnf("fail to parse server port [%s]", ipPort[1])
			continue
		}
		serverConfigs = append(serverConfigs, constant.ServerConfig{
			IpAddr: ipPort[0],
			Port:   port,
		})
	}
	if len(serverConfigs) == 0 {
		return nil, fmt.Errorf("nacos server address is empty")
	}

	// 客户端配置
	clientConfig := constant.ClientConfig{
		NamespaceId:         nacosConfig.Namespace, // 如果不需要命名空间，可以留空
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogLevel:            "debug",
	}
	if nacosConfig.Username != "" && nacosConfig.Password != "" {
		clientConfig.Username = nacosConfig.Username
		clientConfig.Password = nacosConfig.Password
	}

	// 创建配置客户端
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		zlog.Logger.Errorf("fail to create nacos config client, err: %v", err)
		return nil, err
	}
	return configClient, nil
}
