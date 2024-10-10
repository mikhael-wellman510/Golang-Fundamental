package main

import (
	"fmt"
	"strings"
)

func main(){
	nama := [] string {"Mikhael" , "Wellman"}
	
	sayHello2("Hello", nama)
}

func sayHello2(message string , arr [] string){
	name := strings.Join(arr, " ")
	fmt.Println(message , name)
}

