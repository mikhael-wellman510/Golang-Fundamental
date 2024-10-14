package main

import "fmt"

func main() {

	var numA int = 4
	var numB *int = &numA
	// & : Menjadikan Alamat memori / Address
	// * : menjadikan Value (mengubah Alamat memori menjadi Value)

	// Nilai default variable pointer adalah nil
	// Variable pointer tidak bisa menampung yang bukan pointer
	// Variavle pointer menghasilkan Stirng alamat memori (Hexadesimal)

	fmt.Println("Num A : ", numA)          // 4
	fmt.Println("Num B Pointer : ", numB)  //0xc00000a0d8
	fmt.Println("Num A Pointer : ", &numA) //0xc00000a0d8
	fmt.Println("Num B : ", *numB)         // 4

	nums := 12
	var data1 *int = &nums
	var data2 string = "Mikhael"

	fmt.Println(data1)
	fmt.Println("convert jadi value : ", *data1)
	fmt.Println(data2)
	fmt.Println("Convert jadi alamat memori : ", &data2)
}
