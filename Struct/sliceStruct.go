package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {

	person1 := []Person{{Name: "Mikhael", Age: 20}, {Name: "Lion", Age: 30}, {Name: "Deni", Age: 25}}

	fmt.Println(person1)

	for _, data := range person1 {
		fmt.Printf("Hi... nama saya %s , Umur saya %d \n", data.Name, data.Age)
	}
}
