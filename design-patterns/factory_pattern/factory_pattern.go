package main

import (
	"fmt"
	"gonotes/design-patterns/factory_pattern/device"
)

var err error
var newDevice device.Device

func ShowPattern(dType int) {
	newDevice, err = device.CreateNewDevice(dType)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Printf("Device type =>%T\n", newDevice)
	newDevice.TurnOn()
	newDevice.TurnOff()
	fmt.Printf("BatteryLevel => %0.2f\n", newDevice.BatteryLevel())
	fmt.Println("------")
}

func main() {
	ShowPattern(device.PHONE_DEVICE)
	ShowPattern(device.PC_DEVICE)
	ShowPattern(device.TV_DEVICE)
}
