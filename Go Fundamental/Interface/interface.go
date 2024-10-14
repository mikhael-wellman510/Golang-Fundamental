package main

import "fmt"

type Value struct {
	a int
	b int
}

type Calculator interface {
	Perkalian() int
	Pertambahan() int
}

func (v Value) Perkalian() int {

	return v.a * v.b
}

func (v Value) Pertambahan() int {
	return v.a + v.b
}

func main() {

	var calculator Calculator

	calculator = Value{a: 4, b: 5}

	hasilKali := calculator.Perkalian()
	hasilTambah := calculator.Pertambahan()

	fmt.Println(hasilKali)
	fmt.Println(hasilTambah)
}
