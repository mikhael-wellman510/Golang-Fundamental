package main

import "fmt"

func main() {

	// Jika make , biasanya bisa membuat capacity .
	book := make(map[string]string)

	book["Judul"] = "Hantu Jembatan Ancol"
	book["Penulis"] = "Reo Kontol"
	book["tes"] = "Ups"

	fmt.Println("isi map : ", book)
	fmt.Println("Nama Penulis : ", book["Penulis"])

	delete(book, "tes")
	fmt.Println("updated : ", book)

}
