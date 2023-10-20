package main

import (
	"fmt"
	"sync"
)

type account struct {
	balance int
	mutex   sync.Mutex
}

func (a *account) deposit(amount int, chnl chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.balance += amount
	chnl <- true

}
func (a *account) withdraw(amount int, chnl chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if a.balance-amount >= 0 {
		a.balance -= amount
		chnl <- true
	} else {
		chnl <- false
	}
}
func main() {
	acc := account{balance: 0}
	chnl := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(4)
	go acc.deposit(100, chnl, &wg)
	go acc.withdraw(200, chnl, &wg)
	go acc.deposit(500, chnl, &wg)
	go acc.withdraw(400, chnl, &wg)
	go func() {
		wg.Wait()
		close(chnl)
	}()
	for x := range chnl {
		if x {
			fmt.Println("Successfull")
		} else {
			fmt.Println("Failed")
		}
	}
	fmt.Println("Final amount", acc.balance)

}
