// Разработать программу, которая перемножает, делит, складывает, вычитает
// две числовых переменных a,b, значение которых > 2^20.

package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	// Для решения этой задачи есть пакет big
	num1 := big.NewInt(int64(math.Pow(2, 20)) + int64(math.Pow(2, 10)))
	num2 := big.NewInt(int64(math.Pow(2, 21)) - int64(math.Pow(2, 10)))

	sum := new(big.Int).Add(num1, num2)
	sub := new(big.Int).Sub(num1, num2)
	mul := new(big.Int).Mul(num1, num2)

	num1f := new(big.Float).SetInt(num1)
	num2f := new(big.Float).SetInt(num2)

	quof := new(big.Float).Quo(num1f, num2f)

	fmt.Printf("%s + %s = %s\n", num1.String(), num2.String(), sum.String())
	fmt.Printf("%s - %s = %s\n", num1.String(), num2.String(), sub.String())
	fmt.Printf("%s * %s = %s\n", num1.String(), num2.String(), mul.String())
	fmt.Printf("%s / %s = %s\n", num1f.String(), num2f.String(), quof.String())
}
