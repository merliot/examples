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
Parts list below.  If you don't have the hardware, there is a demo mode so you
can run Thing in software-only mode.  (See Running below).

### Parts List:
* Rapsberry Pi (any model except Pico)
* Sixfab Raspberry Pi 4G/LTE Cellular Modem Kit.  I'm using the Telit LE910CF-NF
(North America) LTE module option, $125 for the kit.  It includes a SIM card
for $2/month + data
