```go
package main

import (
	"encoding/json"
	"fmt"
)

var MarshalJSONCalled int
var MarshalJSONCalledExpected int

type customType struct {
	Valid bool
	Value int
}

// MarshalJSON for customType
func (ct *customType) MarshalJSON() ([]byte, error) {
	MarshalJSONCalled++
	fmt.Println("MarshalJSON", ct)
	if !ct.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ct.Value)
}

// UnmarshalJSON for customType
func (ct *customType) UnmarshalJSON(b []byte) error {
	fmt.Println("UnmarshalJSON", ct)
	err := json.Unmarshal(b, &ct.Value)
	ct.Valid = (err == nil)
	return err
}

type testValue struct {
	Name   string
	Number customType
}

type testPointer struct {
	Name   string
	Number *customType
}

func main() {
	n := customType{Valid: true, Value: 1}
	iV := testValue{Name: "TestYo", Number: n}
	iP := testPointer{Name: "TestYo1", Number: &n}
	xV := []testValue{}
	xPP := []*testPointer{}
	xP := []testPointer{}

	iVb, _ := json.Marshal(iV)
	// In this case the method sets is not called because the type is T and the set method expects a *T type
	/*
			Solution

			func (ct customType) MarshalJSON() ([]byte, error) {
			MarshalJSONCalled++
			fmt.Println("MarshalJSON", ct)
			if !ct.Valid {
				return []byte("null"), nil
			}
			return json.Marshal(ct.Value)
		}

	*/
	MarshalJSONCalledExpected++

	fmt.Println("item value Marshal", string(iVb))

	iPb, _ := json.Marshal(iP)
	// In this case the method sets is called because the type is *T
	MarshalJSONCalledExpected++

	fmt.Println("item pointer Marshal", string(iPb))

	xV = append(xV, iV)

	xVb, _ := json.Marshal(xV)
	// In this case the method sets is called because all slice elements underlie a pointer
	MarshalJSONCalledExpected++

	fmt.Println("slice value Marshal", string(xVb))

	xPP = append(xPP, &iP)

	xPPb, _ := json.Marshal(xPP)
	// In this case the method sets is called because the elements of the slice are of type *T
	MarshalJSONCalledExpected++

	fmt.Println("slice pointer of pointer", string(xPPb))

	xP = append(xP, iP)

	xPb, _ := json.Marshal(xP)
	// In this case the method sets is called because the value is of type *T
	MarshalJSONCalledExpected++

	fmt.Println("slice pointer", string(xPb))

	fmt.Printf("---\nMarshal\tExpected: %d\tCalled: %d", MarshalJSONCalledExpected, MarshalJSONCalled)
}

```
**Code** https://go.dev/play/p/68HzW1xk83G