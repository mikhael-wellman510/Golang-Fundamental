package main

import "fmt"

type Biodata struct {
	name string
	umur int
}

func tambahUmur(bio *Biodata, n int) {
	bio.umur = bio.umur + n

}

func main() {

	bio := &Biodata{"Mikhael", 20}

	tambahUmur(bio, 20)

	fmt.Println(bio)
}
