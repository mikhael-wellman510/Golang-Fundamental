package main

import "fmt"

func main() {
	book := make(map[string]string)

	book["Judul"] = "Hantu Jembatan Ancol"
	book["Penulis"] = "Reo Kontol"
	book["tes"] = "Ups"

	fmt.Println("isi map : ", book)
	fmt.Println("Nama Penulis : ", book["Penulis"])

	delete(book, "tes")
	fmt.Println("updated : ", book)

}
