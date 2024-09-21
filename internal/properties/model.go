package properties

import (
	"config-sync/pkg/startup"
	"config-sync/pkg/utils/file_util"
	"path/filepath"
)

// Properties 根配置
type Properties struct {
	Config    *Config    `yaml:"config"`
	Discovery *Discovery `yaml:"discovery"`
	Log       *Log       `yaml:"log"`
}

// Check 检查配置是否正确
func (p *Properties) check() {
	// 检查配置是否正确
	if p.Config != nil {
		// 检查Nacos配置是否正确
		if p.Config.Nacos != nil && len(p.Config.Nacos) > 0 {
			for _, nacos := range p.Config.Nacos {
				nacos.check()
			}
		}
	}
	// 检查注册中心配置是否正确
	if p.Discovery != nil {
		//检查Nacos注册中心配置是否正确
		if p.Discovery.Nacos != nil && len(p.Discovery.Nacos) > 0 {
			for _, nacos := range p.Discovery.Nacos {
				nacos.check()
			}
		}
	}
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
	Template        string   `yaml:"template"`
	RefreshInterval int      `yaml:"refresh_interval"`
	FilePath        string   `yaml:"file_path"`   //同步到的文件目录，绝对路径
	FileSuffix      string   `yaml:"file_suffix"` //文件后缀名
	Command         string   `yaml:"command"`
}

// generateId 设置id
func (y *NacosDiscovery) check() {
	if y.Id == "" || y.ServerAddr == "" || y.Namespace == "" || y.Group == "" || len(y.ServiceNames) == 0 || y.FilePath == "" || y.Template == "" || y.RefreshInterval == 0 {
		panic("NacosConfig must have id, server_addr, namespace, group, property_names, file_path")
	}
	templatePath := filepath.Join(startup.RootConfigPath, string(filepath.Separator), y.Template)
	if ok, _ := file_util.FileExists(templatePath); !ok {
		panic("NacosDiscovery id [" + y.Id + "] template not exists: " + templatePath)
	}
}

// Log 定义了日志配置的结构体
type Log struct {
	Output     string `yaml:"output"`      // 输出方式，例如 "console" 或 "file"
	Level      string `yaml:"level"`       // 日志级别，例如 "info"
	Path       string `yaml:"path"`        // 日志文件路径
	MaxSize    int    `yaml:"max-size"`    // 日志文件最大大小（MB）
	MaxAge     int    `yaml:"max-age"`     // 日志文件保留的最大天数
	MaxBackups int    `yaml:"max-backups"` // 日志文件的最大备份数量
}

// 默认日志配置
var defaultLogConfig = Log{
	Output:     "file",
	Level:      "info",
	Path:       "../logs/info.log",
	MaxSize:    100,
	MaxAge:     30,
	MaxBackups: 10,
}

// SetLogDefaultValues 为 Log 设置默认值
func SetLogDefaultValues(prop *Properties) {
	// 如果没有设置日志配置，则使用默认配置
	if prop.Log == nil {
		prop.Log = &defaultLogConfig
		return
	}

	// 设置默认值
	lc := prop.Log
	if lc.Output == "" {
		lc.Output = defaultLogConfig.Output
	}
	if lc.Level == "" {
		lc.Level = defaultLogConfig.Level
	}
	if lc.Path == "" {
		lc.Path = defaultLogConfig.Path
	}
	if lc.MaxSize == 0 {
		lc.MaxSize = defaultLogConfig.MaxSize
	}
	if lc.MaxAge == 0 {
		lc.MaxAge = defaultLogConfig.MaxAge
	}
	if lc.MaxBackups == 0 {
		lc.MaxBackups = defaultLogConfig.MaxBackups
	}
}
