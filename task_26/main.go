// Разработать программу, которая проверяет, что все символы в строке
// уникальные (true — если уникальные, false etc). Функция проверки должна
// быть регистронезависимой.

// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("abcd — %v\n", IsUniqSybols("abcd"))
	fmt.Printf("abCdefAaf — %v\n", IsUniqSybols("abCdefAaf"))
	fmt.Printf("aabcd — %v\n", IsUniqSybols("aabcd"))
}

func IsUniqSybols(s string) bool {
	low := strings.ToLower(s) // Приводим строку к нижнему регистру
	// Будем добавлять символы строки в map до тех пор пока не кончится строка
	// или мы не найдём повторяющийся символ
	var check = make(map[rune]struct{})
	for _, v := range low {
		if _, ok := check[v]; ok {
			return false
		} else {
			check[v] = struct{}{}
		}
	}
	return true
}
