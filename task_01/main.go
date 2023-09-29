// Дана структура Human (с произвольным набором полей и методов).

// Реализовать встраивание методов в структуре Action
// от родительской структуры Human (аналог наследования).

package main

import "fmt"

type Human struct {
	Name    string
	Surname string
	Age     uint
}

func (h *Human) AboutString() string {
	return fmt.Sprintf(
		"{\n  Имя: %s\n  Фамилия: %s\n  Возраст: %d\n}",
		h.Name,
		h.Surname,
		h.Age,
	)
}

type Action struct {
	Human // Встраивание структуры Human
}

func main() {
	a := Action{Human{Name: "Андрей"}} // Поля можно инициализировать сразу
	a.Human.Surname = "Алейник"        // Вот так можно обратиться к полю
	a.Age = 24                         // Но лучше вот так

	s := a.AboutString() // Вызов метода структуры Human
	fmt.Println(s)
}
