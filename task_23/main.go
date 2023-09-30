// Удалить i-ый элемент из слайса.

package main

import "fmt"

var i = 2

func main() {
	a := []int{0, 10, 20, 30, 40, 50}
	fmt.Print("Исходный слайс: ")
	fmt.Println(a)
	a, err := DeleteByIndex(a, i)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Удаляем a[%d]: ", i)
		fmt.Println(a)
	}
}

func DeleteByIndex(a []int, i int) ([]int, error) {
	if (i < 0) || (len(a) <= i) {
		return a, fmt.Errorf("Неправильный индекс %d", i)
	}
	// Вернём новый срез, из элементов с индексом меньше и больше i
	return append(a[:i], a[i+1:]...), nil
}
