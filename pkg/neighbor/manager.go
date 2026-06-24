package neighbor

import (
	"net"
	"sync"

	"github.com/nishujangra/mini-internet/internal/models"
)

type NeighborManager struct {
	mu       sync.RWMutex
	neighbor map[string]*models.Neighbor
}

func NewNeighborManager() *NeighborManager {
	return &NeighborManager{neighbor: make(map[string]*models.Neighbor)}
}

func (nmg *NeighborManager) Add(conn net.Conn) {
	nmg.mu.Lock()
	defer nmg.mu.Unlock()

	nmg.neighbor[conn.RemoteAddr().String()] = &models.Neighbor{
		Address:   conn.RemoteAddr().String(),
		ASN:       0,
		Connected: true,
		Conn:      conn,
	}
}

func (nmg *NeighborManager) Remove() {}

func (nmg *NeighborManager) Lookup() {}

func (nmg *NeighborManager) List() {}
