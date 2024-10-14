package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Mikhael"
	names[1] = "Aldy"
	names[2] = "Carlos"

	panjang := len(names)

	fmt.Println(names)
	fmt.Printf("Panjang array berikut adalah %d \n", panjang)

	var makanan = [3]string{"Durian", "Semangka", "Jeruk"}

	fmt.Println(makanan[0])
	fmt.Println(makanan[1])
	fmt.Println(makanan[2])

	var values = [...]string{
		"Bogor",
		"Jakarta",
		"Depok",
		"Tangerang",
	}

	panjangArray := len(values)

	fmt.Println(values)
	fmt.Printf("Panjang array adalah %d ", panjangArray)
}
