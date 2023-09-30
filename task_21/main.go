// Реализовать паттерн «адаптер» на любом примере.

package main

import "fmt"

func main() {
	waterSI := NewSIFromL(5)
	waterSI.PrintInfo()
	fmt.Println()
	waterUSA := NewUSAFromPint(AdapterLToPint(waterSI.l)) // Применяем адаптер
	waterUSA.PrintInfo()
}

type pintaU float64

// Единицы объёма амереканской системы единиц
type UnitUSA struct {
	oz     pintaU
	pint   pintaU
	quart  pintaU
	gallon pintaU
}

func (us UnitUSA) PrintInfo() {
	fmt.Printf(
		"Unit USA:\n  %.3f oz\n  %.3f pint\n  %.3f quart\n  %.3f gallon\n",
		us.oz,
		us.pint,
		us.quart,
		us.gallon,
	)
}

// Конструктор структуры амереканской системы единиц из пинты
func NewUSAFromPint(pint pintaU) *UnitUSA {
	return &UnitUSA{
		oz:     16 * pint,
		pint:   pint,
		quart:  pint / 2,
		gallon: pint / 8,
	}
}

type literU float64

// Единицы объёма СИ
type UnitSI struct {
	ml literU
	l  literU
	m3 literU
}

func (si UnitSI) PrintInfo() {
	fmt.Printf("Unit SI:\n  %.3f ml\n  %.3f l\n  %.3f m^3\n", si.ml, si.l, si.m3)
}

// Конструктор структуры единиц СИ из литров
func NewSIFromL(l literU) *UnitSI {
	return &UnitSI{
		ml: 1e3 * l,
		l:  l,
		m3: 1e-3 * l,
	}
}

// Переводит литры в пинты
func AdapterLToPint(l literU) pintaU {
	return pintaU(l * 0.473)
}
