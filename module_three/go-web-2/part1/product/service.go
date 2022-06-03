package product

type Service interface {
	Store(name string, price int) Product
	GetAll() []Product
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
func (s *service) Store(name string, price int) Product {
	newProduct := s.repository.Store(name, price)
	return newProduct
}

func (s *service) GetAll() []Product {
	newProduct := s.repository.GetAll()
	return newProduct
}
