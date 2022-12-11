package repository

import (
	"context"
	"database/sql"
	"fmt"
	"inventory_app/models"
)

type OrderDetailInterface interface {
	GetAll(ctx context.Context, order_id int) ([]models.OrderDetail, error)
	GetDetailByOrderID(ctx context.Context, orderId int) ([]models.OrderDetail, error)
	InsertOrderDetail(ctx context.Context, LastID int, product_id int, product_name string, price float64, quantity int)
}

type orderDetailRepository struct {
	db *sql.DB
}

func NewOrderDetailRepository(db *sql.DB) *orderDetailRepository {
	return &orderDetailRepository{db}
}

func (repo *orderDetailRepository) GetAll(ctx context.Context, order_id int) ([]models.OrderDetail, error) {
	var query string = "SELECT id,order_id,product_id,product_name,price,quantity,total,created_at FROM `order_details` WHERE order_id = ?"
	var ods []models.OrderDetail
	rows, err := repo.db.QueryContext(ctx, query)
	if err != sql.ErrNoRows {
		return nil, err
	}
	for rows.Next() {
		var od models.OrderDetail
		rows.Scan(od)
		ods = append(ods, od)
	}
	return ods, nil
}

func (repo *orderDetailRepository) GetDetailByOrderID(ctx context.Context, orderId int) ([]models.OrderDetail, error) {
	var query string = "SELECT id, order_id, product_id, product_name, price, quantity, total, created_at FROM order_details WHERE order_id = ?"
	var orderDetails []models.OrderDetail

	rows, err := repo.db.QueryContext(ctx, query, orderId)
	if err != sql.ErrNoRows {
		return nil, err
	}
	for rows.Next() {
		var orderDetail models.OrderDetail
		id, order_id, product_id, product_name, price, quantity, total, created_at := orderDetail.GetDetail()
		rows.Scan(id, order_id, product_id, product_name, price, quantity, total, created_at)
		orderDetails = append(orderDetails, orderDetail)
	}
	fmt.Println(orderDetails, "ini di repo orderD line 53")

	return orderDetails, nil
}

func (repo *orderDetailRepository) InsertOrderDetail(ctx context.Context, LastID int, product_id int, product_name string, price float64, quantity int) {
	var query string = "INSERT INTO `order_details` (order_id,product_id,product_name,price,quantity,total) VALUES (?,?,?,?,?,?)"
	total := price * float64(quantity)
	_, err := repo.db.ExecContext(ctx, query, LastID, product_id, product_name, price, quantity, total)
	if err != nil {
		panic(err)
	}
}
