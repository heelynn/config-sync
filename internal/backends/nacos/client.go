package nacos

import (
	"config-sync/internal/properties"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"strconv"
	"strings"
)

// ConfigClientMap 存储配置客户端
var ConfigClientMap map[string]*config_client.IConfigClient = make(map[string]*config_client.IConfigClient)

func initClients() {
	initConfigClientByProperties()
}

// initConfigClientByProperties 初始化配置客户端
func initConfigClientByProperties() {
	if properties.Get().NacosMap == nil || len(properties.Get().NacosMap) == 0 {
		return
	}
	for _, nacosProperties := range properties.Get().NacosMap {
		err := initConfigClient(nacosProperties)
		if err != nil {
			panic(err)
		}
	}
}

func initConfigClient(nacosProperties *properties.NacosProperties) error {
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
		return err
	}
	ConfigClientMap[nacosProperties.Id] = &configClient
	return nil
}
