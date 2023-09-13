package product

import "training/tour/model"

type ProductService struct {
	Name string
}

func NewProductService(name string) *ProductService {
	return &ProductService{
		Name: name,
	}
}

// method to get all data
func (p *ProductService) GetProducts() []*model.Product {
	return nil
}

// method to get data product by id
func (p *ProductService) GetProduct(id int) *model.Product {
	return nil
}

// method to insert data product
func (p *ProductService) InsertProduct(entity *model.Product) (*model.Product, error) {
	return nil, nil
}
