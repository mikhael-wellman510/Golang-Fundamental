package main

import "fmt"

type Biodata struct {
	nama   string
	usia   int
	alamat string
}

func main() {

	var data map[string]Biodata = map[string]Biodata{}

	data["person1"] = Biodata{nama: "Mikhael", usia: 20, alamat: "bogor"}
	data["person2"] = Biodata{nama: "Andre", usia: 18, alamat: "batam"}
	fmt.Println(data)

	for _, data := range data {
		fmt.Println("Nama : ", data.nama, " Usia : ", data.usia, " Tinggal : ", data.alamat)
	}
}
