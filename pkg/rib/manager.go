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

// You insert a route
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

// You never delete route, you delete protocol's entry for that prefix
func (ribmg *RIBManager) Delete(prefix netip.Prefix, protocol models.Protocol) {
	ribmg.mu.Lock()
	defer ribmg.mu.Unlock()

	existing := ribmg.routes[prefix]
	for i, r := range existing {
		if r.Protocol == protocol {
			existing = append(existing[:i], existing[i+1:]...)

			// if no routes left for this prefix
			if len(existing) == 0 {
				delete(ribmg.routes, prefix)
			} else {
				ribmg.routes[prefix] = existing
			}

			return
		}
	}
}

// Lookup is based on ip addr and return route
func (ribmg *RIBManager) Lookup(addr netip.Addr) *models.Route {
	ribmg.mu.RLock()
	defer ribmg.mu.RUnlock()

	var best *models.Route
	bestBits := -1

	for prefix, routes := range ribmg.routes {
		if prefix.Contains(addr) && prefix.Bits() > bestBits {
			bestBits = prefix.Bits()
			r := routes[0]
			best = &r
		}
	}

	return best
}

func (ribmg *RIBManager) List() {}
