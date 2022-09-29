# Relays

This example Thing is a relay controller using a Raspberry Pi and a 4-relay hat to
independently control 4 lights.  Each relay in the 4-relay hat is tied to a GPIO
pin on the Raspberry Pi.  We'll use the [gobot.io](https://gobot.io) library to toggle
the GPIO pins, in turn toggle the relays, and in turn toggle the lights.
