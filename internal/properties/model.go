package properties

// Properties 根配置
type Properties struct {
	Config    *Config    `yaml:"config"`
	Discovery *Discovery `yaml:"discovery"`
}

// 来自配置中心的配置
type Config struct {
	Nacos []*NacosConfig `yaml:"nacos"`
}

// 来自注册中心的配置
type Discovery struct {
	Nacos []*NacosDiscovery `yaml:"nacos"`
}

// CheckId 检查配置的Id是否重复
func (y *Config) CheckId() {
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

type NacosConfig struct {
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
func (y *NacosConfig) check() {
	if y.Id == "" || y.ServerAddr == "" || y.Namespace == "" || y.Group == "" || len(y.PropertyNames) == 0 || y.FilePath == "" {
		panic("NacosConfig must have id, server_addr, namespace, group, property_names, file_path")
	}
}

type NacosDiscovery struct {
	Id              string   `yaml:"id"`
	ServerAddr      string   `yaml:"server_addr"`
	Namespace       string   `yaml:"namespace"`
	Username        string   `yaml:"username"`
	Password        string   `yaml:"password"`
	Group           string   `yaml:"group"`
	ServiceNames    []string `yaml:"service_names"`
	Template        string   `yaml:"tempalte"`
	RefreshInterval string   `yaml:"refresh_interval"`
	FilePath        string   `yaml:"file_path"` //同步到的文件目录，绝对路径
	Command         string   `yaml:"command"`
}

// generateId 设置id
func (y *NacosDiscovery) check() {
	if y.Id == "" || y.ServerAddr == "" || y.Namespace == "" || y.Group == "" || len(y.ServiceNames) == 0 || y.FilePath == "" || y.Template == "" || y.RefreshInterval == "" {
		panic("NacosConfig must have id, server_addr, namespace, group, property_names, file_path")
	}
}
