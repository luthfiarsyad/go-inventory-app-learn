package handlers

import (
	"context"
	"fmt"
	"inventory_app/helper"
	"inventory_app/models"
	"inventory_app/repository"
	"time"
)

type ProductHandler interface {
	// Product Methods
	GetProduct() ([]models.Product, error)
	GetProductByNameAndPrice(productName string, productPrice float64) (models.Product, error)
	GetProductByID(prod int) (models.Product, error)
	StockAppend(oldStock *int, newInput int) *int

	// Order Methods
	GetOrders() ([]models.Order, error)
	UpdateOrderAfterDetailHandler(quantity int, total float64, idOrder int)
	InsertOrderHandler(typeOrder string, custName string, email string, phone string) (int, error)
	GetOrderIdByOrNum(orderType string, orderNum string) int

	// Order Details Methods
	InsertOrderDetail(LastID int, product_id int, product_name string, price float64, quantity int)
	GetAllDetailByOrderID(orderId int) []models.OrderDetail
	InsertOrderDetailSlice(LastID int, product_id int, product_name string, price float64, quantity int) models.OrderDetail
	TotalPrice(price *float64, count int) float64
}

type productHandler struct {
	productInterface     repository.ProductInterface
	orderInterface       repository.OrderInterface
	orderDetailInterface repository.OrderDetailInterface
}

func NewProductHandler(productInterface repository.ProductInterface, orderInterface repository.OrderInterface, orderDetailInterface repository.OrderDetailInterface) *productHandler {
	return &productHandler{productInterface, orderInterface, orderDetailInterface}
}

func (p *productHandler) GetProduct() ([]models.Product, error) {
	var ctx context.Context = context.Background()
	products, err := p.productInterface.GetAll(ctx)
	if err != nil {
		return products, err
	}

	return products, nil
}

// This code is for getting all order
func (p *productHandler) GetOrders() ([]models.Order, error) {
	var ctx context.Context = context.Background()
	orders, err := p.orderInterface.GetAll(ctx)
	// var order models.Order
	if err != nil {
		return orders, err
	}
	return orders, nil
}

// This is the code for getting all order details by order id
func (p *productHandler) GetAllDetailByOrderID(orderId int) []models.OrderDetail {
	var ctx context.Context = context.Background()
	fmt.Println("ini order id", orderId)
	orderDetails, err := p.orderDetailInterface.GetDetailByOrderID(ctx, orderId)
	fmt.Println(orderDetails)
	if err != nil {
		panic(err)
	}
	return orderDetails
}

func (p *productHandler) InsertOrderHandler(typeOrder string, custName string, email string, phone string) (int, error) {
	var ctx = context.Background()
	var order models.Order
	var LastID int
	var num string

	fmt.Println("Ini kebaca!", typeOrder)
	if typeOrder == "PO" {
		order = p.orderInterface.GetLastOrderPOID(ctx)
	} else if typeOrder == "SO" {
		LastID, num = p.orderInterface.GetLastOrderSOID(ctx)
	}
	nums := helper.NumberOrderIncrement(num)
	t := time.Now().Format("2006-01-02")
	date, _ := time.Parse("2006-01-02", t)
	order.SetOrder(typeOrder, nums, custName, email, phone, date)
	orderS := p.orderInterface.InsertOrder(ctx, order)
	i, _, _, _, _, _, _, _, _, _ := orderS.GetOrder()
	LastID = *i
	fmt.Print(LastID)
	return LastID, nil
}

func (p *productHandler) InsertOrderDetail(LastID int, product_id int, product_name string, price float64, quantity int) {
	var ctx context.Context = context.Background()
	p.orderDetailInterface.InsertOrderDetail(ctx, LastID, product_id, product_name, price, quantity)
}

func (p *productHandler) GetOrderIdByOrNum(orderType string, orderNum string) int {
	ctx := context.Background()
	order := p.orderInterface.GetOrderIdByOrderNum(ctx, orderType, orderNum)
	id, _, _, _, _, _, _, _, _, _ := order.GetOrder()
	fmt.Println("ini di handler order", *id)
	return *id
}

func (p *productHandler) GetProductByID(prod int) (models.Product, error) {
	ctx := context.Background()
	product, err := p.productInterface.GetByID(ctx, prod)
	if err != nil {
		panic(err)
	}
	return product, nil
}

func (p *productHandler) GetProductByNameAndPrice(productName string, productPrice float64) (models.Product, error) {
	ctx := context.Background()
	product, err := p.productInterface.GetByNameAndPrice(ctx, productName, productPrice)

	return product, err
}

func (p *productHandler) UpdateOrderAfterDetailHandler(quantity int, total float64, idOrder int) {
	ctx := context.Background()
	p.orderInterface.UpdateOrderAfterDetail(ctx, quantity, total, idOrder)
}

func (p *productHandler) InsertOrderDetailSlice(LastID int, product_id int, product_name string, price float64, quantity int) models.OrderDetail {
	var orderDetail models.OrderDetail
	return orderDetail
}

func (p *productHandler) StockAppend(oldStock *int, newInput int) *int {
	var stock int = *oldStock
	res1 := stock + newInput
	var res *int = &res1
	return res
}

func (p *productHandler) TotalPrice(price *float64, count int) float64 {
	newPrice := *price
	total := newPrice * float64(count)
	return total
}
