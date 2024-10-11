package main

import "fmt"

type ProductService interface {
	addProduct() Product
}

type Product struct {
	id          int
	productName string
}

type Db struct {
	listProduct []Product
}

type Repositori interface {
	save(p Product)
}

func (db *Db) save(p Product) {
	db.listProduct = append(db.listProduct, p)
	fmt.Println("Product saved:", p.productName)
}

func (p Product) addProduct(repo Repositori) Product {

	repo.save(p)
	return p
}

func main() {
	// Inisialisasi repositori
	db := &Db{}

	// Buat produk baru
	produk := Product{id: 1, productName: "Indomie"}
	produk2 := Product{id: 2, productName: "Bakso"}

	produk.addProduct(db)
	produk2.addProduct(db)

	fmt.Println("All products:", db.listProduct)
}
