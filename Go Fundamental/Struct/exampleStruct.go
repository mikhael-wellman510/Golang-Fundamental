package main

import "fmt"

type Biodata struct {
	Nama    string
	Usia    int
	Sekolah []string
}

func main() {
	// Cara 1
	var mahasiswa Biodata

	mahasiswa.Nama = "Mikhael"
	mahasiswa.Usia = 20
	mahasiswa.Sekolah = []string{"SMK TRI DHARAMA", "POLITEKNIK NEGERI MEDIA KREATIF"}

	fmt.Println(mahasiswa)

	// Cara 2
	mahasiswa1 := Biodata{
		Nama:    "Mikhael Wellman",
		Usia:    30,
		Sekolah: []string{"Purwadhika", "Enigma"},
	}

	fmt.Println(mahasiswa1)

	// Cara 3

	mahasiswa2 := Biodata{"Mikhael Wellman", 30, []string{"PNJ", "Undip"}}
	fmt.Println(mahasiswa2)
}
