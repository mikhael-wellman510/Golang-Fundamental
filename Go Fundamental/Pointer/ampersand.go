package main

import "fmt"

type Address struct{
	kota,provinsi,Country string
}

func main(){

	// Ampersand membuat variable mengikuti pointer di awal
	address1 := Address{"Bogor", "Jawa Barat" , "Indonesia"}
	address2 := &address1

	address2.kota = "Depok"

	fmt.Println(address1) // Ikut berubah
	fmt.Println(address2) // berubah

	address3 := Address{"Tangerang" ,"Banten" ,"Banten"}
	
	// *address2 = address3 // ini tetap ke pointer address1
	address2 = &address3 // Membuat address baru
	
	fmt.Printf("Cek : %p \n" , &address1)
	fmt.Printf("Cek : %p \n" , address2)

}