package main

import (
	"fmt"
	"gonotes/design-patterns/singleton_pattern/wallet"
)

func main() {
	myWallet := wallet.GetcurrentWallet()
	fmt.Printf("myWallet -- Pointer %p\n", myWallet)
	fmt.Printf("myWallet -- Balance (%.2f)\n", myWallet.Balance())

	var t float32 = 10.0
	fmt.Printf("myWallet -- Transfer 1 (%.2f)\n", t)
	_, err := myWallet.Transfer(t)
	if err != nil {
		fmt.Println("myWallet -- Error on Transfer", err)
	}

	fmt.Printf("myWallet -- Balance (%.2f)\n", myWallet.Deposit(20.10))

	var t2 float32 = 5.0
	fmt.Printf("myWallet -- Transfer 2 (%.2f)\n", t2)
	_, err2 := myWallet.Transfer(t2)
	if err2 != nil {
		fmt.Println("myWallet -- Error on Transfer", err2)
	}

	fmt.Printf("myWallet -- Balance (%.2f)\n", myWallet.Balance())

	fmt.Println("................")

	myWallet2 := wallet.GetcurrentWallet()
	fmt.Printf("myWallet2 -- Pointer %p\n", myWallet2)
	fmt.Printf("myWallet2 -- Balance (%.2f)\n", myWallet2.Balance())
}
