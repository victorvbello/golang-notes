package main

import (
	"fmt"
	"reflect"
)

type Aircraft interface {
	Fly(flyCode string) (result string)
	Landing()
}

type Airplane struct {
	Code     string
	WingType string
}

func (a *Airplane) Fly(flyCode string) string {
	fmt.Printf("Airplane: %s is flying whit code: %s\n", a.Code, flyCode)
	return fmt.Sprintf("[%s]:%s", a.Code, flyCode)
}

func (a *Airplane) Landing() {
	fmt.Printf("Airplane: %s is landing\n", a.Code)
}

type Helicopter struct {
	Code   string
	Blades int
}

func (h *Helicopter) Fly(flyCode string) error {
	fmt.Printf("Helicopter: %s is flying whit code: %s\n", h.Code, flyCode)
	return nil
}

func (h *Helicopter) Landing() {
	fmt.Printf("Helicopter: %s is landing\n", h.Code)
}

func main() {
	fmt.Println("======= Airplane =======")
	sukhoi := new(Airplane)
	sukhoi.Code = "Su-47"
	sukhoi.WingType = "Forward swept"
	inspectInterface(sukhoi)

	fmt.Println("======= Helicopter =======")
	blackHawk := new(Helicopter)
	blackHawk.Code = "UH-60"
	blackHawk.Blades = 4
	inspectInterface(blackHawk)
}

func inspectInterface(i interface{}) {
	reflectV := reflect.ValueOf(i)
	reflectT := reflectV.Type()

	emptyInterfaceE := reflect.TypeOf((*Aircraft)(nil)).Elem()

	fmt.Println("Input item implements Aircraft interface :", map[bool]string{true: "yes", false: "no"}[reflectT.Implements(emptyInterfaceE)])

	if reflectT.Implements(emptyInterfaceE) {
		flyFunc := reflectV.MethodByName("Fly")
		args := []reflect.Value{reflect.ValueOf("KKJH-1232-XXDES")}
		results := flyFunc.Call(args)
		for i, rv := range results {
			fmt.Printf("Call Fly return: (%d) %v of type:%s \n", i, rv, rv.Kind())
		}
	}
}
