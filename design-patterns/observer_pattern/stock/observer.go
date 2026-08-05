package stock

import "fmt"

type ObserverID string

type Observer interface {
	Update()
}

type API struct{}

func (a *API) Update() {
	fmt.Println("API: Stock was updated")
}

type APP struct{}

func (a *APP) Update() {
	fmt.Println("APP: Stock was updated")
}

type WEB struct{}

func (a *WEB) Update() {
	fmt.Println("WEB: Stock was updated")
}
