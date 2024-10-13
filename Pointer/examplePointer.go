package main

import "fmt"

func main(){

	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("a : " , &numberA)
	fmt.Println(numberB)

	numberA = 5

	fmt.Println(*numberB) // jadi 5 
}