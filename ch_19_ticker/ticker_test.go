package ch_19_ticker

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// ! Seperti setInterval di js

func TestTicker(t *testing.T) {
	var group sync.WaitGroup
	group.Add(1)
	ticker := time.NewTicker(1 * time.Second) // akan di kirim setiap 1 detik

	// !akan kita stop tickernya setelah 5 detik menggunakan time After
	_ = time.AfterFunc(5*time.Second, func() {
		ticker.Stop() //! Mengehntikan ticker
		group.Done()
	})

	for data := range ticker.C {
		fmt.Println(data)
	}

	group.Wait()
}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second) // akan di kirim setiap 1 detik

	for data := range channel {
		fmt.Println(data)
	}

}
