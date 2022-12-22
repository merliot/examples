// file: examples/tutorial/blinkv1/blink.go

package main

import (
	"embed"
	"log"
	"sync"
	"time"

	"github.com/merliot/merle"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

type blink struct {
	sync.Mutex
	adaptor *raspi.Adaptor
	led     *gpio.LedDriver
	Msg     string
	State   bool
}

func (b *blink) init(p *merle.Packet) {
	b.adaptor = raspi.NewAdaptor()
	b.adaptor.Connect()
	b.led = gpio.NewLedDriver(b.adaptor, "11")
	b.led.Start()
	b.State = b.led.State()
}

func (b *blink) run(p *merle.Packet) {
	for {
		b.led.Toggle()
		b.Lock()
		b.State = b.led.State()
		b.Msg = "Update"
		p.Marshal(b)
		b.Unlock()
		p.Broadcast()
		time.Sleep(time.Second)
	}
}

func (b *blink) getState(p *merle.Packet) {
	b.Lock()
	b.Msg = merle.ReplyState
	p.Marshal(b)
	b.Unlock()
	p.Reply()
}

func (b *blink) Subscribers() merle.Subscribers {
	return merle.Subscribers{
		merle.CmdInit:  b.init,
		merle.CmdRun:   b.run,
		merle.GetState: b.getState,
	}
}

//go:embed index.html images
var fs embed.FS

func (b *blink) Assets() merle.ThingAssets {
	return merle.ThingAssets{
		FS: fs,
	}
}

func main() {
	thing := merle.NewThing(&blink{})

	thing.Cfg.Model = "blink"
	thing.Cfg.Name = "blinky"
	thing.Cfg.PortPublic = 80

	log.Fatalln(thing.Run())
}
