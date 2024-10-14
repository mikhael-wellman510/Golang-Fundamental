package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukan huruf : ")
	scanner.Scan() // Ini wajib untuk membaca scanner nya
	input := scanner.Text()

	fmt.Print("\n Masukan angka")
	scanner.Scan()
	inputAngka := scanner.Text()

	angka, _ := strconv.Atoi(inputAngka) // Konversi dari string ke Integer

	fmt.Println("result String : ", input)
	fmt.Println("Result angka : ", angka)
}
