package ch_13_sync_once

import (
	"fmt"
	"sync"
	"testing"
)

// ! digunakan untuk memastikan bahwa function hanya di eksekusi 1 kali saja

var counter int = 0

func OnlyOnceExecute() {
	counter++
}

func TestOnce(t *testing.T) {
	var once sync.Once       // membuat once
	var group sync.WaitGroup // membuat WaitGroup

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			//* menggunakan once, artinya function ini hanya boleh diakses sekali
			// * eksekusi lain pada loop ini tidak akan di hiraukan
			once.Do(OnlyOnceExecute)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println(counter)
}
