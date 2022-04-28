package device

import "fmt"

type PhoneDevice struct {
	Model string
	Brand string
}

func (d *PhoneDevice) TurnOn() {
	fmt.Println("TurnOn Phone")
}

func (d *PhoneDevice) TurnOff() {
	fmt.Println("TurnOff Phone")
}

func (d *PhoneDevice) BatteryLevel() float32 {
	fmt.Println("BatteryLevel Phone")
	return 78.90
}
