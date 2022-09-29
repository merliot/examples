# Relays

This example Thing is a relay controller using a Raspberry Pi and a 4-relay hat to
independently control 4 lights.  Each relay in the 4-relay hat is tied to a GPIO
pin on the Raspberry Pi.  We'll use the [gobot.io](https://gobot.io) library to toggle
the GPIO pins, in turn toggle the relays, and in turn toggle the lights.

![Relays](relays.png)

## User Interface

![UI](relays_ui.png)

## Hardware Setup
### Parts List
* Rapsberry Pi (any model except Pico)
* 4 Channel Relay Shield for Raspberry Pi
* 12v low voltage lights, wire, and 12v battery

The relay hat uses GPIO pins 31, 33, 35, 37 for relay0, relay1, relay2, and relay3.  If your hat uses different GPIO pins, make the adjustments in the code to match.

Each relay on the relay hat is rated for 2A/24V current max, so keep the load within these limits.  You've been warned.

## Software Setup
### Installation
```
$ git clone https://github.com/merliot/examples.git
```
Files for this example are located in examples/relays

examples/relays/

├── cmd

│   └── relays

│       └── main.go	

└── relays.go

Building


$ cd merle

$ ./build examples/relays

Running


$ cd merle

$ ~/go/bin/relays

Any relays previously left on will be turned off at startup.
