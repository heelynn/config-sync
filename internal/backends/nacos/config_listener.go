package nacos

import (
	"config-sync/internal/properties"
	"config-sync/pkg/utils/file_util"
	"config-sync/pkg/utils/os_util"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func initConfigListener() {
	for id, configClient := range ConfigClientMap {
		nacosProperties, err := properties.Get().GetNacosById(id)
		if err != nil {
			fmt.Println("Get nacos properties error:", err)
			continue
		}
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
			// Get nacos properties
			nacosProperties, err := properties.Get().GetNacosById(nacosPropertiesId)
			if err != nil {
				fmt.Println(err)
			}
			// Write to file
			err = file_util.WriteToFile(file_util.GetFileName(nacosProperties.FilePath, dataId), data)
			if err != nil {
				fmt.Println(err)
			}
			// Execute command
			if nacosProperties.Command != "" {
				command, err := os_util.ExecuteCommand(nacosProperties.Command)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(command)
			}
		},
	})

	if err != nil {
		panic(err)
	}
}
