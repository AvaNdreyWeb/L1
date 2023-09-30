// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "snow dog sun"
	fmt.Printf("%s   <->   %s\n", s, ReverseWords(s))
}

func ReverseWords(s string) string {
	res := strings.Split(s, " ")
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return strings.Join(res, " ")
}
