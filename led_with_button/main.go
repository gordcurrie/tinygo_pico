package main

import (
	"machine"
)

var period uint64 = 1e9 / 500

func main() {
	button := machine.GP15
	button.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	led := machine.GP14
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		if button.Get() {
			led.High()
		} else {
			led.Low()
		}
	}
}
