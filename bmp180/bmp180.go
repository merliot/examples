// file: examples/bmp180/bmp180.go

package bmp180

import (
	"embed"
	"math"
	"sync"
	"time"

	"github.com/merliot/merle"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
)

type Bmp180 struct {
	sync.Mutex
	driver      *i2c.BMP180Driver
	Msg         string
	Temperature int
	Pressure    int
}

func NewBmp180() *Bmp180 {
	return &Bmp180{}
}

func (b *Bmp180) init(p *merle.Packet) {
	adaptor := raspi.NewAdaptor()
	adaptor.Connect()
	b.driver = i2c.NewBMP180Driver(adaptor)
	b.driver.Start()
}

func (b *Bmp180) run(p *merle.Packet) {
	for {
		var changed bool = false

		temp, _ := b.driver.Temperature()
		pres, _ := b.driver.Pressure()

		temp = (temp * 1.8) + 32.0
		pres = pres / 1000.0

		newTemp := int(math.Round(float64(temp)))
		newPres := int(math.Round(float64(pres)))

		b.Lock()
		if newTemp != b.Temperature || newPres != b.Pressure {
			b.Msg = "Update"
			b.Temperature = newTemp
			b.Pressure = newPres
			p.Marshal(b)
			changed = true
		}
		b.Unlock()

		if changed {
			p.Broadcast()
		}

		time.Sleep(time.Second)
	}
}

func (b *Bmp180) getState(p *merle.Packet) {
	b.Lock()
	b.Msg = merle.ReplyState
	p.Marshal(b)
	b.Unlock()
	p.Reply()
}

func (b *Bmp180) saveState(p *merle.Packet) {
	b.Lock()
	p.Unmarshal(b)
	b.Unlock()
}

func (b *Bmp180) update(p *merle.Packet) {
	b.saveState(p)
	p.Broadcast()
}

func (b *Bmp180) Subscribers() merle.Subscribers {
	return merle.Subscribers{
		merle.CmdInit:    b.init,
		merle.CmdRun:     b.run,
		merle.GetState:   b.getState,
		merle.ReplyState: b.saveState,
		"Update":         b.update,
	}
}

//go:embed index.html bmp180.go cmd
var fs embed.FS

func (b *Bmp180) Assets() merle.ThingAssets {
	return merle.ThingAssets{
		FS: fs,
	}
}
