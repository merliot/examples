# GPS Locator
This example Thing is a GPS locator running on a Raspberry Pi using a Sixfab
cellular modem to connect to the Internet.  The Sixfab cellular modem also
provides the GPS coordinates.

![Network](gps.png)

## User Interface
Thing's user interface shows a map view with a pin at Thing's current location.
The location is checked and updated every minute.  We'll use Thing Prime running
on a VM to view the map, giving us access to Thing, regardless of it's physical
location.
