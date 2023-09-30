// Реализовать бинарный поиск встроенными методами языка.

package main

import "fmt"

const x = 4
const y = 6

func main() {
	a := []int{-10, 1, 2, 3, 5, 6, 7, 8, 9}
	fmt.Printf("Слайс: %v\n", a)

	if k, v := BinSearch(a, x); v {
		fmt.Printf("a[%d] = %v [BinSearch]\n", k, x)
	} else {
		fmt.Printf("%d - Отсутствует [BinSearch]\n", x)
	}

	if k, v := RecursiveBinSearch(a, y); v {
		fmt.Printf("a[%d] = %v [RecursiveBinSearch]\n", k, y)
	} else {
		fmt.Printf("%d - Отсутствует [RecursiveBinSearch]\n", y)
	}
}

// Бинарный поиск позволяет, на каждой итерации отбросить половину непроверенных
// значений в силу того, что исходные данные отсортированы. Сложность O(logN)
func BinSearch(a []int, x int) (int, bool) {
	var l, r, m int
	l = 0
	r = len(a) - 1

	for l <= r {
		m = (l + r) / 2
		if a[m] == x {
			return m, true
		} else if a[m] < x {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1, false
}

func RecursiveBinSearch(a []int, x int) (int, bool) {
	var (
		l int = 0
		r int = len(a) - 1
		m int = (l + r) / 2
	)
	if r < l {
		return -1, false
	}
	if a[m] == x {
		return m, true
	}
	if x < a[m] {
		return RecursiveBinSearch(a[l:m], x)
	} else {
		k, v := RecursiveBinSearch(a[m+1:r+1], x)
		return m + 1 + k, v
	}
}
