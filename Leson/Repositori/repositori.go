package repositori

type Category struct {
	Id      int
	Product string
}

type CategoryRepository interface {
	Save(id int, product string)
	FindAll() []Category
}
