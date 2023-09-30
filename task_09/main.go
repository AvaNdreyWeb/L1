// Разработать конвейер чисел. Даны два канала: в первый пишутся числа
// (x) из массива, во второй — результат операции x*2, после чего данные
// из второго канала должны выводиться в stdout.

package main

import (
	"fmt"
	"sync"
)

const N = 100

var wg = sync.WaitGroup{}

func main() {
	data := [N]int{}
	for i := 0; i < N; i++ {
		data[i] = 1 + i
	}

	inCh := make(chan int)  // Канал для передачи данных в "Удвоитель"
	outCh := make(chan int) // Канал для передачи данных на вывод

	wg.Add(3)
	go Sender(inCh, data)   // Чтение данных из массива и передача в "Удвоитель"
	go Handler(inCh, outCh) // Получение чисел, их удвоение и отправка на вывод
	go Reciver(outCh)       // Получение чисел и их вывод в stdout
	wg.Wait()
}

func Sender(inCh chan<- int, arr [N]int) {
	for _, x := range arr {
		inCh <- x
	}
	close(inCh)
	wg.Done()
}

func Handler(inCh <-chan int, outCh chan<- int) {
	for x := range inCh {
		outCh <- 2 * x
	}
	close(outCh)
	wg.Done()
}

func Reciver(outCh <-chan int) {
	for x := range outCh {
		fmt.Println(x)
	}
	wg.Done()
}
