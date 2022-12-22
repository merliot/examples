// file: examples/relays/relays.go

package relays

import (
	"embed"
	"sync"

	"github.com/merliot/merle"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

type Relays struct {
	sync.Mutex
	drivers [4]*gpio.RelayDriver
	Msg     string
	States  [4]bool
}

func NewRelays() merle.Thinger {
	return &Relays{}
}

func (r *Relays) run(p *merle.Packet) {
	adaptor := raspi.NewAdaptor()
	adaptor.Connect()

	r.drivers[0] = gpio.NewRelayDriver(adaptor, "31") // GPIO 6
	r.drivers[1] = gpio.NewRelayDriver(adaptor, "33") // GPIO 13
	r.drivers[2] = gpio.NewRelayDriver(adaptor, "35") // GPIO 19
	r.drivers[3] = gpio.NewRelayDriver(adaptor, "37") // GPIO 26

	for _, driver := range r.drivers {
		driver.Start()
		driver.Off()
	}

	select {}
}

func (r *Relays) getState(p *merle.Packet) {
	r.Lock()
	r.Msg = merle.ReplyState
	p.Marshal(r)
	r.Unlock()
	p.Reply()
}

func (r *Relays) saveState(p *merle.Packet) {
	r.Lock()
	p.Unmarshal(r)
	r.Unlock()
}

type MsgClick struct {
	Msg   string
	Relay int
	State bool
}

func (r *Relays) click(p *merle.Packet) {
	var msg MsgClick
	p.Unmarshal(&msg)

	r.Lock()
	r.States[msg.Relay] = msg.State
	r.Unlock()

	if p.IsThing() {
		if msg.State {
			r.drivers[msg.Relay].On()
		} else {
			r.drivers[msg.Relay].Off()
		}
	}

	p.Broadcast()
}

func (r *Relays) Subscribers() merle.Subscribers {
	return merle.Subscribers{
		merle.CmdRun:     r.run,
		merle.GetState:   r.getState,
		merle.ReplyState: r.saveState,
		"Click":          r.click,
	}
}

//go:embed index.html
var fs embed.FS

func (r *Relays) Assets() merle.ThingAssets {
	return merle.ThingAssets{
		FS: fs,
	}
}
