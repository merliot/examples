// file: examples/nano33/cmd/nano33/main.go

// tinygo flash -target=arduino-nano33 -serial usb -ldflags '-X "main.ssid=xxx" -X "main.pass=xxx"' cmd/nano33/main.go
//
// minicom -c on -D /dev/ttyACM0 -b 115200

package main

import (
	"log"
	"strings"

	"github.com/merliot/merle"
	"github.com/merliot/examples/nano33"
)

var (
	ssid string
	pass string
)

func main() {
	nano := nano33.NewNano33()
	macAddr := nano.ConnectAP(ssid, pass)

	thing := merle.NewThing(nano)
	thing.Cfg.Id = strings.Replace(macAddr, ":", "_", -1)
	thing.Cfg.MotherHost = "demos.merliot.org"
	thing.Cfg.MotherUser = "foobar" // not used, but need non-empty string
	thing.Cfg.PortPrivate = 8070
	thing.Cfg.MotherPortPrivate = 8070

	log.Fatalln(thing.Run())
}
