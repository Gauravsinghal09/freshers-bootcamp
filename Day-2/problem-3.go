package main

import (
	"fmt"
	"sync"
)

type Account struct {
	balance uint
	mu      sync.Mutex
	wg      sync.WaitGroup
}

func (account *Account) deposit(amount uint) {
	defer account.wg.Done()
	account.mu.Lock()
	account.balance += amount
	account.getBalance()
	account.mu.Unlock()
}

func (account *Account) withdraw(amount uint) {
	defer account.wg.Done()
	account.mu.Lock()
	if account.balance < amount {
		fmt.Printf("Withdrawl amount %d exceeds balance %d \n", amount, account.balance)
		return
	}
	account.balance -= amount
	account.getBalance()
	account.mu.Unlock()
}

func (account *Account) getBalance() uint {
	fmt.Printf("Current Balance is %d \n", account.balance)
	return account.balance
}

func main() {

	account := Account{}

	account.wg.Add(4)
	go account.deposit(500)
	go account.deposit(10000)
	go account.withdraw(10000)
	go account.withdraw(100)
	account.wg.Wait()

	account.getBalance()
}
