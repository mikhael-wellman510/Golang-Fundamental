package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	var c = 2
	var result = a + b*c

	fmt.Printf("hasil dari %d + %d x %d adalah %d ! \n", a, b, c, result)

	var x = 10
	x += 5

	fmt.Println(x)
	x += 10
	fmt.Println("cek", x)

}
