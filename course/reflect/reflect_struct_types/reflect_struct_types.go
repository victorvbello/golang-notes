package main

import (
	"fmt"
	"reflect"
)

type CustomStringType string

type MyTestStruct struct {
	ID       int32            `flag:"id" detail:"id of item"`
	Name     string           `flag:"name" detail:"name of item"`
	Alias    CustomStringType `flag:"alias" detail:"alias of item"`
	Distance float32          `flag:"distance" detail:"distance of item"`
}

func main() {
	Item := MyTestStruct{1, "Plane", "The bird", 7453.89}
	inspectValue(Item)
	changeValue(&Item)
	fmt.Println("======= finalValue =======")
	fmt.Printf("Id: %d, Name: %s, Alias: %s, Distance: %.4f\n", Item.ID, Item.Name, Item.Alias, Item.Distance)
}

func inspectValue(i interface{}) {
	fmt.Println("======= inspectValue =======")
	reflectV := reflect.ValueOf(i)
	if reflectV.Kind() != reflect.Struct {
		fmt.Println("Kind of input is not a struct")
		return
	}
	reflectT := reflect.TypeOf(i)

	for i := 0; i < reflectT.NumField(); i++ {
		fieldReflectT := reflectT.Field(i)
		fieldReflectV := reflectV.Field(i)
		fmt.Printf("Field Name: '%s', Field kind: '%s', Field type: '%s', Filed value: '%v'\n", fieldReflectT.Name, fieldReflectT.Type.Kind(), fieldReflectT.Type, fieldReflectV.Interface())
	}

	fmt.Println("--- Tags ---")

	for i := 0; i < reflectT.NumField(); i++ {
		fieldReflectT := reflectT.Field(i)
		fmt.Printf("Field Name: '%s' flag: '%s', detail: '%s'\n", fieldReflectT.Name, fieldReflectT.Tag.Get("flag"), fieldReflectT.Tag.Get("detail"))
	}
}

func changeValue(i interface{}) {
	fmt.Println("======= changeValue =======")
	reflectV := reflect.ValueOf(i)
	if reflectV.Kind() != reflect.Ptr {
		fmt.Println("Kind of input is not a pointer")
		return
	}
	reflectE := reflectV.Elem()

	if reflectE.Kind() != reflect.Struct {
		fmt.Println("Kind of input is not a struct")
		return
	}

	// Change value of ID
	reflectE.Field(0).SetInt(123)
	// Change value of Name
	reflectE.Field(1).SetString("Plane-V2")
	// Change value of Distance
	reflectE.Field(3).SetFloat(105643.77)

	reflectT := reflectE.Type()

	for i := 0; i < reflectT.NumField(); i++ {
		fieldReflectT := reflectT.Field(i)
		fieldReflectV := reflectE.Field(i)
		fmt.Printf("Field Name: '%s', Field kind: '%s', Field type: '%s', Filed value: '%v'\n", fieldReflectT.Name, fieldReflectT.Type.Kind(), fieldReflectT.Type, fieldReflectV.Interface())
	}
}
