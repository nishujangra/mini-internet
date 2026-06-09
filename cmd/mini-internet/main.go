package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/goccy/go-yaml"
	"github.com/nishujangra/mini-internet/internal/models"
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

	// Accept connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}

		log.Printf(
			"neighbor connected from %s",
			conn.RemoteAddr(),
		)

		conn.Close()
	}
}
