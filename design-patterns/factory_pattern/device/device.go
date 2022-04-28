package device

import "errors"

type Device interface {
	TurnOn()
	TurnOff()
	BatteryLevel() float32
}

const (
	PHONE_DEVICE = iota
	PC_DEVICE
	TV_DEVICE
)

func CreateNewDevice(deviceType int) (Device, error) {
	switch deviceType {
	case PHONE_DEVICE:
		return new(PhoneDevice), nil
	case PC_DEVICE:
		return new(PcDevice), nil
	default:
		return nil, errors.New("device nof found")
	}
}
