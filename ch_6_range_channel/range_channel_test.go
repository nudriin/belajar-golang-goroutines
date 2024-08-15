package ch_6_range_channel

import (
	"fmt"
	"strconv"
	"testing"
)

func TestRangeChannel(t *testing.T) {

	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Data ke-" + strconv.Itoa(i)
		}
		// ! jika tidak di close akan error
		close(channel) //! menutup channel
	}()

	// ! for range nya seperti ini untuk mengambil data
	for data := range channel {
		fmt.Println(data) //! perulangan akan berhenti ketika channel ditutup (close)
	}
}
