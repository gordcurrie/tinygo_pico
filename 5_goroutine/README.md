# LED with button

# Equipment

* Raspberry Pi Pico
* 7 * male to male jumper wires
* 1 * LED
* 1 * 330 ohm resistor (orange, orange, brown, gold)
* 1 * phototransistor
* 1 * 10k ohm resistor (brown, black, orange, gold)

# Assembly

* 3.3v (out) pin -> + red power rail
* red power rail -> phototransistor 1
* phototransistor 2 -> GP26
* phototransistor 2 -> 10k resistor 1
* 10k resistor 2 -> blue power rail
* gnd pin -> blue power rail
* blue power rail -> LED 1
* LED 2 -> resistor 1
* resistor 2 -> GP14

![Diagram](https://github.com/gordcurrie/tinygo_pico/blob/main/5_goroutine/phototransistor.png)

