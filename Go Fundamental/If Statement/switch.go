package main

import "fmt"

func main() {

	name := "jokowi"

	switch name {
	case "muliono":
		fmt.Println("Yo ndak tau ")
	case "jokowi":
		fmt.Println("Silahkan ambil sepedah nya")
	case "gibran":
		fmt.Println("Di bantu paman")
	default:
		fmt.Println("Tidak Ditemukan")
	}

	angka := 4

	switch angka {
	case 1:
		fmt.Println("Asik")
	case 2:
		fmt.Println("Bjirr")
	case 3, 4, 5, 6:
		fmt.Println("Good")
	default:
		fmt.Println("else")
	}

}
