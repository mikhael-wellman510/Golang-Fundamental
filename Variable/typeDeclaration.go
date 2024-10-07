package main

import "fmt"

func main(){
	type NoKtp string

	var ktpMikhael NoKtp = "1212000123"
	var example = "12aa12239"
	// Harus di Casting 
	var ktpUdin NoKtp= NoKtp(example) 

	fmt.Printf("ktp : %s \n " , ktpMikhael)
	fmt.Printf("Tes : %s \n " , ktpUdin)
}