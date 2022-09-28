# Temperature/Pressure Sensor

This example uses a cheap ($10) BMP180 barometric pressure/temperature/altitude sensor and a Raspberry Pi  to display temperature and pressure on dial gauges.

<img src="bmp180_sensor.png" width=100>

## User Interface
The User Interface uses gauge.js javascript library.

![UI](bmp180.png)

## Hardware Setup
* Raspberry Pi (all model except Pico)
* BMP180  pressure/temperature/altitude sensor

Wire the BMP180 to the Raspberry Pi:

```
Sensor	Raspberry Pi
---------------------
VCC	3V3 (Pin 1)
GND	GND (Pin 6)
SCL	SCL (pin 5)
SDA	SDA (Pin 3)
```

## Software Setup
* Use the raspi-config tool to enable  I2C interface on Raspberry Pi
```
sudo raspi-config
```
* Install i2c-utils
```
sudo apt update
sudo apt install -y i2c-tools
```
* Verify I2C connections.  There should be at least one address in use.  (if you see all "--" for all the addresses, something is wrong in finding the I2C device: check wiring).
```
sudo i2cdetect -y 1
```

Files for this example are located in examples/bmp180:
```
examples/bmp180/
├── bmp180.go
├── bmp180.png
├── bmp180_sensor.png
├── build
├── cmd
│   └── bmp180
│       └── main.go
└── README.md
```

Building


$ cd merle

$ ./build examples/bmp180

Running


$ cd merle

$ ~/go/bin/bmp180 -h

Usage of /home/merle/go/bin/bmp180:

  -TLS uint

        TLS port

  -prime

        Run as Thing Prime

  -rhost string

        Remote host

  -ruser string

        Remote user (default "merle")

