package main

import "fmt"

type People struct {
	name string
	age int
	address string
}

func myData()map[string]People{

	person :=[] People{{"Mikhael" , 20 , "bogor"} ,
			{"Aldy" , 23 , "Jakarta"},
			{"Deni" , 25 , "Tangerang"},
			{"Carlos" ,19 , "Depok"}}

	data := make(map[string]People)

	for i, val := range person{
			if val.age >= 20 {
				key := fmt.Sprintf("No %d" , i+1)
				data[key] = val
			}
	}
	
	
	return data
}

func main(){

	value := myData()

	fmt.Println(value["No 1"])
	 
}