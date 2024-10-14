package main

import "fmt"

func main() {

	// Append akan membuat refrensi baru / membuat alamat memori baru
	days := []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	updateDays := append(days, "Senin-Libur")

	fmt.Println("days : ", days)

	// Dia membuat memori baru , dan tidak akan merubah isi refrence asli nya
	fmt.Println("Update Days : ", updateDays)

	newDays := days[1:3] // tidak akan berubah walaupun di append

	updateNewDays := append(newDays, "Kamis-Libur") // Membuat alamat memori baru

	fmt.Println("new Days : ", newDays)
	fmt.Println("update New Days ", updateNewDays)

}
