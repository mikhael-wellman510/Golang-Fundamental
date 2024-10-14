package main

import "fmt"

func main() {

	var person = []map[string]interface{}{
		{"nama": "mikhael", "usia": 20},
		{"nama": "deni", "usia": 15},
		{"nama": "aldi", "usia": 24},
	}

	for _, data := range person {
		fmt.Printf("hallo , my name %s , usia saya : %d \n", data["nama"], data["usia"])
	}
}
