// Разработать программу, которая будет последовательно отправлять значения
// в канал, а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.

package main

import (
	"context"
	"flag"
	"fmt"
	"sync"
	"time"
)

func main() {
	var N = flag.Int("N", 2, "Время работы программы в секундах")
	flag.Parse()

	done := make(chan bool)

	// Горутина, отсчитывающая N секунд
	go func(done chan<- bool, n int) {
		time.Sleep(time.Duration(n) * time.Second)
		done <- true
	}(done, *N)

	msgs := make(chan interface{})
	ctx, cancel := context.WithCancel(context.Background())
	go Writer(ctx, msgs) // Горутина, которая пишет в канал

	var wg sync.WaitGroup
	wg.Add(1)
	go Reader(&wg, msgs) // Горутина, которая читает канал

	<-done
	cancel()  // Завершаем запись в канал
	wg.Wait() // Если нужно максимально точное время, то надо убрать, но возможна потеря данных
}

func Reader(wg *sync.WaitGroup, ch <-chan interface{}) {
	for x := range ch {
		fmt.Printf("Прочитано: %v\n", x)
	}
	wg.Done()
}

func Writer(ctx context.Context, ch chan<- interface{}) {
	for {
		select {
		case <-ctx.Done():
			close(ch)
			return

		default:
			ch <- 42
		}
	}
}
