package product

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (*Service) List() []Product {
	return allProducts
}

func (*Service) Get(id int) (*Product, error) {
	return &allProducts[id], nil
}
