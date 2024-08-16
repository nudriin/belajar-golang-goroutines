package ch_7_select_channel

import (
	"fmt"
	"testing"
	"time"
)

// * DIGUNAKAN UNTUK MENGAMBIL DATA DARI BEBERAPA CHANNEL SECARA SEKALIGUS

func TestSelectChannel(t *testing.T) {

	channel1 := make(chan string)
	channel2 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		channel1 <- "Nurdin"
		channel2 <- "Hishasy"
	}()

	counter := 0

	for {
		// ! menggunakan select untuk mengambil data data dari beberapa channel
		select {
		case data := <-channel1:
			fmt.Println(data)
			counter++
		case data := <-channel2:
			fmt.Println(data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}
