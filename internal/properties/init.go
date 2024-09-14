package properties

func InitProperties() {
	parser := newYamlParser()
	buildPropertiesMap(parser)
}

func Get() *PropertiesMap {
	return propertiesMap
}
