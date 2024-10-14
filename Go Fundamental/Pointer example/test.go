package main

import "fmt"

// merubah
func changes(original *int, val int) {

	*original = val
}

// Tidak berubah
func changes2(original *int, val int) {
	original = &val

}

func main() {

	num := 5
	num2 := 11
	changes(&num, 10)
	fmt.Println(num)
	changes2(&num2, 20)

	fmt.Println(num2)
}
