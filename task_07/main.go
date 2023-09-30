// Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const N = 15

var r = rand.New(rand.NewSource(time.Now().UnixNano()))
var wg sync.WaitGroup
var rwmu sync.RWMutex

func main() {
	var (
		dataMap     = map[int]string{}
		dataSyncMap sync.Map
	)

	wg.Add(2 * N)
	for i := 0; i < N; i++ {
		go WriterMap(dataMap, i)            // Тут запись в обычную map с использованием мьютекса
		go WriterSyncMap(&dataSyncMap, N+i) // Тут запись в потокобезопасную sync.Map
	}
	wg.Wait()

	fmt.Println("Обычная map:")
	for k, v := range dataMap {
		fmt.Printf("  %d: %s\n", k, v)
	}

	fmt.Printf("\nsync.Map:\n")
	dataSyncMap.Range(func(k, v interface{}) bool {
		fmt.Printf("  %d: %s\n", k, v)
		return true
	})
}

func WriterMap(data map[int]string, i int) {
	key := r.Intn(10)
	rwmu.Lock()
	data[key] = fmt.Sprintf("Горутина %d", i)
	rwmu.Unlock()
	wg.Done()
}

func WriterSyncMap(data *sync.Map, i int) {
	key := r.Intn(10)
	data.Store(key, fmt.Sprintf("Горутина %d", i))
	wg.Done()
}
