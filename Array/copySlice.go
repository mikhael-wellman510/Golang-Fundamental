package main

import "fmt"

func main() {

	var hari = []string{"senin", "selasa", "rabu", "kamis", "jumat"}

	dataHari := make([]string, 3, 5)

	copy(dataHari, hari) // param1 : tujuan copy nya , param2 : yg mau di copy

	fmt.Printf("Hari : %s \n", hari)
	fmt.Println("data hari ", dataHari)

	dataHari[2] = "rabu-baru"

	fmt.Printf("Hari : %s \n", hari)
	fmt.Println("data hari ", dataHari)

}
