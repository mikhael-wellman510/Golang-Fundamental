package main

import "fmt"

func main(){
	// Jika hanya declare nama variable . wajib menggunakan tipe data
	var name string

	name = "Mikhael Wellman"
	fmt.Println(name)

	name = "Mikhael Sitorus"
	fmt.Println(name)
    
	// Golang bisa mendeteksi tipe data nya langsung 
	var nama = "Mikhael Wellman Gans"
	fmt.Println(nama)

	nama = "Mikhael Gokilss"
	fmt.Println("Nama gaul nya : " , nama)


	// Variable tanpa menggunakan Var

	sekolah := "SMP Negeri 20"
	fmt.Println("Nama Smp : " ,sekolah)

	sekolah = "SMK Tri Dharma 1"
	fmt.Println("Nama SMK : " , sekolah)

	kuliah := "Politeknik Negeri Media Kreatif Jakarta"
	fmt.Println("Kampus saya di : " , kuliah)



}