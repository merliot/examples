// file: examples/nano33/cmd/nano33/main.go

// tinygo flash -target=arduino-nano33 -serial usb examples/nano33/cmd/nano33/main.go
//
// minicom -c on -D /dev/ttyACM0 -b 115200

package main

import (
	"log"

	"github.com/merliot/merle"
	"github.com/merliot/examples/nano33"
)

func main() {
	thing := merle.NewThing(nano33.NewNano33())

	thing.Cfg.MotherHost = "example.org"
	thing.Cfg.MotherUser = "foobar" // not used, but need not-empty string
	thing.Cfg.PortPrivate = 8080
	thing.Cfg.MotherPortPrivate = 8080

	log.Fatalln(thing.Run())
}
