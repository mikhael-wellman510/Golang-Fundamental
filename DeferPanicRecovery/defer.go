package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil Function")
	message := recover() // Membuat function ini tetap berjalan
	fmt.Println("Terjadi Panic : ", message)
}

func runApplication(err bool) {
	// Defer akan selalu terpanggil di paling terakhir setelah menjalankan semua baris program
	defer logging()

	if !err {
		panic("Gw Panic")
	}

	fmt.Println("Run Application")
	fmt.Println("Succes")
}

func main() {
	runApplication(false)
}
