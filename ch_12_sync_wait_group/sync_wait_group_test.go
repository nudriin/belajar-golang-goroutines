package ch_12_sync_wait_group

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// ! Wait group digunakan untuk menunggu sebuah proses goroutine selesai sebelum aplikasi selesai
// ! sama sepert sleep yang sudah dilakukan pada contoh sebelumnya
// ! wait group lebih direkomndasikan

func RunAsync(group *sync.WaitGroup) {
	defer group.Done() //! untuk menandai bahwa sudah selesai

	group.Add(1) // ! untuk menandai ada proses goroutine

	fmt.Println("Sleep")
	time.Sleep(2 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsync(group)
	}

	group.Wait() // ! menunggu sampai selesai
	fmt.Println("Complete")
}
