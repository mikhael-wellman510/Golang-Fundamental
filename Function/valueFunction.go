package main

import "fmt"


func goodBye(val... string)string{
	
	say := "hi , " + val[0] + " dan " + val[1]

	return say

}

func main(){

	nama := []string{"Mikhael" ,"Lion"}

	say := goodBye

	fmt.Println(say(nama...))
	
}