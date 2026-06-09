package models

type Config struct {
	Router    RouterConfig     `yaml:"router"`
	Neighbors []NeighborConfig `yaml:"neighbors"`
}
