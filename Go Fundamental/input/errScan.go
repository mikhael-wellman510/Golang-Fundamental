package main

import "fmt"

func main() {

	var inputAngka int

	_, err := fmt.Scanln(&inputAngka)

	fmt.Println(err) // Kalau benar ,. err akan menghasilkan <nil>

	if err != nil {
		fmt.Println("Yang anda masukan bukan angka")
	}

}
