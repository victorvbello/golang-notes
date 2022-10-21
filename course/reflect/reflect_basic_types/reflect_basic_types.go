// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"reflect"
	"strings"
)

type LocalFloat float32

func main() {
	var x1 float32 = 10.81
	var x2 LocalFloat = 11.82
	var x3 float64 = 12.83

	fmt.Println("- x1 float32", x1)
	inspectValue(x1)
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("- x2 LocalFloat", x2)
	inspectValue(x2)
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("- x3 float64", x3)
	inspectValue(x3)
}

func inspectValue(i interface{}) {

	//reflect.TypeOf()  				=> 	information of type
	//reflect.ValueOf() 				=>	reflect.Value
	//reflect.ValueOf().Type()	=>	reflect.Type of reflect.Value
	//reflect.ValueOf().Kind()	=>	original type of value

	// interface to reflect.Value
	fmt.Println("======= interface to reflect.Value =======")
	v := reflect.ValueOf(i)
	fmt.Printf("v.Type = %s, reflect.TypeOf = %s\n", v.Type(), reflect.TypeOf(i))
	fmt.Println("v.kind [original type]", v.Kind())

	fmt.Println("-- switch i.(type)")
	switch i.(type) {
	case float32:
		fmt.Println("[float32] Real value", v.Float())
		fmt.Printf("v.kind == reflect.Float32 %t\n", v.Kind() == reflect.Float32)
	case float64:
		fmt.Println("[float64] Real value", v.Float())
		fmt.Printf("v.kind == reflect.Float64 %t\n", v.Kind() == reflect.Float64)
	case LocalFloat:
		fmt.Println("[LocalFloat] Real value", v.Float())
		fmt.Printf("v.kind == reflect.Float64 %t\n", v.Kind() == reflect.Float64)
	}

	// reflect.Value to interface

	fmt.Println("======= reflect.Value to interface =======")
	originalValue := v.Interface()

	fmt.Println("-- switch originalValue.(type)")
	switch iV := originalValue.(type) {
	case float32:
		fmt.Printf("[float32] Original value =  %v  reflect.ValueOf().Float() = %v \n", iV, v.Float())
	case float64:
		fmt.Printf("[float64] Original value =  %v  reflect.ValueOf().Float() = %v \n", iV, v.Float())
	case LocalFloat:
		fmt.Printf("[LocalFloat] Original value =  %v  reflect.ValueOf().Float() = %v \n", iV, v.Float())
	}

	// change original value from reflect
	fmt.Println("======= change original value from reflect =======")
	initValue := i
	fmt.Println("Original float32 value", initValue)
	reflectV := reflect.ValueOf(initValue)
	fmt.Println("reflect value", reflectV)
	fmt.Println("Can set reflectV", reflectV.CanSet())
	fmt.Println("-- pointer value")
	switch iV := initValue.(type) {
	case float32:
		reflectPointerV := reflect.ValueOf(&iV)
		pointerElement := reflectPointerV.Elem() // Pointer Dereferencing
		fmt.Println("Can set reflectPointerV", pointerElement.CanSet())
		fmt.Println("Init value ", iV)
		pointerElement.SetFloat(float64(iV) + 0.3)
		fmt.Println("New value ", iV)
	case float64:
		reflectPointerV := reflect.ValueOf(&iV)
		pointerElement := reflectPointerV.Elem() // Pointer Dereferencing
		fmt.Println("Can set reflectPointerV", pointerElement.CanSet())
		fmt.Println("Init value ", iV)
		pointerElement.SetFloat(float64(iV) + 0.4)
		fmt.Println("New value ", iV)
	case LocalFloat:
		reflectPointerV := reflect.ValueOf(&iV)
		pointerElement := reflectPointerV.Elem() // Pointer Dereferencing
		fmt.Println("Can set reflectPointerV", pointerElement.CanSet())
		fmt.Println("Init value ", iV)
		pointerElement.SetFloat(float64(iV) + 0.5)
		fmt.Println("New value ", iV)

	}
}
