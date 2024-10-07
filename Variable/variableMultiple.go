package main

import "fmt"

func main(){

	// Di golang , semua variable harus di pakai . kalau tidak dia akan error 

	var(
		firstName = "Mikhael"
		middleName = "Wellman"
		lastName = "Sitorus"
	)

	fmt.Println("nama Depan : " ,firstName , ", nama Tengah : "  , middleName, ", nama Belakang : " , lastName)
}