package main

import (
	"fmt"
	"gonotes/design-patterns/builder_pattern/car"
)

func main() {
	newCarBuilder := car.NewCarBuilder()
	teslaModelS := newCarBuilder.AddDoors().AddWheels().AddBattery().AddWindows().Build()
	fmt.Printf("Tesla Model S is: %#v \n", teslaModelS)
	teslaModelX := newCarBuilder.AddDoors().AddWheels().Build()
	fmt.Printf("Tesla Model X is: %#v \n", teslaModelX)
	teslaModel3 := newCarBuilder.AddDoors().AddWheels().AddBattery().AddWindows().AddPaint().Build()
	fmt.Printf("Tesla Model 3 is: %#v \n", teslaModel3)
}
