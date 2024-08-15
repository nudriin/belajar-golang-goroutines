package ch_2_channel

import (
	"fmt"
	"testing"
	"time"
)

func SendData(data string) string {
	return data
}

func TestCreateChannel(t *testing.T) {
	channel := make(chan string) // * Membuat channel
	defer close(channel)         //* Menutup channel apabila sudah selesai

	// * Mengirim data ke channel menggunakan go routine
	go func() {
		time.Sleep(2 * time.Second)
		// channel <- "Sasuke"
		channel <- SendData("Nurdin")
		fmt.Println("Selesai mengirim data ke channel")
	}()

	// * Mengambil data dari channel
	data := <-channel
	fmt.Println(data)
	// fmt.Println(<-channel)
}

func TestChannelWithoutRecive(t *testing.T) {
	channel := make(chan string) // * Membuat channel
	defer close(channel)         //* Menutup channel apabila sudah selesai

	// * Mengirim data ke channel menggunakan go routine
	go func() {
		time.Sleep(2 * time.Second)
		channel <- SendData("Nurdin")
		fmt.Println("Selesai mengirim data ke channel") // tidak akan di print karena tidak ada penerima
	}()

	time.Sleep(5 * time.Second)

}

// * Akan terjadi error Deadlock
func TestChannelWithoutSender(t *testing.T) {
	channel := make(chan string) // * Membuat channel
	defer close(channel)         //* Menutup channel apabila sudah selesai

	// * data tidak dikirim ke channel
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Selesai mengirim data ke channel")
	}()

	// * Mengambil data dari channel
	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}
