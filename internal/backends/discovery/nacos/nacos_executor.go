package nacos

import (
	"config-sync/internal/backends/discovery"
	"config-sync/internal/properties"
	"config-sync/internal/sync"
	"config-sync/pkg/http/client"
	"config-sync/pkg/zlog"
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

type NacosExecutor struct {
	nacosDiscovery *properties.NacosDiscovery
}

func NewNacosExecutor(nacosDiscovery *properties.NacosDiscovery) (*NacosExecutor, error) {
	if nacosDiscovery == nil {
		return nil, fmt.Errorf("nacosDiscovery config is nil")
	}
	if nacosDiscovery.ServerAddr == "" || nacosDiscovery.Namespace == "" || nacosDiscovery.Template == "" ||
		nacosDiscovery.RefreshInterval == 0 || nacosDiscovery.FilePath == "" || nacosDiscovery.ServiceNames == nil ||
		len(nacosDiscovery.ServiceNames) == 0 {
		return nil, fmt.Errorf("nacosDiscovery config is invalid or missing some fields")
	}
	return &NacosExecutor{
		nacosDiscovery: nacosDiscovery,
	}, nil
}

func (n *NacosExecutor) Execute() error {
	// TODO: implement
	hosts := strings.Split(n.nacosDiscovery.ServerAddr, ",")
	httpClient := client.NewHttpClientHosts(hosts, client.GET, "/nacos/v1/ns/instance/list")
	httpClient.AddHeader("Content-Type", "application/x-www-form-urlencoded")
	httpClient.AddParam("namespaceId", n.nacosDiscovery.Namespace)
	httpClient.AddParam("groupName", n.nacosDiscovery.Group)
	httpClient.AddParam("healthyOnly", "true")
	for _, serviceName := range n.nacosDiscovery.ServiceNames {
		httpClient.AddParam("serviceName", serviceName)
		code, body, err := httpClient.DoInstances()
		if err != nil {
			return err
		}
		if code != 200 {
			zlog.Logger.Errorf("nacosDiscovery request failed with code %d , body %s", code, body)
			continue
		}
		// 结果解析
		var instance NacosInstance
		err = json.Unmarshal(body, &instance)
		if err != nil {
			zlog.Logger.Errorf("nacosDiscovery unmarshal failed with err %s", err)
		}
		// 写入文件
		var instances []discovery.InstanceResult = make([]discovery.InstanceResult, 0)
		if instance.Hosts == nil || len(instance.Hosts) == 0 {
			// if instance.Hosts is nil or empty, skip this instance
			zlog.Logger.Warnf("nacosDiscovery instance hosts is nil")
			return nil
		} else {
			for _, host := range instance.Hosts {
				instances = append(instances, discovery.InstanceResult{
					Host:   host.IP,
					Port:   host.Port,
					Weight: int(host.Weight + 0.5),
				})

			}
		}
		discoveryResults := discovery.DiscoveryResult{
			Name:      serviceName,
			Instances: instances,
		}
		// 生成文件名
		fileName := serviceName
		if n.nacosDiscovery.Template != "" {
			fileName = fileName + "." + n.nacosDiscovery.FileSuffix
		}
		path := filepath.Join(n.nacosDiscovery.FilePath, string(filepath.Separator), fileName)

		// 获取模板生成的内容
		content, err := discovery.GenerateTemplate(n.nacosDiscovery.Template, discoveryResults)
		if err != nil {
			return err
		}
		// 检查文件是否有变化，并执行命令
		err = sync.CheckFileChangedAndExecuteCommand(path, content, n.nacosDiscovery.Command)
		if err != nil {
			return err
		}

	}

	return nil

}

func (n *NacosExecutor) TickerExecute() error {
	// 创建一个Ticker，设置时间间隔
	ticker := time.NewTicker(time.Duration(n.nacosDiscovery.RefreshInterval) * time.Second)

	// 使用defer确保在函数结束时停止Ticker
	//defer ticker.Stop()

	err := n.Execute()
	if err != nil {
		return err
	}

	// 循环等待Ticker发送的时间值
	go func() {
		for {
			select {
			case <-ticker.C:
				// 每分钟执行一次的代码
				n.Execute()
			}
		}
	}()
	return nil

}

type NacosInstance struct {
	Name                     string     `json:"name"`
	GroupName                string     `json:"groupName"`
	Clusters                 string     `json:"clusters"`
	CacheMillis              int        `json:"cacheMillis"`
	Hosts                    []HostInfo `json:"hosts"`
	LastRefTime              int64      `json:"lastRefTime"`
	Checksum                 string     `json:"checksum"`
	AllIPs                   bool       `json:"allIPs"`
	ReachProtectionThreshold bool       `json:"reachProtectionThreshold"`
	Valid                    bool       `json:"valid"`
}

type HostInfo struct {
	InstanceId                string            `json:"instanceId"`
	IP                        string            `json:"ip"`
	Port                      int               `json:"port"`
	Weight                    float32           `json:"weight"`
	Healthy                   bool              `json:"healthy"`
	Enabled                   bool              `json:"enabled"`
	Ephemeral                 bool              `json:"ephemeral"`
	ClusterName               string            `json:"clusterName"`
	ServiceName               string            `json:"serviceName"`
	Metadata                  map[string]string `json:"metadata"`
	InstanceHeartBeatInterval int               `json:"instanceHeartBeatInterval"`
	InstanceHeartBeatTimeOut  int               `json:"instanceHeartBeatTimeOut"`
	IPDeleteTimeout           int               `json:"ipDeleteTimeout"`
}
