package main

import (
	"fmt"
	"machine"
	"time"

	"tinygo.org/x/drivers/servo"
)

func main() {
	machine.InitADC()
	c := make(chan uint16, 1)

	go getVals(c)
	adjustSevo(c)
}

func getVals(c chan uint16) {
	pot := machine.ADC{Pin: 26}
	for {
		c <- pot.Get()
		time.Sleep(time.Millisecond * 50)
	}
}

func adjustSevo(c chan uint16) {
	pwm := machine.PWM7
	servo, err := servo.New(pwm, machine.GP15)
	if err != nil {
		fmt.Printf("%#v\n", servo)
		return
	}

	perviousAngle := int16(0)

	for {
		a := <-c

		angle := scale(float32(a))
		if angle <= perviousAngle-100 || angle >= perviousAngle+100 {
			fmt.Printf("Potentiometer reading: %d setting serove to: %d \n", a, angle)
			servo.SetMicroseconds(angle)
			perviousAngle = angle
		}
	}
}

const (
	fromMin float32 = 23000
	fromMax float32 = 65535
	toMin   float32 = 1000
	toMax   float32 = 2200
)

func scale(value float32) int16 {
	v := float32(value)
	return int16(((v - fromMin) * (toMax - toMin) / (fromMax - fromMin)) + toMin)
}
