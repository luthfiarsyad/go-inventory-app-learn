package models

import "time"

type OrderDetail struct {
	id           int
	order_id     int
	product_id   int
	product_name string
	price        float64
	quantity     int
	total        float64
	created_at   time.Time
}

// This func returns id, order_id, product_id, product_name, price, quantity, total, created_at
func (od *OrderDetail) GetDetail() (*int, *int, *int, *string, *float64, *int, *float64, *time.Time) {
	id, orderId, productId, productName, price, quantity, total, created_at := od.id, od.order_id, od.product_id, od.product_name, od.price, od.quantity, od.total, od.created_at

	return &id, &orderId, &productId, &productName, &price, &quantity, &total, &created_at
}

// This func sets id, order_id, product_id, product_name, price, quantity, total, created_at to OrderDetail struct
func (od *OrderDetail) SetDetail(orderId int, productId int, productName string, price float64, quantity int, total float64, created_at time.Time) {
	 od.order_id, od.product_id, od.product_name, od.price, od.quantity, od.total, od.created_at = orderId, productId, productName, price, quantity, total, created_at
}

// This func returns OrderDetails ID
func (od *OrderDetail) GetDetailID() *int {
	return &od.id
}
