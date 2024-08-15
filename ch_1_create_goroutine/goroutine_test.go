package ch_1_create_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func SayHello() {
	fmt.Println("Hello")
}

func TestSayHello(t *testing.T) {
	// * Cukup dengan mnambahkan "go" sebelum memanggil function
	// function ini akan menjadi async
	go SayHello() // ?menjalankan function menggunakan goroutine
	fmt.Println("test")

	// * digunakan untuk menunggu 1 detik agar goroutine tidak langsung di close ketika program selesai
	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Printf("Number %d\n", number)
}

func TestDisplayNumber(t *testing.T) {
	for i := 0; i < 100_000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(15 * time.Second)
}
