package main

import "fmt"

func subTotal(number ...int)int{

	total := 0

	for _, val := range number {
		total += val
	}

	return total;

}

func cekTotal(val...int)int{
	
	fmt.Println(val)

	return 0;
}

func main(){

	result := subTotal(10,5,15,30)
	data := []int{10,20,30}
	cekTotal(data...)

	fmt.Println(result)

}