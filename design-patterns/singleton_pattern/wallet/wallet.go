package wallet

import (
	"errors"
	"fmt"
	"sync"
)

type walletUnique struct {
	balance float32
}

var currentWallet *walletUnique
var once sync.Once

func createInstance() *walletUnique {
	return new(walletUnique)
}

func GetCurrentWallet() *walletUnique {
	once.Do(func() {
		currentWallet = createInstance()
		fmt.Println("Instance of walletUnique create")
	})
	fmt.Println("Current instance returned")
	return currentWallet
}

func (w *walletUnique) Deposit(v float32) float32 {
	w.balance += v
	return w.balance
}

func (w *walletUnique) Transfer(v float32) (float32, error) {
	if v > w.balance {
		return w.balance, errors.New("insufficient funds")
	}
	w.balance -= v
	return w.balance, nil
}

func (w *walletUnique) Balance() float32 {
	return w.balance
}
