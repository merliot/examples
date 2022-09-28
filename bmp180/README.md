# Temperature/Pressure Sensor

This example uses a cheap ($10) BMP180 barometric pressure/temperature/altitude sensor and a Raspberry Pi  to display temperature and pressure on dial gauges.
![Sensor](bmp180_sensor.png)

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
