package main

import "fmt"

func main(){

	// Constant di golang tidak bisa di re-assigned
	// Jika variable Constant tidak error jika tidak di pakai


	const firstName string = "Mikhael"
	const lastName = "Wellman"

	// firstName = "Deni"
	// lastName = "Roston"

	//  Multiple Constant

	const (
		namaDepan string = "Mikhael"
		namaBelakang  = "Wellman"
	)

	// Tidak bisa di re-Assign
	// namaDepan = "Deni"

	fmt.Println("Nama Depan : " ,  namaDepan , ", Nama Belakang : " , namaBelakang)

}