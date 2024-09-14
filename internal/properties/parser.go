package properties

// PropertiesMap is a struct that contains all properties
type PropertiesMap struct {
	NacosMap map[string]*NacosConfig
}

type Parser interface {
	Parse(filePath string) *Properties
	GetParser() *Parser
	GetFilePath() string
}

func checkPropertiesIdDuplicate(properties *Properties) {
	if properties != nil {

		// Check duplicate id for config
		if properties.Config != nil {
			if properties.Config.Nacos != nil && len(properties.Config.Nacos) > 0 {
				idMap := make(map[string]bool)
				for _, nacos := range properties.Config.Nacos {
					if _, ok := idMap[nacos.Id]; ok {
						panic("Duplicate id: " + nacos.Id)
					}
					idMap[nacos.Id] = true
				}
			}
		}
	}
}
