package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {

	if pembagi == 0 {
		return 0, errors.New("Tidak bisa membagi Nol")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {

	res, err := Pembagian(12, 0)

	if err == nil {
		fmt.Println("result : ", res)
	} else {
		fmt.Println("Err : ", err.Error())
	}
}
