package main

import (
	"fmt"
	repositori "golang-fundamental/Leson/Repositori"
)

func main() {

	// Instance Interface
	var barang repositori.CategoryRepository = &repositori.Category{}

	barang.Save(1, "Indomie")
	barang.Save(2, "Susu")
	barang.Save(3, "Sarden")

	data := barang.FindAll()
	fmt.Println(barang)
	fmt.Println(data)
}
