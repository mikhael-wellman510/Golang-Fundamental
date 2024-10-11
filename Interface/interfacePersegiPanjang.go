package main

import "fmt"

type Hitung interface {
	Luas() int
	Keliling() int
}

type Nilai struct {
	panjang int
	lebar   int
}

func (n Nilai) Luas() int {

	return n.panjang * n.lebar
}

func (n Nilai) Keliling() int {
	return 2 * (n.panjang + n.lebar)
}

func main() {

	var hitung Hitung

	hitung = Nilai{panjang: 4, lebar: 2}
	luas := hitung.Luas()
	keliling := hitung.Keliling()

	fmt.Println(luas)
	fmt.Println(keliling)
}
