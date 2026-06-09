package models

type Config struct {
	Router    RouterConfig     `yaml:"router"`
	Neighbors []NeighborConfig `yaml:"neighbors"`
}

type RouterConfig struct {
	ID       string `yaml:"id"`
	ASN      uint32 `yaml:"asn"` // Autonomous System Number -> Unique number assigned to specific network
	ListenIP string `yaml:"listen_ip"`
	Port     string `yaml:"port"`
}

type NeighborConfig struct {
	Address string `yaml:"address"`
	ASN     uint32 `yaml:"asn"`
}
