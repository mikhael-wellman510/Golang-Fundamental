package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Addres string
}

func (P Person) sayHello() {
	fmt.Printf("Hello %s , Salam kenal ! \n", P.Name)
}

func (P Person) fullBio() string {
	// Sprint F untuk menyimpan ke dalam variable
	// bio := fmt.Sprintf("Hallo , Nama saya :  %s . Umur saya : %d . Saya tinggal di : %s \n", P.Name, P.Age, P.Addres)
	bio := "Hallo , Nama saya " + P.Name + " , Saya Tinggal di " + P.Addres
	return bio
}

func main() {
	person1 := Person{"Mikhael Wellman", 18, "Bogor"}
	person2 := Person{"Agus Yudhoyono", 20, "Depok"}
	person1.sayHello()
	person2.sayHello()
	var result1 string = person1.fullBio()
	var result2 string = person2.fullBio()
	fmt.Println(result1)
	fmt.Println(result2)
}
