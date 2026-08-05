package main

import (
	"fmt"

	"github.com/victorvbello/gonotes/design-patterns/decorator_pattern/shakes"
)

func main() {
	shake := new(shakes.MilkShake)
	shakeWithChocolate := shakes.NewChocolateTopping(shake)
	shakeWithStrawberry := shakes.NewStrawberryTopping(shakeWithChocolate)
	shakeWithDonut := shakes.NewDonutTopping(shakeWithStrawberry)

	fmt.Println("Description:", shakeWithDonut.GetDescription())
	fmt.Println("Cost:", shakeWithDonut.Cost())
}
