package ch_3_channel_as_param

import (
	"fmt"
	"testing"
	"time"
)

// ! Tanpa perlu menulsikan pointer, secara default channel menggunakan pointer
func ParamChannel(channel chan []string, data []string) {
	time.Sleep(2 * time.Second)
	channel <- data
}

func TestParamChannel(t *testing.T) {
	channel := make(chan []string)
	defer close(channel)

	//! pass by reference
	go ParamChannel(channel, []string{"Nurdin", "Hishasy", "Sunny", "Summer"})

	data := <-channel
	fmt.Println(data)
}
