package main

import "fmt"

type Kampus struct {
	NamaKampus string
	Jurusan    string
}

type Mahasiswa struct {
	Nama string
	Usia int
	Kampus
}

func main() {
	// Cara 1
	var mahasiswa1 Mahasiswa

	mahasiswa1.Nama = "Mikhael Wellman"
	mahasiswa1.Usia = 20
	mahasiswa1.NamaKampus = "Polimedia"
	mahasiswa1.Jurusan = "Multimedia"

	fmt.Println(mahasiswa1)
	fmt.Println("Jurusan : ", mahasiswa1.Kampus.Jurusan)
	fmt.Println("Jurusan : ", mahasiswa1.Jurusan)
}
