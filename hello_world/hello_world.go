package main

import (
	"embed"
	"log"

	"github.com/merliot/merle"
)

type hello struct {
}

func (h *hello) Subscribers() merle.Subscribers {
	return merle.Subscribers{
		merle.CmdRun: merle.RunForever,
	}
}

//go:embed index.html
var fs embed.FS

func (h *hello) Assets() merle.ThingAssets {
	return merle.ThingAssets{
		FS: fs,
	}
}

func main() {
	thing := merle.NewThing(&hello{})
	thing.Cfg.PortPublic = 80
	log.Fatalln(thing.Run())
}
