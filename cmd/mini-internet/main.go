package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/goccy/go-yaml"
	"github.com/nishujangra/mini-internet/internal/models"
)

func main() {
	confFile, err := ioutil.ReadFile("configs/router.yaml")
	if err != nil {
		log.Printf("confFile err   #%v ", err)
	}

	var config models.Config

	err = yaml.Unmarshal(confFile, &config)
	if err != nil {
		log.Printf("unmarshal err   #%v ", err)
	}

	fmt.Printf("Router ID: %s\n", config.Router.ID)
}
