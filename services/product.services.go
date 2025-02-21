package services

import (
	"database/sql"
	"gin_gonic_products_rest_api/model"
)

type ProductService struct {
	Db *sql.DB
}

func NewProductService(db *sql.DB) ProductServiceInterface {
	return &ProductService{Db: db}
}

func (p *ProductService) CreateProduct(post model.PostProduct) error {
	_, err := p.Db.Exec(
		"INSERT INTO go_products(Name, Price, Quantity) VALUES ($1, $2, $3)",
		post.Name, post.Price, post.Quantity,
	)
	return err
}

func (p *ProductService) GetAllProducts() ([]model.Product, error) {
	rows, err := p.Db.Query("SELECT id, name, price, quantity FROM go_products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (p *ProductService) GetOneProduct(id uint) (*model.Product, error) {
	var product model.Product
	err := p.Db.QueryRow("SELECT id, name, price, quantity FROM go_products WHERE id=$1", id).Scan(&product.ID, &product.Name, &product.Price, &product.Quantity)

	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductService) UpdateProduct(id uint, product model.PostProduct) error {
	_, err := p.Db.Exec(
		"UPDATE go_products SET name=$1, price=$2, quantity=$3 WHERE id=$4",
		product.Name, product.Price, product.Quantity, id,
	)
	return err
}

func (p *ProductService) DeleteProduct(id uint) error {
	_, err := p.Db.Exec("DELETE FROM go_products WHERE id = $1", id)
	return err
}
