// Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// Остановка через закрытие канала
	fmt.Println("Способ 1: Закрытие канала")
	stopChan := make(chan struct{})
	go workerCloseChannel(stopChan)

	time.Sleep(1 * time.Second)
	close(stopChan)
	time.Sleep(1 * time.Second)
	fmt.Printf("\n\n")

	// Остановка через условие
	fmt.Println("Способ 2: Выход по условию")
	flagChan := make(chan bool)
	go workerFlagChannel(flagChan)

	time.Sleep(1 * time.Second)
	flagChan <- true
	time.Sleep(1 * time.Second)
	fmt.Printf("\n\n")

	// Остановка через контекст
	fmt.Println("Способ 3: Выход с использованием контекста")
	ctx, cancel := context.WithCancel(context.Background())
	go workerContext(ctx)

	time.Sleep(1 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
	fmt.Printf("\n\n")

	// Остановка по выполнению
	// (Использование WaitGroup чтобы дождаться выполнения всех горутин)
	fmt.Println("Способ 4: Обычное заваершение")
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go workerWaitGroup(&wg, i)
	}

	time.Sleep(1 * time.Second)
	wg.Wait()
	fmt.Printf("\n\n")

	// Использование runtime.Goexit()
	fmt.Println("Способ 5: Использование runtime.Goexit()")
	go workerRuntimeExit()

	time.Sleep(1 * time.Second)
	fmt.Printf("\n\n")

	// Паника в горутине
	fmt.Println("Способ 6: Использование паники")
	go workerPanic()

	time.Sleep(1 * time.Second)
	fmt.Printf("\n\n")
}

func workerCloseChannel(stopChan chan struct{}) {
	for {
		select {
		case <-stopChan:
			fmt.Println(" - Горутина выполнена   [workerCloseChannel]")
			return
		default:
			fmt.Println(" - Горутина работает... [workerCloseChannel]")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func workerFlagChannel(flagChan chan bool) {
	for {
		select {
		case <-flagChan:
			fmt.Println(" - Горутина выполнена   [workerFlagChannel]")
			return
		default:
			fmt.Println(" - Горутина работает... [workerFlagChannel]")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func workerContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(" - Горутина выполнена   [workerContext]")
			return
		default:
			fmt.Println(" - Горутина работает... [workerContext]")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func workerWaitGroup(wg *sync.WaitGroup, i int) {
	defer fmt.Printf(" - (%d) Горутина выполнена   [workerWaitGroup]\n", i)
	fmt.Printf(" - (%d) Горутина работает... [workerWaitGroup]\n", i)
	time.Sleep(100 * time.Millisecond)
	wg.Done()
}

func workerRuntimeExit() {
	defer fmt.Println(" - Горутина выполнена   [workerRuntimeExit]")
	for {
		fmt.Println(" - Горутина работает... [workerRuntimeExit]")
		time.Sleep(100 * time.Millisecond)
		runtime.Goexit()
	}
}

func workerPanic() {
	defer fmt.Println(" - Горутина выполнена    [workerPanic]")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			fmt.Println(" - Горутина востановлена [workerPanic]")
		}
	}()
	for {
		fmt.Println(" - Горутина работает...  [workerPanic]")
		time.Sleep(100 * time.Millisecond)
		panic(" - Горутина паникует     [workerPanic]")
	}
}
