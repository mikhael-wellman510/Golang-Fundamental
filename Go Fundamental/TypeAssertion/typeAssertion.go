package main

import (
	"fmt"
)

func random() any {

	return 15
}

func main(){

	// Interface jika menjadi interface{} -> ini jadi tipe data
	

	result := random()
	
	switch result.(type) {
	case string :
		fmt.Println("String " , result)
	case int :
		fmt.Println("Integer " , result)
	case bool :
		fmt.Println("Bool " , result)
	default:
		fmt.Println("Unknown")
	}



}