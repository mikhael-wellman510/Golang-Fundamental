package main

type CategoryRepository interface {
	Save(id int, produk string)
}
