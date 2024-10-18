package repositori

var listProduct []Category

func (c *Category) Save(id int, product string) {
	c.Id = id
	c.Product = product
	listProduct = append(listProduct, *c)
}

func (c *Category) FindAll() []Category {

	return listProduct
}
