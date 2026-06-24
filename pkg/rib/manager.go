package rib

import (
	"net/netip"
	"sort"
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

func (ribmg *RIBManager) Insert(route models.Route) {
	ribmg.mu.Lock()
	defer ribmg.mu.Unlock()

	existing := ribmg.routes[route.Prefix]
	for i, r := range existing {
		if r.Protocol == route.Protocol {
			existing[i] = route
			ribmg.routes[route.Prefix] = existing
			return
		}
	}

	ribmg.routes[route.Prefix] = append(existing, route)

	sort.Slice(ribmg.routes[route.Prefix], func(i, j int) bool {
		return ribmg.routes[route.Prefix][i].AdminDistance < ribmg.routes[route.Prefix][j].AdminDistance
	})

}

func (ribmg *RIBManager) Delete() {}

func (ribmg *RIBManager) Lookup() {}

func (ribmg *RIBManager) List() {}
