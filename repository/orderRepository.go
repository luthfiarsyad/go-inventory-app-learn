package repository

import (
	"context"
	"database/sql"
	"fmt"
	"inventory_app/models"
)

type OrderInterface interface {
	GetAll(ctx context.Context) ([]models.Order, error)
	InsertOrder(ctx context.Context, order models.Order) models.Order
	GetLastOrderPOID(ctx context.Context) models.Order
	GetLastOrderSOID(ctx context.Context) (int, string)
	GetOrderIdByOrderNum(ctx context.Context, orderType string, orderNum string) models.Order
	UpdateOrderAfterDetail(ctx context.Context, quantity int, total float64, idOrder int)
}

type orderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *orderRepository {
	return &orderRepository{db}
}

func (repo *orderRepository) GetAll(ctx context.Context) ([]models.Order, error) {
	var query string = "SELECT id,type,number,customer_name, email,phone, date, quantity, total, created_at FROM orders"
	var orders []models.Order
	rows, _ := repo.db.QueryContext(ctx, query)
	for rows.Next() {
		var order models.Order
		id, typeOrder, number, customerName, email, phone, date, quantity, total, created_at := order.GetOrder()
		rows.Scan(id, typeOrder, number, customerName, email, phone, date, quantity, total, created_at)
		orders = append(orders, order)
	}
	return orders, nil
}

func (repo *orderRepository) InsertOrder(ctx context.Context, order models.Order) models.Order {
	var query string = "INSERT INTO `orders` (`type`,`number`,customer_name,email,phone,`date`,quantity,total) VALUES (?,?,?,?,?,?,?,?)"
	//Date format = 27-11-2022
	_, typeOrder, number, customerName, email, phone, date, quantity, total, _ := order.GetOrder()
	res, err := repo.db.ExecContext(ctx, query, typeOrder, number, customerName, email, phone, *date, quantity, total)
	if err != nil {
		return order
	}
	lastInsertedId, _ := res.LastInsertId()
	id := int(lastInsertedId)
	order.SetID(&id)
	// _, err = stmt.ExecContext(ctx, query, typeOrder, number, customerName, email, phone, date, quantity, total)
	// if err != nil {
	// 	return order, idOrder
	// }

	return order
	// if err != nil {
	// 	return nil
	// }
	// orderID, _ := res.LastInsertId()
	// return &orderID
}

// This function is for getting the last ID and last Purchase Order ID
func (repo *orderRepository) GetLastOrderPOID(ctx context.Context) models.Order {
	var query string = "SELECT id, number FROM orders WHERE type = 'PO' ORDER BY `id` DESC LIMIT 1"
	var order models.Order
	id, _, number, _, _, _, _, _, _, _ := order.GetOrder()
	err := repo.db.QueryRowContext(ctx, query).Scan(&id, &number)
	if err != nil {
		return order
	}
	fmt.Print(id)
	return order
}

// This function is for getting the last ID and last Sales Order ID
func (repo *orderRepository) GetLastOrderSOID(ctx context.Context) (int, string) {
	var LastID int
	var number string
	var query string = "SELECT id, `number` FROM orders WHERE type = 'SO' ORDER BY `id` DESC LIMIT 1"
	repo.db.QueryRowContext(ctx, query).Scan(LastID, number)
	return LastID, number
}

func (repo *orderRepository) GetOrderIdByOrderNum(ctx context.Context, orderType string, orderNum string) models.Order {
	var order models.Order
	var query string = "SELECT id, type, number, customer_name, email, phone, date, quantity, total, created_at FROM orders WHERE TYPE = '?' AND NUMBER = '?'"
	id, typeOrder, number, customerName, email, phone, date, quantity, total, created_at := order.GetOrder()
	repo.db.QueryRowContext(ctx, query, orderType, orderNum).Scan(id, typeOrder, number, customerName, email, phone, date, quantity, total, created_at)

	fmt.Println("Ini id di repo order", orderType, orderNum, order)
	return order
}

func (repo *orderRepository) UpdateOrderAfterDetail(ctx context.Context, quantity int, total float64, idOrder int) {
	var query string = "UPDATE orders SET quantity = ?, total= ? WHERE id=?"
	repo.db.ExecContext(ctx, query, quantity, total, idOrder)
}
