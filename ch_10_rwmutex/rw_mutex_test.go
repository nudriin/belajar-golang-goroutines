package ch_10_rwmutex

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// * Rw mutex adalah read write mutex

type BankAccount struct {
	RWMtx    sync.RWMutex
	Ballance int
}

// * Mutex untuk write
func (acc *BankAccount) AddBalance(amount int) {
	acc.RWMtx.Lock()
	acc.Ballance += amount
	acc.RWMtx.Unlock()
}

// * Mutex untuk read
func (acc *BankAccount) GetBalance() int {
	acc.RWMtx.RLock() //! RLock untuk read
	ballance := acc.Ballance
	acc.RWMtx.RUnlock() // ! RUnlock untuk read
	return ballance
}

func TestRWMutex(t *testing.T) {

	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)

	fmt.Println("Final Ballance:", account.GetBalance())
}
