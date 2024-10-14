package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukan angka : ")
	scanner.Scan()
	input := scanner.Text() // Menghasilkan String

	if i, err := strconv.Atoi(input); err == nil {
		fmt.Print(i, " Adalah Angka !")
	} else {
		fmt.Println("Bukan angka")
	}

}
