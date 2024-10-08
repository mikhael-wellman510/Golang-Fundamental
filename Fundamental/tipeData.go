package main

import "fmt"

func main() {

	// Uint itu ke angka positif. ga bisa angka negativ
	var angkaPositif1 uint8 = 255
	fmt.Println("Positif 1 : ", angkaPositif1)
	var angkaPositif2 uint16 = 65535
	fmt.Println("Positif 2 : ", angkaPositif2)
	var angkaPositif3 uint32 = 4294967295
	fmt.Println("Positif 3 : ", angkaPositif3)

	var angkaNegative int8 = -128
	var angkaNegative2 int16 = -32768

	fmt.Println("negative ", angkaNegative)
	fmt.Println("Negative 2 ", angkaNegative2)
}
