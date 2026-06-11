package rib

import (
	"net/netip"
	"sync"

	"github.com/nishujangra/mini-internet/internal/models"
)

// Routing Information Base // Stores all known routes.
type RIBManager struct {
	mu     sync.RWMutex
	routes map[netip.Prefix][]models.Route
}

func NewRIBManager() *RIBManager {
	return &RIBManager{routes: make(map[netip.Prefix][]models.Route)}
}

func (ribmg *RIBManager) Insert() {}

func (ribmg *RIBManager) Delete() {}

func (ribmg *RIBManager) Lookup() {}

func (ribmg *RIBManager) List() {}
