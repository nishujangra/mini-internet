package models

// Runtime
type Router struct {
	ID       string
	ASN      uint32
	ListenIP string
	Port     string

	// This is runtime state, when connection is estabilished, will update the Neighbors
	Neighbors map[string]*Neighbor
}
