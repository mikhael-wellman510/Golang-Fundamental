package main

import (
	"fmt"
)

func main(){

	nama := []string{}

	// create
	nama = append(nama, "Mikhael" ,"Deni" ,"Carlos" ,"Aldi")

	fmt.Println("Data : " ,nama)
	// Deleted
	idx := 1

	nama = append(nama[:idx] ,nama[idx+1:]... )
	
	fmt.Println("Deleted : " , nama)

	// Deleted by name

	siswa := [] string{"Mike" ,"Aldo" ,"Handi" ,"Kevin"}
	hapus := "Handi"
	siswa = removeNameFromSlice(siswa,hapus) 
	
	fmt.Println("data siswa terbaru : " ,siswa)


	//Updated

	idxUpdt := 2

	nama[idxUpdt] = "Aldi Baru"

	fmt.Println("Updated data terbaru : " , nama)


	// Update by name
	mahasiswa := []string {"Mike", "Carlos" ,"Aldi" ,"Deni"}

	update := "Carlos"
	toBe := "Carlos-Baru"

	mahasiswa = updateData(mahasiswa, update,toBe)
	fmt.Println("updated fix : " , mahasiswa)

}

func updateData(mahasiswa [] string , name string , toBe string)[]string{

	for idx , siswa := range mahasiswa {
		if(name == siswa){
			mahasiswa[idx] = toBe
		}
	}

	
	return mahasiswa
}

func removeNameFromSlice(names [] string , namesRemove string)[]string{
	temp := names[:0] // Di buat kosong

	for _ , nama := range names {
		if (nama != namesRemove){
			temp = append(temp, nama)
		}
	}

	return temp
}