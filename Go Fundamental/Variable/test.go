package main

import "fmt"

func main(){
	nama := "Mikhael"
	
	fmt.Println("Awal : " , nama);

	nama = "Deni"

	fmt.Println("Akhir : " , nama)

	sekolah := "Polimedia"
	kampus := sekolah
	fmt.Println("sekolah : " , sekolah)
	fmt.Println("Kampus : " ,  kampus)

	sekolah = "Smk Tri Dharma 1"
	fmt.Println(sekolah)
	
}