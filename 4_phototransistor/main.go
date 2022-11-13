package main

import (
	"machine"
)

func main() {
	machine.InitADC()

	photoSensor := machine.ADC{Pin: 26}

	var lightLevelThreshold uint16 = 1300

	led := machine.GP14
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		if photoSensor.Get() > lightLevelThreshold {
			led.Low()
		} else {
			led.High()
		}
	}
}
