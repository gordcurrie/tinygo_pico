# LED with button

# Equipment

* Raspberry Pi Pico
* 6 * male to male jumper wires
* 1 * button with pull up resistor
* 1 * LED
* 1 * 330 ohm resistor (orange, orange, brown, gold)

# Assembly

* 3.3v (out) pin -> + red power rail
* red power rail -> button 1
* button 2 -> GP15
* gnd pin -> blue power rail
* blue power rail -> LED 1
* LED 2 -> resistor 1
* resistor 2 -> GP14

![Diagram](https://github.com/gordcurrie/tinygo_pico/blob/main/3_led_with_button/pico_LED_button.png)

