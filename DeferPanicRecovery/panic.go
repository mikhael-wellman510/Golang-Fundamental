package main

import "fmt"

func endApp() {
	fmt.Println("End App")
}

func runApp(err bool) {

	if !err {
		// Dia akan menghientikan program
		panic("Error Bos")
	}

	fmt.Println("Tidak Error")
}

func main() {
	runApp(true)
}
