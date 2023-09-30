// Реализовать пересечение двух неупорядоченных множеств.

package main

import "fmt"

type Set map[int]struct{} // Множество удобно реализовать через map

func (s Set) Print() {
	if len(s) == 0 {
		fmt.Printf("{}\n")
		return
	}
	fmt.Print("{")
	for k := range s {
		fmt.Printf("%d, ", k)
	}
	fmt.Print("\b\b}\n")
}

func main() {
	A := Set{
		1: struct{}{},
		3: struct{}{},
		5: struct{}{},
		7: struct{}{},
	}
	B := Set{
		7: struct{}{},
		2: struct{}{},
		5: struct{}{},
		4: struct{}{},
	}
	C := Intersection(A, B)

	fmt.Print("    A: ")
	A.Print()
	fmt.Print("    B: ")
	B.Print()
	fmt.Print("A & B: ")
	C.Print()
}

func Intersection(s1, s2 Set) Set {
	si := Set{}
	// Проходимся по ключам одного множества и проверяем, есть ли они в другом.
	// Если есть, то добавляем ключ в новое множество
	for k := range s2 {
		if _, ok := s1[k]; ok {
			si[k] = struct{}{}
		}
	}
	return si
}
