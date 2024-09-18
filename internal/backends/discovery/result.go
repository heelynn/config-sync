package discovery

// DiscoveryResult 服务发现结果
type DiscoveryResult struct {
	Name      string
	Instances []InstanceResult
}

type InstanceResult struct {
	Host   string
	Port   int
	Weight int
}
