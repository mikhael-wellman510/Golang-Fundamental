package main

import "fmt"

func main(){

	name := "Mikhael"
	sekolah := "SMK Tri Dharma"

	var res string = biodata(name,sekolah)
	fmt.Println(res)
}

func biodata(name string , sekolah string)string{

	return "Nama saya " + name + ",Saya sekolah di : " + sekolah 

}