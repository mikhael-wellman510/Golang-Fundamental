package main

import (
	"fmt"
)

func main(){
	
	var data [] string

	nama := "mikhael"


	for _ , n := range nama {
		data = append(data, string(n))
	}

	fmt.Println(data)


	fmt.Println(data[1])

}