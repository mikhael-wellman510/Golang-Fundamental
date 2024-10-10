package main

import (
	"fmt"
)

func main(){

	// For Manual
	for i := 0 ; i < 10 ; i++ {
		fmt.Println("Counter ke - " , i)
	}

	nama := [...]string {"Mikhael" , "Deni" , "Lion" ,"Aldi"}


	for i := 0 ; i < len(nama) ; i++ {
		fmt.Println("Nama : " , nama[i])
	} 

	// for Range

	var data string = "mike"

	for i , v := range data{
		fmt.Printf("Index : %d , value : %c \n" , i ,v)
	}

	day := [] string {"Senin" ,"Selasa" ,"Rabu"}
	// _ = itu karena tidak d pakai init nya
	for _, hari := range day {
		fmt.Println("Hari : " , hari)
	}

	for i , _ := range day {
		fmt.Println("Hari Ke : " , i)
	}



}