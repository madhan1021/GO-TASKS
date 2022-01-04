//RW mutex

package main

import (
	"errors"
	"fmt"
	"sync"
)

type Account struct {
	Balance int
	sync.RWMutex
}

func (a *Account) GetBalance() int {
	a.RLock()
	defer a.RUnlock()
	return a.Balance //once the function gets returned we cant go for the next line so we have to use defer for the unlock purpose

}

func (a *Account) Deposit(amount int) {
	a.Lock()
	a.Balance = a.Balance + amount
	a.Unlock()
}

func (a *Account) Withdraw(amount int) error {
	a.Lock()
	defer a.Unlock()
	if a.Balance >= amount {
		a.Balance = a.Balance - amount
		return nil
	}
	return errors.New("insufficient balance")

}

func main() {
	acc := Account{Balance: 1000}
	done := make(chan struct{})
	go func() {
		defer func() { done <- struct{}{} }()

		err := acc.Withdraw(1000)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Successfully withdrawn 10000")

	}()
	go func() {
		defer func() { done <- struct{}{} }()

		err := acc.Withdraw(2000)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Successfully withdrawn 2000")
	}()
	<-done
	<-done
}
