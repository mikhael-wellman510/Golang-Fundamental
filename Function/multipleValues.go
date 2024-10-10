package main

import "fmt"

func getFullName() (string, string){

	return "Mikhael" , "Wellman"
}

func getSchool()(string, string){

	return "Polimedia" ,"Pnj"
}

func main(){

	firstName , lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)

	school1 , _ := getSchool()
	fmt.Println(school1)
}