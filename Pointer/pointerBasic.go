package main

import "fmt"

type Biodata struct{
	name string
	age int
	address string
}

func main(){

	// Pass by Value
	biodataA := Biodata{"Mikhael" , 25 , "Bogor"}
	biodataB := biodataA

	biodataB.name = "Lion"

	fmt.Println(biodataA) // Tidak Berubah
	fmt.Println(biodataB) // Berubah

	biodataC := Biodata{"Andre" , 23 , "Batam"}
	biodataD := &biodataC
	biodataD.name="Baskara"

	fmt.Println(biodataC) // berubah
	fmt.Println(biodataD) // berubah
	



}