// Реализовать быструю сортировку массива (quicksort)
// встроенными методами языка.

package main

import "fmt"

const N = 10

func main() {
	arr := [N]int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}
	fmt.Println("Исходный массив:", arr)

	quickSort(&arr, 0, len(arr)-1)
	fmt.Println("Отсортированный массив:", arr)
}

func quickSort(arr *[N]int, l, r int) {
	// Массив из одного или нуля элементов считаем отсортированным
	if l < r {
		pi := partition(arr, l, r)
		// Поскольку опорный элемент уже отсортирован, осталось отсортировать
		// левую и правую часть соответственно
		quickSort(arr, l, pi-1)
		quickSort(arr, pi+1, r)
	}
}

// Выбираем опорный элемент и разделяем массив на две части.
// Слева от опорного будут расположаны элементы, которые меньше или равны
// опорному, а справа от него большие
func partition(arr *[N]int, l, r int) int {
	p := arr[r]
	i := l - 1

	for j := l; j < r; j++ {
		if arr[j] <= p {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[r] = arr[r], arr[i+1]
	return i + 1
}
