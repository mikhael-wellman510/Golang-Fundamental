package main

import (
	"fmt"
)


func myMap(name string)map[string]string{
	people := make(map[string]string)

	people["name"] = name
	people["address"] = "Bogor"
	people["Hobby"] = "Coding"
	
	
	return people

}

func data(name []string , umur []int)map[int]string{
	
	var data map[int]string = map[int]string{}

	for i , key := range name {
		data[i] = key
	}

	return data

}

func main(){

	result := myMap("mikhael")
	fmt.Println(result)

	name := []string{"Mikhael","Aldy" ,"Lion"}
	umur := []int{20,25,30}

	result2 := data(name,umur)
	fmt.Println(result2)
}