package main

import "fmt"

type Total struct {
	hasil int
}

func tambah(data *Total, val int) {
	data.hasil = data.hasil + val + 10

}

func main() {
	data := &Total{}

	tambah(data, 20)
	tambah(data, 20)
	fmt.Println(data)

}
