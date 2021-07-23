package main

import (
	"errors"
	"fmt"
	"sync"
)

type Account struct {
	balance uint
	mu      sync.Mutex
}

func (account *Account) deposit(wg *sync.WaitGroup, amount uint) {
	defer wg.Done()
	account.mu.Lock()
	account.balance += amount
	account.getBalance()
	account.mu.Unlock()
}

func (account *Account) withdraw(wg *sync.WaitGroup, amount uint, err chan<- error) {
	defer func() {
		account.mu.Unlock()
		wg.Done()
	}()
	account.mu.Lock()
	if account.balance < amount {
		err <- errors.New(fmt.Sprintf("Withdrawl amount %d exceeds balance %d \n", amount, account.balance))
		return
	} else {
		account.balance -= amount
		account.getBalance()
	}
}

func (account *Account) getBalance() uint {
	fmt.Printf("Current Balance is %d \n", account.balance)
	return account.balance
}

func main() {
	var wg sync.WaitGroup
	account := Account{}
	account.balance = 15000
	err := make(chan error, 2)

	wg.Add(4)

	go account.deposit(&wg, 500)
	go account.deposit(&wg, 10000)
	go account.withdraw(&wg, 10000, err)
	go account.withdraw(&wg, 100, err)
	wg.Wait()
	close(err)
	for er := range err {
		fmt.Println(er)
	}

	account.getBalance()
}
