// file: examples/bmp180/bmp180.go

package bmp180

import (
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

const html = `
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<style>
		#overlay {
			position: fixed;
			display: none;
			width: 100%;
			height: 100%;
			top: 0;
			left: 0;
			right: 0;
			bottom: 0;
			background-color: rgba(0,0,0,0.5);
			z-index: 2000;
			cursor: wait;
		}
		#offline {
			position: absolute;
			top: 50%;
			left: 50%;
			font-size: 50px;
			color: white;
			transform: translate(-50%,-50%);
		}
		</style>
	</head>
	<body style="background-color:orange">
		<canvas id="temp_gauge"></canvas>
		<canvas id="pres_gauge"></canvas>
		<div id="overlay">
			<div id="offline">Offline</div>
		</div>

		<script src="//cdn.rawgit.com/Mikhus/canvas-gauges/gh-pages/download/2.1.7/radial/gauge.min.js"></script>

		<script>
			var conn
			var online = false

			temp_gauge = document.getElementById("temp_gauge")
			pres_gauge = document.getElementById("pres_gauge")

			var tempGauge = new RadialGauge({
				renderTo: temp_gauge,
				majorTicks: [0,20,40,60,80,100,120],
				minorTicks: 10,
				highlights: [
					{from: 80, to: 100, color: "orange"},
					{from: 100, to: 120, color: "red"},
				],
				maxValue: 120,
				units: "F",
				title: "Temperature",
				width: 300,
				height: 300,
				valueInt: 0,
				valueDec: 0,
			})
			var presGauge = new RadialGauge({
				renderTo: pres_gauge,
				majorTicks: [40,60,80,100,120],
				minorTicks: 10,
				highlights: [
					{from: 0, to: 101.325, color: "#c9df8a"},
					{from: 101.325, to: 120, color: "#6fc4db"},
				],
				maxValue: 120,
				units: "kPa",
				title: "Pressure",
				width: 300,
				height: 300,
				valueInt: 0,
				valueDec: 0,
			})

			function getState() {
				conn.send(JSON.stringify({Msg: "_GetState"}))
			}

			function getIdentity() {
				conn.send(JSON.stringify({Msg: "_GetIdentity"}))
			}

			function save(msg) {
				tempGauge.value = msg.Temperature
				presGauge.value = msg.Pressure
			}

			function show() {
				overlay = document.getElementById("overlay")
				if (online) {
					tempGauge.draw()
					presGauge.draw()
					overlay.style.display = "none"
				} else {
					overlay.style.display = "block"
				}
			}

			function connect() {
				conn = new WebSocket("{{.WebSocket}}")

				conn.onopen = function(evt) {
					getIdentity()
				}

				conn.onclose = function(evt) {
					online = false
					show()
					setTimeout(connect, 1000)
				}

				conn.onerror = function(err) {
					conn.close()
				}

				conn.onmessage = function(evt) {
					msg = JSON.parse(evt.data)
					console.log('bmp180', msg)

					switch(msg.Msg) {
					case "_ReplyIdentity":
					case "_EventStatus":
						online = msg.Online
						getState()
						break
					case "_ReplyState":
					case "Update":
						save(msg)
						show()
						break
					}
				}
			}

			connect()
		</script>
	</body>
</html>`

func (b *Bmp180) Assets() merle.ThingAssets {
	return merle.ThingAssets{
		HtmlTemplateText: html,
	}
}
