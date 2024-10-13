package main

import "fmt"

func newMap(name string)map[string]string {

	if name != "" {
		return map[string]string{
			"name": name,
		}
	}else {
		return nil
	}
}

func tes () any{

	return nil
}



func main(){

	result := newMap("")
	
	fmt.Println("Res : " , result)
	res := tes()

	fmt.Println(res)

}