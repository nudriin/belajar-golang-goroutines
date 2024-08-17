package ch_15_sync_map

// * Mirip map, namun aman untuk concurencies

import (
	"fmt"
	"sync"
	"testing"
)

func TestSyncMap(t *testing.T) {
	var maps sync.Map
	var group sync.WaitGroup

	AddToMap := func(mapS *sync.Map, value int, groupS *sync.WaitGroup) {
		defer groupS.Done()

		groupS.Add(1)
		mapS.Store(value, value) // * memasukan data ke maps emnggunakan method store
	}

	for i := 0; i < 100; i++ {
		go AddToMap(&maps, i, &group)
	}

	group.Wait()

	// * Membaca data yang ada pada map
	maps.Range(func(key, value any) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
