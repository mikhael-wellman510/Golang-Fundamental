package main

import "fmt"

type Cek func(string) bool

func validation(name string , cek Cek){
	if cek(name) {
		fmt.Println("Kata terlalu kasar")
	}else{
		fmt.Println("Succes Registration")
	}

}

func main(){

	cek := func(val string) bool{
		return val == "anjing"
	}

	validation("mikhael", cek)

	validation("anjing" , func(val string)bool{
		return val == "anjing"
	})

}