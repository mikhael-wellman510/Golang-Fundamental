package main

import "fmt"

type Person struct {
	nama string
	umur int
}

func (p *Person) getName() string {

	return p.nama
}

func main() {

	var data any = &Person{"Mikhael", 28}

	res := data.(*Person).nama

	fmt.Println("Cek : ", res)

}
