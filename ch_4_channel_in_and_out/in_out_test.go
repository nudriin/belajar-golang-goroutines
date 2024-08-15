package ch_4_channel_in_and_out

import (
	"fmt"
	"testing"
	"time"
)

// ! function ini hanya dapat digunakan untuk memasukan data ke channel
func OnlyInChannel(channel chan<- string, data string) {
	time.Sleep(2 * time.Second)
	channel <- data
}

// ! function ini hanya dapat digunakan untuk mengambil data dari channel
func OnlyOutChannel(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOut(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyInChannel(channel, "Nurdin")
	go OnlyOutChannel(channel)

	time.Sleep(3 * time.Second)
}
