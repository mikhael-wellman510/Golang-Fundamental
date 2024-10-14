package main

import "fmt"

type Biodata struct{
	name string
	age int
	address string

}

func main(){

	biodata1 := Biodata{"Mikhael" , 25 , "Bogor"}
	biodata2 := &biodata1

	biodata2.name = "Deni"

	fmt.Println(biodata1) 
	fmt.Println(biodata2)

	// Dengan asterisk . dia akan merubah semua refrence awal
	*biodata2 =Biodata{"Deni" , 30 , "Bekasi"}

	fmt.Println(biodata1) 
	fmt.Println(biodata2)





}