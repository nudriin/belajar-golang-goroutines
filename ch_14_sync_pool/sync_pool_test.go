package ch_14_sync_pool

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// * Design pattern object pool digunakan untuk menyimpan data
// * dan dapat mengambil data dari poolnya
// * biasa digunakan untuk manage koneksi ke database

func TestPool(t *testing.T) {
	var pool sync.Pool
	var group sync.WaitGroup

	poolWithDefault := sync.Pool{
		New: func() interface{} {
			return "Default data"
		},
	}

	// * memasukan data ke pool
	pool.Put("Nurdin")
	pool.Put("Naruto")
	pool.Put("Gaara")

	for i := 0; i < 10; i++ {
		go func() {
			group.Add(1)
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data) // * mengembalikan lagi datanya agar tidak hilang dari pool
			group.Done()
		}()
	}

	// * memasukan data ke pool
	poolWithDefault.Put("Nurdin")
	poolWithDefault.Put("Naruto")
	poolWithDefault.Put("Gaara")

	for i := 0; i < 10; i++ {
		go func() {
			data := poolWithDefault.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			poolWithDefault.Put(data) // * mengembalikan lagi datanya agar tidak hilang dari pool
		}()
	}

	time.Sleep(11 * time.Second)

}
