//go:build tinygo
// +build tinygo

package nano33

import (
	"machine"
	"time"

	"github.com/merliot/merle"
	"tinygo.org/x/drivers/lsm6ds3"
	"tinygo.org/x/drivers/wifinina"
)

// Access point info
const ssid = "Feldman Starlink"
const pass = "itsasecret"

func init() {
	// 2 sec delay otherwise some printlns are missed at startup in serial output
	time.Sleep(2 * time.Second)
}

func (n *nano33) connectAP(ssid, pass string) {
	// These are the default pins for the Arduino Nano33 IoT.
	spi := machine.NINA_SPI

	// Configure SPI for 8Mhz, Mode 0, MSB First
	spi.Configure(machine.SPIConfig{
		Frequency: 8 * 1e6,
		SDO:       machine.NINA_SDO,
		SDI:       machine.NINA_SDI,
		SCK:       machine.NINA_SCK,
	})

	// This is the ESP chip that has the WIFININA firmware flashed on it
	adaptor := wifinina.New(spi,
		machine.NINA_CS,
		machine.NINA_ACK,
		machine.NINA_GPIO0,
		machine.NINA_RESETN)
	adaptor.Configure()

	// Connect to access point
	time.Sleep(2 * time.Second)
	println("Connecting to " + ssid)
	err := adaptor.ConnectToAccessPoint(ssid, pass, 10*time.Second)
	if err != nil { // error connecting to AP
		for {
			println(err)
			time.Sleep(time.Second)
		}
	}

	println("Connected.")

	time.Sleep(2 * time.Second)
	ip, _, _, err := adaptor.GetIP()
	for ; err != nil; ip, _, _, err = adaptor.GetIP() {
		println(err.Error())
		time.Sleep(time.Second)
	}
	println(ip.String())
}

func (n *nano33) init(p *merle.Packet) {
	n.connectAP(ssid, pass)
}

func (n *nano33) run(p *merle.Packet) {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	machine.I2C0.Configure(machine.I2CConfig{})
	accel := lsm6ds3.New(machine.I2C0)
	err := accel.Configure(lsm6ds3.Configuration{})
	if err != nil {
		for {
			println("Failed to configure", err.Error())
			time.Sleep(time.Second)
		}
	}

	for {
		if !accel.Connected() {
			println("LSM6DS3 not connected")
			time.Sleep(time.Second)
			continue
		}

		tempC, _ := accel.ReadTemperature()
		tempC /= 1000
		if tempC != n.TempC {
			n.TempC = tempC
			p.Marshal(n).Broadcast()
		}

		led.Low()
		time.Sleep(time.Millisecond * 500)

		led.High()
		time.Sleep(time.Millisecond * 500)
	}
}
