
package repository

import "product-service/internals/entity"

type ProductRepository struct {
	ProductDB []*entity.Product
}

func (pr *ProductRepository) Add(product *entity.Product) error {
	pr.ProductDB = append(pr.ProductDB, product)
	return nil
}

func (pr *ProductRepository) GetAll() ([]*entity.Product, error) {
	return pr.ProductDB, nil
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		ProductDB: []*entity.Product{},
	}
}
