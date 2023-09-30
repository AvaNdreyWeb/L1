// Реализовать собственную функцию sleep.

package main

import (
	"fmt"
	"time"
)

const N = 5

func main() {
	fmt.Printf("Ждём %d сек...\n", N)
	Sleep(time.Duration(N) * time.Second)
	fmt.Println("Готово!")
}

func Sleep(duration time.Duration) {
	// Функция After ждёт помежуток времени, а затем посылает текущее время в
	// возващаемый канал, а в нашей реализации Sleep, мы просто ждём это значение
	<-time.After(duration)
}
