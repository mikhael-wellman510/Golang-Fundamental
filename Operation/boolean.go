package main

import "fmt"

func main() {

	var left bool = false
	var right bool = true

	resultAnd := left && right
	resultOr := left || right

	fmt.Printf("And %t \n", resultAnd) // False
	fmt.Printf("or %t \n", resultOr)   // True
	fmt.Printf("Negasi %t \n", !left)  // True

	var nama string = "Baskara"
	var nilaiAkhir = 80
	var absensi = 90

	var lulusNilaiAkhir bool = nilaiAkhir > 90
	var lulusAbsensi bool = absensi > 85

	var lulus bool = lulusNilaiAkhir && lulusAbsensi

	fmt.Printf("Apakah %s lulus kuliah ? %t ", nama, lulus)

}
