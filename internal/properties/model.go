package properties

type Properties struct {
	Nacos []*NacosProperties `yaml:"nacos"`
}

// CheckId 检查配置的Id是否重复
func (y *Properties) CheckId() {
	//检查Nacos配置的Id是否重复
	if y.Nacos != nil && len(y.Nacos) > 0 {
		seen := make(map[string]bool, len(y.Nacos))
		for _, nacos := range y.Nacos {
			if _, ok := seen[nacos.Id]; ok {
				panic("Duplicate Nacos Id: " + nacos.Id)
			}
			seen[nacos.Id] = true
		}
	}
}

type NacosProperties struct {
	Id            string   `yaml:"id"`
	ServerAddr    string   `yaml:"server_addr"`
	Namespace     string   `yaml:"namespace"`
	Username      string   `yaml:"username"`
	Password      string   `yaml:"password"`
	Group         string   `yaml:"group"`
	PropertyNames []string `yaml:"property_names"`
	FilePath      string   `yaml:"file_path"` //同步到的文件目录，绝对路径
	Command       string   `yaml:"command"`   //发生变化时执行的命令
}

// generateId 设置id
func (y *NacosProperties) generateId() {
	y.check()
	if y.Id == "" {
		y.Id = y.ServerAddr + "_" + y.Namespace + "_" + y.Group
	}
}

// generateId 设置id
func (y *NacosProperties) check() {
	if y.ServerAddr == "" || y.Namespace == "" || y.Group == "" || len(y.PropertyNames) == 0 || y.FilePath == "" {
		panic("Nacos Properties is invalid")
	}
}
