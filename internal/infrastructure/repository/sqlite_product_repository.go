package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/kutsuna/simple-go-backend-template/internal/domain/model"
	"github.com/kutsuna/simple-go-backend-template/internal/domain/repository"
)

type sqliteProductRepository struct {
	db *sql.DB
}

func NewSQLiteProductRepository(db *sql.DB) repository.ProductRepository {
	return &sqliteProductRepository{db: db}
}

func (r *sqliteProductRepository) FindByID(id string) (*model.Product, error) {
	var product model.Product
	query := `SELECT id, name, price FROM products WHERE id = ?`
	if err := r.db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Price); err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *sqliteProductRepository) FindAll() ([]*model.Product, error) {
	var products []*model.Product
	query := `SELECT id, name, price FROM products`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}

func (r *sqliteProductRepository) Save(product *model.Product) error {
	product.ID = uuid.New().String()
	query := `INSERT INTO products (id, name, price) VALUES (?, ?, ?)`
	_, err := r.db.Exec(query, product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}

	return nil
}

func (r *sqliteProductRepository) Update(product *model.Product) error {
	query := `UPDATE products SET name = ?, price = ? WHERE id = ?`
	_, err := r.db.Exec(query, product.Name, product.Price, product.ID)
	return err
}

func (r *sqliteProductRepository) Delete(id int) error {
	query := `DELETE FROM products WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}
