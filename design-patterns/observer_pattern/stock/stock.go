package stock

import "fmt"

type Subject interface {
	RegisterObserver(id ObserverID, observer Observer)
	UnregisterObserver(id ObserverID)
	NotifyObservers()
}

type Stock struct {
	Symbol    string
	Price     float64
	observers map[ObserverID]Observer
}

func NewStock(symbol string, price float64) *Stock {
	return &Stock{Symbol: symbol, Price: price, observers: make(map[ObserverID]Observer)}
}

func (s *Stock) UpdatePrice(price float64) {
	s.Price = price
	s.NotifyObservers()
}

func (s *Stock) RegisterObserver(id ObserverID, observer Observer) {
	fmt.Println("Registering observer:", id)
	s.observers[id] = observer
}

func (s *Stock) UnregisterObserver(id ObserverID) {
	fmt.Println("Unregistering observer:", id)
	delete(s.observers, id)
}

func (s *Stock) NotifyObservers() {
	fmt.Println("Notifying observers")
	for _, observer := range s.observers {
		observer.Update()
	}
}
