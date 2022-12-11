package repository

import (
	"context"
	"database/sql"
	"fmt"
	"inventory_app/models"
)

type ProductInterface interface {
	GetAll(ctx context.Context) ([]models.Product, error)
	UpdateAddStock(ctx context.Context, add int, name string, price float64)
	UpdateReduceStock(ctx context.Context, red int, name string, price float64)
	GetByID(ctx context.Context, id int) (models.Product, error)
	GetByNameAndPrice(ctx context.Context, name string, price float64) (models.Product, error)
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *productRepository {
	return &productRepository{db}
}

// id,name,price,stock,created_at
func (repo *productRepository) GetAll(ctx context.Context) ([]models.Product, error) {
	var query string = "SELECT id,name,price,stock,created_at FROM products"
	var products []models.Product

	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return products, err
	}

	for rows.Next() {
		var product models.Product
		rows.Scan(product.GetID(), product.GetProductName(), product.GetPrice(), product.GetStock(), product.GetCreated())
		products = append(products, product)
	}

	return products, nil
}

func (repo *productRepository) GetByNameAndPrice(ctx context.Context, name string, productPrice float64) (models.Product, error) {
	var product models.Product
	var query string = "SELECT id, `name`, price, stock, created_at FROM products where `name` = ? AND price = ?"

	row := repo.db.QueryRowContext(ctx, query, name, productPrice)
	id, names, price, stock, created_at := product.GetAll()
	row.Scan(&id, &names, &price, &stock, &created_at)
	product.SetProduct(id, names, price, stock, created_at)
	return product, nil
}

// This function is for reducing stock while making purchase order. This function need reduce val, name of product and price of the product
func (repo *productRepository) UpdateAddStock(ctx context.Context, add int, name string, productPrice float64) {
	var query string = "UPDATE products SET stock = ? WHERE `name` = ?  AND `price` = ?"

	product, _ := repo.GetByNameAndPrice(ctx, name, productPrice)
	var stock int = *product.GetStock()
	stock += add

	_, err := repo.db.ExecContext(ctx, query, stock, name, productPrice)
	if err != nil {
		panic(err)
	}
}

// This function is for reducing stock while making sales order. This function need reduce val, name of product and price of the product
func (repo *productRepository) UpdateReduceStock(ctx context.Context, red int, name string, productPrice float64) {
	var query string = "UPDATE products SET stock = ? WHERE `name` = ?  AND `price` = ?"

	product, _ := repo.GetByNameAndPrice(ctx, name, productPrice)
	var stock int = *product.GetStock()
	stock -= red

	_, err := repo.db.ExecContext(ctx, query, stock, name, productPrice)
	if err != nil {
		panic(err)
	}
}

func (repo *productRepository) GetByID(ctx context.Context, id int) (models.Product, error) {
	var product models.Product
	var query string = "SELECT id,`name`,price,stock FROM products WHERE id = ?"
	ids, name, price, stock, _ := product.GetAll()

	repo.db.QueryRowContext(ctx, query, id).Scan(ids, name, price, stock)
	fmt.Println(product)
	return product, nil
}
