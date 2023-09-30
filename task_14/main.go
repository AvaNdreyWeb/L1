// Разработать программу, которая в рантайме способна определить тип
// переменной: int, string, bool, channel из переменной типа interface{}.

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var (
		xInt   int     = 42
		xStr   string  = "Hello, World!"
		xBool  bool    = true
		xFloat float32 = 3.14
	)
	xChan := make(chan int)

	PrintType(xStr)
	PrintType(xChan)
	PrintType(xFloat)
	PrintType(xBool)
	PrintType(xInt)
}

func PrintType(x interface{}) {
	// Для определения канала есть пакет reflect
	if reflect.TypeOf(x).Kind() == reflect.Chan {
		fmt.Printf("%v - Канал\n", x)
		return
	}

	// Для всего остального есть switch x.(type)
	switch x.(type) {
	case int:
		fmt.Printf("%v - Целое число\n", x)
	case string:
		fmt.Printf("%v - Строка\n", x)
	case bool:
		fmt.Printf("%v - Логический тип\n", x)
	default:
		fmt.Printf("%v - Неизвестный тип\n", x)
	}
}
