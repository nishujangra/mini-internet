package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/netip"
	"os"

	"github.com/goccy/go-yaml"
	"github.com/nishujangra/mini-internet/internal/models"
	"github.com/nishujangra/mini-internet/pkg/neighbor"
	"github.com/nishujangra/mini-internet/pkg/rib"
)

func main() {
	confFile, err := os.ReadFile("configs/router.yaml")
	if err != nil {
		log.Printf("confFile err   #%v ", err)
	}

	var config models.Config

	err = yaml.Unmarshal(confFile, &config)
	if err != nil {
		log.Printf("unmarshal err   #%v ", err)
	}

	fmt.Printf("Router ID: %s\n", config.Router.ID)
	for i, neigh := range config.Neighbors {
		fmt.Printf("Address of Neighbor %d is %v\n", i, neigh)
	}

	// Start the listner (default listen at tcp)
	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%s", config.Router.ListenIP, config.Router.Port))
	if err != nil {
		fmt.Printf("Error in starting listner")
	}

	nmg := neighbor.NewNeighborManager()
	ribmg := rib.NewRIBManager()

	// TEST
	ribmg.Insert(models.Route{
		Prefix:        netip.MustParsePrefix("10.0.0.0/24"),
		NextHop:       netip.MustParseAddr("192.168.1.1"),
		Protocol:      models.OSPF,
		AdminDistance: 110,
		Metric:        10,
	})

	go func() {
		http.HandleFunc("/routes", func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode(ribmg.List())
		})
		http.ListenAndServe(":8080", nil)
	}()

	// Accept connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}

		nmg.Add(conn)

		log.Printf(
			"neighbor connected from %s",
			conn.RemoteAddr(),
		)

		go func(c net.Conn) {
			defer c.Close()
		}(conn)
	}
}
