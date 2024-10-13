package main

import "fmt"

type Address3 struct{
	City , Province, Country string
}

func changeCountry(addres1 *Address3){
	addres1.City = "Tangerang"
	addres1.Province = "Banten"

}

func main(){

	address1 := Address3{"Bogor" , "Jawa Barat" , "Indonesia"}

	changeCountry(&address1)

	fmt.Println(address1)
	

}