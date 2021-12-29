// Package hero make heroes can attack and fly
package hero

import (
	"fmt"
	"strings"
)

// Hero struct form make new heroes
type Hero struct {
	// Hero Name
	Name string
	// Hero Alias
	Alias string
}

// Attack is a method set, permit the hero launch attack
func (h Hero) Attack() {
	fmt.Println("---> Attack!!")
	fmt.Println("POW!! ")
	fmt.Println("\tPOW!! ")
	fmt.Println("\t\tPOW!! ")
}

// Fly is a method set, permit the hero fly
func (h Hero) Fly(kms int) {
	distance := "\t"
	fmt.Println("---> FLY!!")
	for i := 0; i < kms; i++ {
		fmt.Printf("%s----->\n", strings.Repeat(distance, i))
	}
}
