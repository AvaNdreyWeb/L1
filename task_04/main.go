// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные
// из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.

// Программа должна завершаться по нажатию Ctrl+C.
// Выбрать и обосновать способ завершения работы всех воркеров.

package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func main() {
	// go run main.go -N=100
	var N = flag.Int("N", 5, "Количество воркеров")
	flag.Parse()

	msgs := make(chan interface{})                          // Канал для записи произвольных данных
	ctx, cancel := context.WithCancel(context.Background()) // Контекст с отменой для окончания публикации данных в канал
	go Publisher(ctx, msgs)                                 // Горутина публикующая данные в канал

	var wg sync.WaitGroup
	wg.Add(*N)
	for i := 0; i < *N; i++ {
		go Worker(&wg, msgs, i) // Воркеры, читающие данные из канала
	}

	// Ждём получения сигнала Ctrl + C
	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, os.Interrupt)
	<-signalCh

	cancel()  // Останавливаем запись в канал
	wg.Wait() // Дожидаемся выполнения воркеров
}

func Worker(wg *sync.WaitGroup, ch <-chan interface{}, i int) {
	// Читаем данные из канала пока он открыт
	for x := range ch {
		fmt.Printf("Воркер %d: %v\n", i, x)
	}
	wg.Done()
}

func Publisher(ctx context.Context, ch chan<- interface{}) {
	data := []interface{}{42, true, "Привет", 3.14}
	l := len(data)
	i := 0
	for {
		select {
		case <-ctx.Done(): // В случае отмены контекста
			close(ch) // Закрываем канал
			return    // Выходим

		default: // По умолчанию
			ch <- data[i]   // Публикуем данные в канал
			i = (i + 1) % l // Цикличный обход индексов слайса data
		}
	}
}
