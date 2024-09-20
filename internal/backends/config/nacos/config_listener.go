package nacos

import (
	"config-sync/pkg/utils/file_util"
	"config-sync/pkg/utils/os_util"
	"config-sync/pkg/zlog"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func initConfigListener() {
	for id, configClient := range configClientMap {
		nacosProperties := nacosConfigMap[id]

		for _, propertyName := range nacosProperties.PropertyNames {
			listenConfig(*configClient, propertyName, nacosProperties.Group, id)
		}

	}

}

func listenConfig(configClient config_client.IConfigClient, dataId, group string, nacosPropertiesId string) {
	err := configClient.ListenConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {
			zlog.Logger.Infof("Config changed, dataId: %s, group: %s", dataId, group)
			// Get nacos properties
			nacosProperties := nacosConfigMap[nacosPropertiesId]
			// Write to file
			fileName := file_util.GetFileName(nacosProperties.FilePath, dataId)
			err := file_util.WriteToFile(fileName, data)
			if err != nil {
				zlog.Logger.Error(err)
			}
			zlog.Logger.Infof("nacos config generated config written to file,fileName: %s", fileName)
			zlog.Logger.Debugf("[%s] write content  \n %s ", fileName, data)
			// Execute command
			if nacosProperties.Command != "" {
				command, err := os_util.ExecuteCommand(nacosProperties.Command)
				if err != nil {
					zlog.Logger.Error(err)
				}
				zlog.Logger.Infof("\n -- Command Result--\n %s\n ", command)
			}
		},
	})

	if err != nil {
		panic(err)
	}
}
