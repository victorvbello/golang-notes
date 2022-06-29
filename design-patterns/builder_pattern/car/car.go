package car

import "strings"

const (
	CAR_PART_WINDOWS = "W"
	CAR_PART_DOORS   = "D"
	CAR_PART_PAINT   = "P"
	CAR_PART_BATTERY = "B"
	CAR_PART_WHEELS  = "H"
)

type Car struct {
	Doors               bool
	Paint               bool
	Windows             bool
	Wheels              bool
	Battery             bool
	ConstructPercentage float64
}

type CarBuilder struct {
	manifest string
}

func NewCarBuilder() *CarBuilder {
	return new(CarBuilder)
}

func (c *CarBuilder) AddDoors() *CarBuilder {
	c.manifest += CAR_PART_DOORS
	return c
}

func (c *CarBuilder) AddPaint() *CarBuilder {
	c.manifest += CAR_PART_PAINT
	return c
}

func (c *CarBuilder) AddWindows() *CarBuilder {
	c.manifest += CAR_PART_WINDOWS
	return c
}

func (c *CarBuilder) AddWheels() *CarBuilder {
	c.manifest += CAR_PART_WHEELS
	return c
}

func (c *CarBuilder) AddBattery() *CarBuilder {
	c.manifest += CAR_PART_BATTERY
	return c
}

func (c *CarBuilder) Build() *Car {
	newCar := &Car{
		Doors:               strings.Contains(c.manifest, CAR_PART_DOORS),
		Paint:               strings.Contains(c.manifest, CAR_PART_PAINT),
		Windows:             strings.Contains(c.manifest, CAR_PART_WINDOWS),
		Wheels:              strings.Contains(c.manifest, CAR_PART_WHEELS),
		Battery:             strings.Contains(c.manifest, CAR_PART_BATTERY),
		ConstructPercentage: float64(len(c.manifest)) * 20.0,
	}
	c.manifest = ""
	return newCar
}
