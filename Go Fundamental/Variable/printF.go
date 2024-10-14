package main

import "fmt"

func main(){
	// Angka 
	// Uint -> Tidak ada bilangan negative
	// int -> Bisa bilangan negative
	var angka uint = 10
	fmt.Printf("angka : %d \n" , angka)

	var negativeNumber int = -1234
	fmt.Printf("Angka Negative %d \n" , negativeNumber) 

	angka1 := 20
	fmt.Printf("tes Angka : %d \n" ,angka1)


	// Huruf
	var huruf string = "Hello"
	fmt.Printf("huruf : %s \n" , huruf)

	huruf1 := "Mikhael Wellman"
	fmt.Printf("Nama Saya : %s \n" , huruf1)

	const nama = "Deni"
	fmt.Printf("Nama Asli saya : %s \n " , nama)

	// Decimal
	var angkaDecimal float32 = 30.5
	fmt.Printf("Angka Decimal : %f \n " , angkaDecimal )
	angkaDecimal1 := 30.5
	fmt.Printf("Tes : %.2f \n " , angkaDecimal1)


	// Boolean
	var isExist bool = true
	isOpen := false
	fmt.Printf("Apakah exist : %t \n " , isExist)
	fmt.Printf("Apakah Open : %t \n " , isOpen)

}