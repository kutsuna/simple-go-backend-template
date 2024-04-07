package repository

import "github.com/kutsuna/simple-go-backend-template/internal/domain/model"

type ProductRepository interface {
	FindByID(id string) (*model.Product, error)
	FindAll() ([]*model.Product, error)
	Save(product *model.Product) error
	Update(product *model.Product) error
	Delete(id int) error
}
