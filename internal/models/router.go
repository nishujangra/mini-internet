package models

type RouterConfig struct {
	ID  string `yaml:"id"`
	ASN uint32 `yaml:"asn"` // Autonomous System Number -> Unique number assigned to specific network
}
