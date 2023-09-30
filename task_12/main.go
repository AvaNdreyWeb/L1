// Имеется последовательность строк - (cat, cat, dog, cat, tree)
// создать для нее собственное множество.

package main

import "fmt"

type Set map[string]struct{}

func NewSet(args ...string) Set {
	s := Set{}
	// Поскольку ключи в map уникальны, просто пройдёмся по последовательности
	// и будем использовать значения, как ключи
	for _, v := range args {
		s[v] = struct{}{}
	}
	return s
}

func (s Set) Print() {
	if len(s) == 0 {
		fmt.Printf("{}\n")
		return
	}
	fmt.Print("{")
	for k := range s {
		fmt.Printf("%s, ", k)
	}
	fmt.Print("\b\b}\n")
}

func main() {
	s := NewSet("cat", "cat", "dog", "cat", "tree")
	fmt.Println("Последовательность:", "cat", "cat", "dog", "cat", "tree")
	fmt.Print("Множество: ")
	s.Print()
}
