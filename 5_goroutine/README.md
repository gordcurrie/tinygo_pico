# LED with button

# Equipment

* Raspberry Pi Pico
* 7 * male to male jumper wires
* 4 * LED
* 4 * 330 ohm resistor (orange, orange, brown, gold)
* 1 * rotary potentiometer

# Assembly

* 3.3v (out) pin -> + red power rail
* red power rail -> potentiometer 1
* potentiometer 2 -> GP26
* gnd pin -> blue power rail
* blue power rail -> LED 1
* LED 2 -> 330 ohm resistor 1
* 330 ohm resistor 2 -> GP12
* blue power rail -> LED 1
* LED 2 -> 330 ohm resistor 1
* 330 ohm resistor 2 -> GP13
* blue power rail -> LED 1
* LED 2 -> 330 ohm resistor 1
* 330 ohm resistor 2 -> GP14
* blue power rail -> LED 1
* LED 2 -> 330 ohm resistor 1
* 330 ohm resistor 2 -> GP15


![Diagram](https://github.com/gordcurrie/tinygo_pico/blob/main/6_multiple_leds/multiple_leds.png)

