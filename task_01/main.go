package main

import "fmt"

type Human struct {
	Name string
	Age  uint
}

type Action struct {
	Human
}

func (h *Human) AboutString() string {
	return fmt.Sprintf("Name: %s, Age: %d", h.Name, h.Age)
}

func main() {
	a := Action{Human{Name: "Andrey", Age: 24}}
	s := a.AboutString()
	fmt.Println(s)
}
