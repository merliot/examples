# GPS Locator
This example Thing is a GPS locator running on a Raspberry Pi using a Sixfab
cellular modem to connect to the Internet.  The Sixfab cellular modem also
provides the GPS coordinates.

![Network](gps.png)

## User Interface
Thing's user interface shows a map view with a pin at Thing's current location.
The location is checked and updated every minute.

![UI](gps_ui.png)

## Hardware Setup
Parts list below.  If you don't have the hardware, don't worrry, you can run
in demo mode on any system with Go installed.  (See Running below).

### Parts List:
* Rapsberry Pi (any model except Pico)
* [Sixfab](https://sixfab.com) Raspberry Pi 4G/LTE Cellular Modem Kit.  I'm using the Telit LE910CF-NF
(North America) LTE module option, $125 for the [kit](https://sixfab.com/product/raspberry-pi-4g-lte-modem-kit/).
It includes a SIM card for $2/month + data.

### Sixfab Setup
* Follow the Sixfab [instructions](https://docs.sixfab.com/) to install the Sixfab software on the Raspberry Pi.
* Activate the Sixfab SIM

## Software Setup
### Installation
git clone https://github.com/merliot/examples.git

Files for this example are located in examples/gps/ and examples/telit.
```
examples/gps
├── build
├── cmd
│   └── gps
│       └── main.go
├── gps.go
├── gps.png
├── gps_ui.png
└── README.md

examples/telit/
└── telit.go
```

Thing Prime runs on a cloud VM.  I'm using a Linode VM to host Thing Prime.


