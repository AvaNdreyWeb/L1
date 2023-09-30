// К каким негативным последствиям может привести данный фрагмент кода,
// и как это исправить? Приведите корректный пример реализации.

// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

// func main() {
//   someFunc()
// }

// Проблема 1. Здесь происходит копирование большой строки
// Проблема 2. Отсутствует проверка длины строки при взятии слайса
package main

import (
	"fmt"
	"strings"
)

var justString string

func buildHugeString(size int) *strings.Builder {
	var builder strings.Builder

	for i := 0; i < size; i++ {
		builder.WriteString("X")
	}

	return &builder // Решение 1. Возвращаем не строку а её билдер
}

func someFunc() {
	builder := buildHugeString(1 << 16)
	hugeString := builder.String() // Решение 1. Копирование отсутствует
	if len(hugeString) > 100 {     // Решение 2. Проверяем длину строки
		justString = hugeString[:100]
	} else {
		justString = hugeString
	}
}

func main() {
	someFunc()
	fmt.Println(justString)
}
