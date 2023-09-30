// Реализовать структуру-счетчик, которая будет инкрементироваться
// в конкурентной среде. По завершению программа должна выводить
// итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	rwmu sync.RWMutex
	val  uint
}

func (c *Counter) Inc() {
	c.rwmu.Lock()   // Закрываем для записи
	c.val += 1      // Меняем значение
	c.rwmu.Unlock() // Открываем для записи
}

func (c *Counter) Get() uint {
	c.rwmu.RLock()         // Закрываем для чтения
	defer c.rwmu.RUnlock() // Открываем для чтения после возврата значения
	return c.val           // Возвращаем значение
}

const N = 10 // Количество горутин

var wg sync.WaitGroup

func main() {
	c := Counter{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(c *Counter) {
			c.Inc()
			fmt.Printf("Cчётчик: %d\n", c.Get())
			wg.Done()
		}(&c)
	}
	wg.Wait()
	fmt.Printf("Итого: %d\n", c.Get())
}
