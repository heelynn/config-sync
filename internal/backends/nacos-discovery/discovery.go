package nacos_discovery

import (
	"config-sync/internal/properties"
	"config-sync/pkg/zlog"
)

// nacosConfigMap 存储Nacos配置
var nacosConfigMap map[string]*properties.NacosDiscovery = make(map[string]*properties.NacosDiscovery)

func initClients() {
	if properties.Prop == nil || properties.Prop.Discovery.Nacos == nil || len(properties.Prop.Discovery.Nacos) == 0 {
		zlog.Logger.Warnf("discovery.nacos config is empty, skip init discovery client")
	}
	bindNacosDiscpveryMap()
}

// bindNacosDiscpveryMap 绑定Nacos配置到map，id为key
func bindNacosDiscpveryMap() {
	for _, nacosProperties := range properties.Prop.Discovery.Nacos {
		nacosConfigMap[nacosProperties.Id] = nacosProperties
	}
}
