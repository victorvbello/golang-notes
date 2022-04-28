package device

import "fmt"

type PcDevice struct {
}

func (d *PcDevice) TurnOn() {
	fmt.Println("TurnOn PC")
}

func (d *PcDevice) TurnOff() {
	fmt.Println("TurnOff PC")
}

func (d *PcDevice) BatteryLevel() float32 {
	fmt.Println("BatteryLevel PC")
	return 0.0
}
