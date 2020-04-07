package service

import (
	"product-service/internals/entity"
	"product-service/internals/repository"
)

type ProductService struct {
	ProductRepository *repository.ProductRepository
}

// ListProducts ...
func (ps *ProductService) ListProducts() ([]*entity.Product, error) {
	products, err := ps.ProductRepository.GetAll()
	if err != nil {
		return []*entity.Product{}, err
	}
	return products, nil
}

// ImportProduct ...
func (ps *ProductService) ImportProduct(p *entity.Product) error{
	err := ps.ProductRepository.Add(p)
	if err != nil {
		return err
	}
	return nil
}


func NewProductService() *ProductService {
	return &ProductService{
		ProductRepository: repository.NewProductRepository(),
	}
}