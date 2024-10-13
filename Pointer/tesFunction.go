package main

import "fmt"

func change(original *int , val int){
	 *original = val
}

func main(){

	var num int = 5

	change(&num , 10)

	fmt.Println(num)

}