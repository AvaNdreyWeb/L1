// Поменять местами два числа без создания временной переменной.

package main

import "fmt"

func main() {

	var a, b int
	// Способ I
	a, b = 42, -18
	fmt.Printf("До обмена: a = %d, b = %d\n", a, b)
	a, b = b, a
	fmt.Printf("I: a = %d, b = %d\n", a, b)

	// Способ II
	a, b = 314, 27
	fmt.Printf("\nДо обмена: a = %d, b = %d\n", a, b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("II: a = %d, b = %d\n", a, b)
}
