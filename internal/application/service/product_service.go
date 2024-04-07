package service

import (
	"github.com/kutsuna/simple-go-backend-template/internal/domain/model"
	"github.com/kutsuna/simple-go-backend-template/internal/domain/repository"
)

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) GetProductByID(id string) (*model.Product, error) {
	return s.repo.FindByID(id)
}

func (s *ProductService) GetAllProducts() ([]*model.Product, error) {
	return s.repo.FindAll()
}

func (s *ProductService) CreateProduct(name string, price float64) (*model.Product, error) {
	newProduct := &model.Product{
		Name:  name,
		Price: price,
	}
	return newProduct, s.repo.Save(newProduct)
}

func (s *ProductService) UpdateProduct(productId string, name string, price float64) (*model.Product, error) {
	product := &model.Product{
		ID:    productId,
		Name:  name,
		Price: price,
	}
	err := s.repo.Update(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) DeleteProduct(id int) error {
	return s.repo.Delete(id)
}
