package models

import (
	"net"
)

type Neighbor struct {
	Address string
	ASN     uint32

	Connected bool // default = false

	Conn net.Conn // default = nil
}
