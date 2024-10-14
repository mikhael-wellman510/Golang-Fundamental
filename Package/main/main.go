package main

import (
	"fmt"
	initpackage "golang-fundamental/Package/InitPackage"
	"golang-fundamental/Package/helper"
	"golang-fundamental/Package/persegi"
)

func main() {

	// Mengambil dari package helper
	result := helper.SayHello("Mikhael")
	fmt.Println(result)

	luasPersegi := persegi.Luas(5)
	fmt.Println(luasPersegi)

	kelilingPersegi := persegi.Keliling(10)
	fmt.Println(kelilingPersegi)

	// Init Package

	getDatabase := initpackage.GetDatabase()
	fmt.Println("get db : ", getDatabase)

}
