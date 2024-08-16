package ch_9_sync_mutex

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// * Mutex dapat digunakan untuk mengatasi masalah race condition
// * Mutex digunakan untuk melakukan locking maupun unlocking
// * Jadi ketika ada 1 goroutine yang sedang mengakses sebuah data, maka data tersebut akan di lock
// * dan hanya ada satu goroutine yang dapat mengakses data tersebut,
// * sehingga goroutine lainnya akan di block sampai data tersebut di unlock
func TestMutex(t *testing.T) {

	x := 0

	var mutex sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				// * Proses ini akan di lock
				mutex.Lock()
				x += 1
				// * Kemudian akan dibuka kembali untuk proses berikutnya
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Mutex :", x) // akan selalu fix 100000 datanya
}

// ! RACE CONDITION
func TestRaceCondition(t *testing.T) {

	x := 0
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x += 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Race Condition", x) //* Akan bervariasi karena terjadi race condition
}
