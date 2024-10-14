package main

import "fmt"

type Mahasiswa struct{
	name string
	nip int
}

func (m *Mahasiswa) changeName(name string){
	m.name = name
}

func main(){
	mhs := Mahasiswa{"Mikhael" , 19910073}

	fmt.Println(mhs)
	mhs.changeName("Deni")
	fmt.Println(mhs)
	
}