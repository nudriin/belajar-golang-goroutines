package ch_5_buffered_channel

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateBufferChannel(t *testing.T) {

	// ! Membuat buffered channel
	channel := make(chan string, 4) // 4 adalah jumlah buffernya
	defer close(channel)

	go func() {
		time.Sleep(4 * time.Second)
		channel <- "Nurdin"
		channel <- "Hishasy"
		channel <- "Sunny"
	}()

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)

	fmt.Println(len(channel)) // melihat jumlah data pada buffer
	fmt.Println(cap(channel)) // melihat kapasitas buffer
}
