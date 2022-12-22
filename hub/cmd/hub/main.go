package main

import (
	"flag"
	"log"

	"github.com/merliot/merle"
	"github.com/merliot/examples/hub"
)

func main() {
	thing := merle.NewThing(hub.NewHub())

	thing.Cfg.Model = "hub"
	thing.Cfg.Name = "hubby"

	thing.Cfg.PortPublic = 80
	thing.Cfg.PortPrivate = 6000

	flag.UintVar(&thing.Cfg.PortPublicTLS, "TLS", 0, "TLS port")
	flag.Parse()

	log.Fatalln(thing.Run())
}
