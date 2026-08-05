package main

import (
	"github.com/google/uuid"
	"github.com/victorvbello/gonotes/design-patterns/observer_pattern/stock"
)

func main() {

	voo := stock.NewStock("VOO", 400.0)

	apiObserver := &stock.API{}
	appObserver := &stock.APP{}
	webObserver := &stock.WEB{}

	apiObserverID := stock.ObserverID(uuid.New().String())
	voo.RegisterObserver(apiObserverID, apiObserver)
	appObserverID := stock.ObserverID(uuid.New().String())
	voo.RegisterObserver(appObserverID, appObserver)
	webObserverID := stock.ObserverID(uuid.New().String())
	voo.RegisterObserver(webObserverID, webObserver)
	appObserverID2 := stock.ObserverID(uuid.New().String())
	voo.RegisterObserver(appObserverID2, appObserver)

	voo.UpdatePrice(410.0)

	voo.UnregisterObserver(appObserverID2)

}
