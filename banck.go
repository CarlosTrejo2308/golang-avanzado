package main

import (
	"fmt"
	"sync"
)

var balance int

func Deposit(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	balance = balance + amount
	fmt.Println("Deposited", amount)
}

func Withdraw(amount int, wg *sync.WaitGroup) bool {
	defer wg.Done()

	// Wait a second to simulate a delay
	//time.Sleep(time.Millisecond)

	if amount > balance {
		fmt.Println("Insufficient funds")
		return false
	}
	balance = balance - amount
	return true
}

func Balance() int {
	return balance
}

func main() {
	var wg sync.WaitGroup

	balance = 500
	wg.Add(2)
	go Deposit(200, &wg)
	go Withdraw(700, &wg)

	wg.Wait()
	fmt.Println(Balance())
}
