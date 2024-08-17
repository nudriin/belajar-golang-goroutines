package ch_18_timer

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// ! Seperti setTimeout di js
func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now()) // di print duluan

	time := <-timer.C
	fmt.Println(time) // data akan masuk setelah 5 detik
}

func TestAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now()) // di print duluan

	time := <-channel
	fmt.Println(time) // data akan dikirim setelah 5 detik

}

func TestAfterFunction(t *testing.T) {
	var group sync.WaitGroup
	group.Add(1)

	time.AfterFunc(5*time.Second, func() {
		fmt.Println(time.Now())
		fmt.Println("Execute after 5 second")
		group.Done()
	})
	fmt.Println(time.Now()) // akan di print dulu

	group.Wait()
}
