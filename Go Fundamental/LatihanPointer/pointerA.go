package main

import "fmt"

func main() {

	// Menggunakan Pointer
	x := 10
	changeValue(&x)
	fmt.Println(x)

	// tanpa Pointer

	y := 20
	gantiValue(y)
	fmt.Println(y)

}

func changeValue(x *int) {
	*x = 20
}

func gantiValue(y int) {
	y = 50
}
