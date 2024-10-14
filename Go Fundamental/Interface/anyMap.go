package main

import (
	"fmt"
)

func main() {

	data := make(map[string]interface{})

	data["Nama"] = "Mikhael"
	data["Umur"] = 20

	fmt.Println(data)

	var biodataSiswa map[any]string = map[any]string{}

	biodataSiswa["Mikhael"] = "Nama"
	biodataSiswa[20] = "umur"

	bio := make(map[interface{}]string)

	bio[1] = "Deni"
	bio["Nama"] = "Andree"

	fmt.Println("Cek : ", bio)

	var biodata map[string]interface{} = map[string]interface{}{
		"nama": "Mikhael",
		"Umur": 20,
	}

	fmt.Println(biodata)

	biodata2 := map[string]interface{}{
		"Nama":  "Mike",
		"usia":  20,
		"berat": 12.5,
	}

	fmt.Println(biodata2)

}
