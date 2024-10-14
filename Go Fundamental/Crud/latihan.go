package main

import "fmt"

type ProductNew struct {
	id          string
	productName string
	price       int
	qty         int
}

type Database struct {
	listProducts []ProductNew
}

// Method
func (db *Database) save(p ProductNew) {
	db.listProducts = append(db.listProducts, p)
}

func (db *Database) findById(id string) *ProductNew {
	for i, _ := range db.listProducts {
		if db.listProducts[i].id == id {
			return &db.listProducts[i]
		}
	}

	return nil
}

func (db *Database) findAll() []ProductNew {
	return db.listProducts
}

func (db *Database) deleted(id string) {

	var idx int

	for i, val := range db.listProducts {
		if val.id == id {
			idx = i
		}
	}

	data1 := db.listProducts[:idx]
	data2 := db.listProducts[idx+1:]

	db.listProducts = append(data1, data2...)
}

// Service
func addProducts(db *Database, products ProductNew) {
	productRepositori := db
	productRepositori.save(products)

}

func findById(db *Database, id string) {
	productRepositori := db
	res := productRepositori.findById(id)
	fmt.Println(res)
}

func findAll(db *Database) []ProductNew {
	result := db.findAll()
	return result
}

func deletedProductsById(db *Database, id string) {
	productsRepositori := db
	productsRepositori.deleted(id)
}

func updatedProducts(db *Database, updateProduct ProductNew) {
	productRepositori := db
	data := productRepositori.findById(updateProduct.id)

	if data != nil {
		data.id = updateProduct.id
		data.price = updateProduct.price
		data.productName = updateProduct.productName
		data.qty = updateProduct.qty
	} else {
		fmt.Println("nil")
	}

}

func main() {
	// Instance db
	db := &Database{}

	// controller
	p1 := ProductNew{"uu1", "Sabun", 4000, 20}
	p2 := ProductNew{"uu2", "Indomie", 3000, 100}
	p3 := ProductNew{"uu3", "Susu", 1000, 40}
	p4 := ProductNew{"uu4", "Rokok", 13000, 5}
	addProducts(db, p1)
	addProducts(db, p2)
	addProducts(db, p3)
	addProducts(db, p4)

	findById(db, "uu1")

	deletedProductsById(db, "uu3")
	deletedProductsById(db, "uu4")
	updated := ProductNew{"uu1", "Beras", 30000, 10}
	updatedProducts(db, updated)
	result := findAll(db)
	fmt.Println(result)

}
