package repository

import (
	"database/sql"
	"errors"
	"test/domain"
)

type ProductRepository struct {
	DB *sql.DB
}

func (repo *ProductRepository) GetProductByName(product domain.Product) error {
	query := "SELECT COUNT(*) FROM products WHERE productname = $1"
	var count int
	err := repo.DB.QueryRow(query, product.ProductName).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("product already exists")
	}
	return nil
}

func (repo *ProductRepository) CreateProduct(product domain.Product) error {
	query := "INSERT INTO products (productname, price, category) VALUES ($1, $2, $3)"
	_, err := repo.DB.Exec(query, product.ProductName, product.Price, product.Catergory)
	if err != nil {
		return err
	}
	return nil
}

func (repo *ProductRepository) GetAllProduct() ([]domain.Product, error) {
	query := "SELECT productname, price, category FROM products"
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []domain.Product{}
	for rows.Next() {
		product := domain.Product{}
		err := rows.Scan(&product.ProductName, &product.Price, &product.Catergory)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (repo *ProductRepository) DeleteProduct(product domain.Product) error {
	query := "DELETE FROM products WHERE productname = $1"
	_, err := repo.DB.Exec(query, product.ProductName)
	if err != nil {
		return err
	}
	return nil
}
