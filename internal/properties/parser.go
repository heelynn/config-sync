package properties

import "errors"

// PropertiesMap is a struct that contains all properties
type PropertiesMap struct {
	NacosMap map[string]*NacosProperties
}

func (*PropertiesMap) GetNacosById(id string) (*NacosProperties, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	return propertiesMap.NacosMap[id], nil
}

var propertiesMap *PropertiesMap

type Parser interface {
	Parse(filePath string) *Properties
	GetParser() *Parser
	GetFilePath() string
}

func putPropertiesId(properties *Properties) {
	if properties == nil {
		return
	}
	// Set id for nacos
	if properties.Nacos != nil && len(properties.Nacos) > 0 {
		for _, nacos := range properties.Nacos {
			nacos.generateId()
		}
	}
}

func checkPropertiesIdDuplicate(properties *Properties) {
	if properties != nil {
		// Check duplicate id for nacos
		if properties.Nacos != nil && len(properties.Nacos) > 0 {
			idMap := make(map[string]bool)
			for _, nacos := range properties.Nacos {
				if _, ok := idMap[nacos.Id]; ok {
					panic("Duplicate id: " + nacos.Id)
				}
				idMap[nacos.Id] = true
			}
		}
	}
}

// GetPropertiesMap returns the properties map
func buildPropertiesMap(parser Parser) {
	path := parser.GetFilePath()
	properties := parser.Parse(path)
	if properties == nil {
		panic("properties is nil")
	}
	// Set id for nacos ...
	putPropertiesId(properties)
	// Check duplicate id for nacos ...
	checkPropertiesIdDuplicate(properties)
	config := PropertiesMap{NacosMap: make(map[string]*NacosProperties)}
	// Set id for nacos
	if properties.Nacos != nil && len(properties.Nacos) > 0 {
		for _, nacos := range properties.Nacos {
			config.NacosMap[nacos.Id] = nacos
		}
	}
	propertiesMap = &config

}
