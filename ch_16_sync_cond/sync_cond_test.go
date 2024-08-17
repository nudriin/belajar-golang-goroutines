package ch_16_sync_cond

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group sync.WaitGroup

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock() // melakukan lock
	cond.Wait()   // untuk pngondisian
	fmt.Println("Done", value)
	cond.L.Unlock() // membuka lock
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	// ! Mnegirim sinyal agar bisa lanjut dan tidak perlu wait lagi untuk cond
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		time.Sleep(1 * time.Second)
	// 		cond.Broadcast() //! Bisa juga gini
	// 	}
	// }()

	group.Wait()
}
