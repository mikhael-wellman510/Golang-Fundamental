package main

import "fmt"

func data() any {

	return 1
}

func data2() any {
	return "Hello World"
}

func main() {

	// interface kosong juga bisa disebut sebagi anny
	var tes any = data()
	var tes2 any = data2()

	fmt.Println(tes)
	fmt.Println(tes2)

	var free interface{}

	free = "Mikhael"
	fmt.Println(free)

	free = 200
	fmt.Println(free)

	free = true
	fmt.Println(free)
}
