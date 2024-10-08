package main

import "fmt"

func main() {

	name := new(string)
	fmt.Println(name)
	// Untuk Menampilkan hasil / nilai nya dari alamat memori nya
	fmt.Print(*name)
}
