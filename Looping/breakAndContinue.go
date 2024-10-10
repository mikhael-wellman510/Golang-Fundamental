package main

import "fmt"

func main(){

	// Break akan langsung menghentikan perulangan
	for i := 0 ; i < 10 ; i++ {
		if(i==5){
			break
		}

		fmt.Println("Loop ke : " , i)
	}

	// Continue akan melewatkan eksekusi dan lanjut looping 
	for i := 0 ; i < 10 ; i++{
		
		if (i == 3 || i == 7) {
			continue
			}

		fmt.Println("idx ke : " , i)
}
}