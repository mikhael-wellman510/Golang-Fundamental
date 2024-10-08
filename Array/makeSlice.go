package main

import "fmt"

func main() {

	nama := make([]string, 2, 5)

	nama[0] = "Mikhael"
	nama[1] = "Deni"
	// nama[2] = "Anto" ini bakal error . karena maksimal 2

	// Dengan menggunakan append , dia akan mengalokasikan ulang memori
	namaBaru := append(nama, "Aldy", "edo", "ikbal", "andi")

	// Jika blm memenuhi capacity . nama juga akan terupdate
	// Jika capacity sudah terpenuhi . makan nama tidak akan terupdate
	namaBaru[0] = "Mikhael-Baru"

	fmt.Println("nama : ", nama)
	fmt.Println("Nama baru : ", namaBaru)
	fmt.Println(len(nama)) // 2
	fmt.Println(cap(nama)) // 5

	fmt.Println(len(namaBaru)) // 5
	fmt.Println(cap(namaBaru)) // 5
	// Jika tdi cap nya 5 , jika data nya trus bertambah ,.
	//Dia akan berlipat . dari 5 ke 10  , 10 ke 15
}
