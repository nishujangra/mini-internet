package models

import "net/netip"

type Protocol uint8

const (
	Connected Protocol = iota // to generate successive valules in golang in consts
	OSPF
	EBGP // external BGP
	IBGP // internal BGP
)

type Route struct {
	Prefix        netip.Prefix
	NextHop       netip.Addr
	Interface     string
	Metric        uint32
	Protocol      Protocol
	AdminDistance uint8 // AD, distance from administrative router, useful in multiple routes
}

// Routing Information Base // Stores all known routes.
type RIB struct {
	routes map[netip.Prefix][]Route
}
