package main

import (
	"machine"
)

const lightLevelThreshold uint16 = 1500

func main() {
	machine.InitADC()
	c := make(chan uint16, 1)

	go getVal(c)
	adjustLight(c)
}

func getVal(c chan uint16) {
	photoSensor := machine.ADC{Pin: 26}
	for {
		c <- photoSensor.Get()
	}
}

func adjustLight(c chan uint16) {
	led := machine.GP14
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		val := <-c
		if val > lightLevelThreshold {
			led.Low()
		} else {
			led.High()
		}
	}

}
