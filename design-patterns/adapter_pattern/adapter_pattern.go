package main

import (
	"fmt"

	"github.com/victorvbello/gonotes/design-patterns/adapter_pattern/human"
)

func main() {

	ad := &human.Android{ID: "AND-T-1"}
	ad2 := &human.Android{ID: "AND-T-2"}
	hu := &human.Person{Name: "Victor"}
	hu2 := &human.Person{Name: "Rodrigo"}

	humanList := []human.Human{hu, hu2, human.NewAndroidAdapter(ad), human.NewAndroidAdapter(ad2)}

	for _, h := range humanList {
		fmt.Printf("Name: %s\n", h.GetName())
		fmt.Printf("Walking %d steps\n", 5)
		h.Walk(5)
		fmt.Printf("Talking %d words\n", 3)
		h.Talk(3)
	}
}
