package main

import (
	"fmt"
)

func main(){

	// Tipe Data   || Min  || Max
	// int 8          -128    127
	// int 16         -32768  32767
	// int 32  		  -2147483648 2147483647
	// int 64 		  -9223372036854775808  -9223372036854775808


	// Jika menggunakan tipe data nya dan tidak sesuai maka akan overflows
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai64)
	fmt.Println("Nilai 32 : " , nilai32)
	fmt.Println("Nilai 64 : " , nilai64)
	fmt.Println("Nilai 16 : " , nilai16)

	var tes32 int16 = 2000
	var tes8 int8 = int8(tes32)
	fmt.Println(tes32)
	fmt.Println(tes8)


	nama := "Mikhael Wellman" // Nama tidak erro . karena di panggil di reference
	idxNama := nama[1];
	idxHuruf := string(idxNama)
	
	fmt.Println("angka byte : " , idxNama)
	fmt.Println("huruf : " , idxHuruf)
}