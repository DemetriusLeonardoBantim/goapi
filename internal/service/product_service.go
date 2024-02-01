package service

import (
	"github.com/DemetriusLeonardoBantim/goapi/internal/database"
	"github.com/DemetriusLeonardoBantim/goapi/internal/entity"
)

type ProductService struct {
	ProductDB database.ProductDB
}

func NewCProductService(productDB database.ProductDB) *ProductService {
	return &ProductService{ProductDB: productDB}
}

func (ps *ProductService) GetProducts() ([]*entity.Product, error) {
	products, err := ps.ProductDB.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) GetProduct(id string) (*entity.Product, error) {
	products, err := ps.ProductDB.GetProduct(id)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) GetProductByCategoryID(categoryID string) ([]*entity.Product, error) {
	products, err := ps.ProductDB.GetProductByCategoryID(categoryID)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) CreateProduct(name, description, category_id, image_url string, price float64) (*entity.Product, error) {
	product := entity.NewProduct(name, description, price, category_id, image_url)
	_, err := ps.ProductDB.CreateProducts((product))
	if err != nil {
		return nil, err
	}
	return product, nil
}
