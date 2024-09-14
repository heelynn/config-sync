package properties

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func newYamlParser() *YamlParser {
	return &YamlParser{}
}

type YamlParser struct{}

func (parser *YamlParser) GetParser() *Parser {
	//TODO implement me
	var p Parser = parser
	return &p
}

func (parser *YamlParser) Parse(yamlFile string) *Properties {

	// 读取YAML文件内容
	yamlData, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		panic(err)
	}

	// 创建一个变量来接收解析后的数据
	var prop Properties

	// 解析YAML文件
	err = yaml.Unmarshal(yamlData, &prop)
	if err != nil {
		panic(err)
	}

	return &prop
}

func (parser *YamlParser) GetFilePath() string {
	return "application.yaml"
}
