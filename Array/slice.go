package main

import "fmt"

func main() {
	month := []string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}

	slice1 := month[2:5] // Hasil nya [Maret,April,Mei]
	slice2 := month[2:2] // Hasil nya []
	slice3 := month[:2]  // Hasil nya [Januari, Februari]
	slice4 := month[9:]  // Hasil nya [Oktober , November , Desember]
	slice5 := month[:]   // Hasil nya semua isi array

	// ini akan tetap merefrensi ke Month , Tidak membuat memori baru
	slice4[1] = "November-Baru"
	slice1[2] = "Mei-Baru"

	fmt.Printf("Hasil 1 : %s \n", slice1)
	fmt.Println("Hasil 2 : ", slice2)
	fmt.Println("Hasil 3 : ", slice3)
	fmt.Println("Hasil 4 : ", slice4)
	fmt.Println("Hasil 5 : ", slice5)

	fmt.Println("Month : ", month)

	nama := [...]string{"mike", "aldi", "deni"}

	hasil := nama[1:2]
	fmt.Println(hasil)
}
