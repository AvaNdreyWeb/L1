// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными
// параметрами x,y и конструктором.

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// |AB| =  √((Ax - Bx)^2 + (Ay - By)^2)
func (a *Point) GetDistance(b *Point) float64 {
	x := a.x - b.x
	y := a.y - b.y
	return math.Sqrt(x*x + y*y)
}

func main() {
	A := NewPoint(0, 0)
	B := NewPoint(6, 8)
	dist := A.GetDistance(B)
	fmt.Printf("Расстояние: %.2f\n", dist)
}
