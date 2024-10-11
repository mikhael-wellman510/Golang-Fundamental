package main

import "fmt"

func main() {
	var data map[string]any

	data = map[string]any{
		"name": "Mikhael",
		"umur": 20,
	}

	// jika memakai any dan tipe data nya int , harus di casting dulu

	fmt.Println(data["umur"].(int) * 20)

}
