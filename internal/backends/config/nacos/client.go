package nacos

import (
	"config-sync/internal/properties"
	"config-sync/pkg/zlog"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"strconv"
	"strings"
)

// configClientMap 存储Nacos配置客户端
var configClientMap map[string]*config_client.IConfigClient = make(map[string]*config_client.IConfigClient)

// nacosConfigMap 存储Nacos配置
var nacosConfigMap map[string]*properties.NacosConfig = make(map[string]*properties.NacosConfig)

func initClients() {
	if properties.Prop == nil || properties.Prop.Config.Nacos == nil || len(properties.Prop.Config.Nacos) == 0 {
		zlog.Logger.Warnf("config.nacos config is empty, skip init config client")
		return
	}
	bindNacosConfigMap()
	initConfigClientByProperties()
}

// bindNacosConfigMap 绑定Nacos配置到map，id为key
func bindNacosConfigMap() {
	for _, nacosProperties := range properties.Prop.Config.Nacos {
		nacosConfigMap[nacosProperties.Id] = nacosProperties
	}
}

// initConfigClientByProperties 初始化配置客户端
func initConfigClientByProperties() {
	// if has no nacos config, return

	zlog.Logger.Infof("init nacos config clients")
	for _, nacosProperties := range nacosConfigMap {
		err := initConfigClient(nacosProperties)
		if err != nil {
			panic(err)
		}
	}
}

func initConfigClient(nacosProperties *properties.NacosConfig) error {
	if nacosProperties == nil {
		return nil
	}
	// Nacos服务器地址
	var serverConfigs []constant.ServerConfig
	for _, server := range strings.Split(nacosProperties.ServerAddr, ",") {
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
		return fmt.Errorf("nacos server address is empty")
	}

	// 客户端配置
	clientConfig := constant.ClientConfig{
		NamespaceId:         nacosProperties.Namespace, // 如果不需要命名空间，可以留空
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogLevel:            "debug",
	}
	if nacosProperties.Username != "" && nacosProperties.Password != "" {
		clientConfig.Username = nacosProperties.Username
		clientConfig.Password = nacosProperties.Password
	}

	// 创建配置客户端
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		zlog.Logger.Errorf("fail to create nacos config client, err: %v", err)
		return err
	}
	configClientMap[nacosProperties.Id] = &configClient
	return nil
}
