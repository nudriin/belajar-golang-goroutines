package ch_8_default_select

import (
	"fmt"
	"testing"
	"time"
)

// * DIGUNAKAN UNTUK MENGAMBIL DATA DARI BEBERAPA CHANNEL SECARA SEKALIGUS

func TestDefaultSelectChannel(t *testing.T) {

	channel1 := make(chan string)
	channel2 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		channel1 <- "Nurdin"
		channel2 <- "Hishasy"
	}()

	counter := 0

	for {
		// * apabila datanya tidak ada, namun dicoba untuk diambil
		// * maka akan terjadi sebuah error yaitu deadlock
		// * untuk itu kita bisa tambahkan default, jika tidak ada datanya, maka akan di eksekusi defaultnya
		select {
		case data := <-channel1:
			fmt.Println(data)
			counter++
		case data := <-channel2:
			fmt.Println(data)
			counter++
		default:
			fmt.Println("Menunggu data")
		}
		if counter == 2 {
			break
		}
	}
}
