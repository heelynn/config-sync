package nacos

import (
	"config-sync/internal/properties"
	"config-sync/pkg/zlog"
)

func init() {
	if properties.Prop == nil || properties.Prop.Discovery.Nacos == nil || len(properties.Prop.Discovery.Nacos) == 0 {
		zlog.Logger.Warnf("nacos discovery is disabled")
		return
	}
	for _, nacos := range properties.Prop.Discovery.Nacos {
		zlog.Logger.Infof("init nacos discovery %s", nacos.Id)
		executor, err := NewNacosExecutor(nacos)
		if err != nil {
			zlog.Logger.Errorf("init nacos discovery [%s] fail, %v", nacos.Id, err)
			continue
		}
		//err = executor.Execute()
		err = executor.TickerExecute()
		if err != nil {
			zlog.Logger.Errorf("execute nacos discovery [%s] fail, %v", nacos.Id, err)
		}
	}

}
