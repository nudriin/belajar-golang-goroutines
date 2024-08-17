package ch_17_atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var counter int64 = 0
	var group sync.WaitGroup

	for i := 0; i < 1000; i++ {
		go func() {
			group.Add(1)
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&counter, 1) // menggunakan atomic agar tidak terjadi race condition
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println(counter)
}
