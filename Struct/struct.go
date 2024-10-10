package main

import "fmt"

type Customer struct {
	Name, Addres string
	Age          int
}

func main() {
	var eko Customer

	fmt.Println("Init Awal : ", eko)

	eko.Name = "Eko Khanedy"
	eko.Addres = "Bogor"
	eko.Age = 30

	fmt.Println("Result : ", eko)
}
