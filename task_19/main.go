// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

package main

import "fmt"

func main() {
	s := "你好, Привет, Hello! ⚙"
	fmt.Printf("%s   <->   %s\n", s, Reverse(s))
}

func Reverse(s string) string {
	res := []rune(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return string(res)
}
