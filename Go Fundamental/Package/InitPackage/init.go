package initpackage

import "fmt"

var connection string

func Tes() {
	fmt.Println("Tes")
}

// Init akan di jalankan terlebih dahulu
func init() {
	fmt.Println("Init ---")
	Tes()
	connection = "Postgre Sql Connect"
}

func GetDatabase() string {

	fmt.Println("Get Db 00000")

	return connection
}
