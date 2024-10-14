package main

import "fmt"

type Address2 struct{
	kota , provinsi , negara string
}

func main(){

	address1 := new(Address2)
	address2 := address1

	address2.negara= "Indonesia"

	fmt.Println(address1)

}