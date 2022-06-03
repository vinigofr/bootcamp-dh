package product

type Repository interface {
	Store(name string, price int) Product
	GetAll() []Product
}

type repository struct{}

var globalProducts []Product
var globalID = 1

func (repository) Store(name string, price int) Product {
	var newProduct = Product{Name: name, Price: price, Id: globalID}
	globalProducts = append(globalProducts, newProduct)
	globalID++
	return newProduct
}

func (repository) GetAll() []Product {
	return globalProducts
}

func NewRepository() Repository {
	return &repository{}
}
