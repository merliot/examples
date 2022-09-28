// file: examples/tutorial/blinkv1/blink.go

package main

import (
	"log"
	"time"

	"github.com/merliot/merle"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

type blink struct {
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
		b.State = b.led.State()
		b.Msg = "Update"
		p.Marshal(b)
		p.Broadcast()
		time.Sleep(time.Second)
	}
}

func (b *blink) getState(p *merle.Packet) {
	b.Msg = merle.ReplyState
	p.Marshal(b)
	p.Reply()
}

func (b *blink) Subscribers() merle.Subscribers {
	return merle.Subscribers{
		merle.CmdInit:  b.init,
		merle.CmdRun:   b.run,
		merle.GetState: b.getState,
	}
}

const html = `
<!DOCTYPE html>
<html lang="en">
	<body>
		<img id="LED" style="width: 400px">

		<script>
			image = document.getElementById("LED")

			conn = new WebSocket("{{.WebSocket}}")

			conn.onopen = function(evt) {
				conn.send(JSON.stringify({Msg: "_GetState"}))
			}

			conn.onmessage = function(evt) {
				msg = JSON.parse(evt.data)
				console.log('blink', msg)

				switch(msg.Msg) {
				case "_ReplyState":
				case "Update":
					image.src = "/{{.AssetsDir}}/images/led-" +
						msg.State + ".png"
					break
				}
			}
		</script>
	</body>
</html>`

func (b *blink) Assets() *merle.ThingAssets {
	return &merle.ThingAssets{
		AssetsDir:        "examples/tutorial/assets",
		HtmlTemplateText: html,
	}
}

func main() {
	thing := merle.NewThing(&blink{})

	thing.Cfg.Model = "blink"
	thing.Cfg.Name = "blinky"
	thing.Cfg.PortPublic = 80

	log.Fatalln(thing.Run())
}
