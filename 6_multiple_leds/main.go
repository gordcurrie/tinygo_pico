package main

import (
	"machine"
	"time"
)

const (
	th1 uint16 = 15000
	th2 uint16 = 25000
	th3 uint16 = 35000
	th4 uint16 = 40000
)

func main() {
	machine.InitADC()
	c := make(chan uint16, 1)

	go getVal(c)
	adjustLight(c)
}

func getVal(c chan uint16) {
	pot := machine.ADC{Pin: 26}
	for {
		c <- pot.Get()
		time.Sleep(time.Millisecond * 50)
	}
}

func adjustLight(c chan uint16) {
	led12 := machine.GP12
	led12.Configure(machine.PinConfig{Mode: machine.PinOutput})

	led13 := machine.GP13
	led13.Configure(machine.PinConfig{Mode: machine.PinOutput})

	led14 := machine.GP14
	led14.Configure(machine.PinConfig{Mode: machine.PinOutput})

	led15 := machine.GP15
	led15.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		val := <-c
		if val > th1 {
			led12.Low()
		} else {
			led12.High()
		}
		if val > th2 {
			led13.Low()
		} else {
			led13.High()
		}
		if val > th3 {
			led14.Low()
		} else {
			led14.High()
		}
		if val > th4 {
			led15.Low()
		} else {
			led15.High()
		}
		println(val)
	}
}
