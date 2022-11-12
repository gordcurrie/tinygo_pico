package main

import (
	"machine"
	"time"
)

var period uint64 = 1e9 / 500

func main() {
	led := machine.LED // onboard LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		led.High()
		time.Sleep(time.Millisecond * 500)
		led.Low()
		time.Sleep(time.Millisecond * 500)
	}
}
