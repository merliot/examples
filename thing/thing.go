package main

import (
	"embed"

	"github.com/merliot/merle"
)

type thing struct {
}

func (t *thing) Subscribers() merle.Subscribers {
	return merle.Subscribers{
		merle.CmdRun: merle.RunForever,
	}
}

//go:embed index.html
var fs embed.FS

func (t *thing) Assets() merle.ThingAssets {
	return merle.ThingAssets{
		FS: fs,
	}
}

func main() {
	thing := merle.NewThing(&thing{})
	thing.Cfg.PortPublic = 80
	thing.Run()
}
