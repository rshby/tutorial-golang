package product

import "training/tour/model"

type IProductService interface {
	GetProducts() []*model.Product
	GetProduct(id int) *model.Product
	InsertProduct(entity *model.Product) (*model.Product, error)
}
