package main

import "fmt"

func main() {

	var data map[string]string = map[string]string{}
	data["Nama"] = "Mikhael"
	data["Alamat"] = "Bogor"
	fmt.Println(data)
	fmt.Println("Nama nya adalah : ", data["Nama"])

	biodata := map[string]string{}

	biodata["nama"] = "Mikhael"
	biodata["alamat"] = "Bogor"

	fmt.Println(biodata)
	fmt.Println(biodata["nama"])
	fmt.Println(biodata["alamat"])

	person := map[string]string{
		"name":    "Mikhael",
		"address": "Bogor",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)
}
