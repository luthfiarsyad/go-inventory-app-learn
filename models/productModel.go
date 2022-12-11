package models

import "time"

type Product struct {
	id         int
	name       string
	price      float64
	stock      int
	created_at time.Time
}

// Atributes of product is id, name, price, stock, created_at
func (p *Product) GetAll() (*int, *string, *float64, *int, *time.Time) {
	return &p.id, &p.name, &p.price, &p.stock, &p.created_at
}

func (p *Product) GetID() *int {
	return &p.id
}

func (p *Product) GetProductName() *string {
	return &p.name
}

func (p *Product) GetPrice() *float64 {
	return &p.price
}

func (p *Product) GetStock() *int {
	return &p.stock
}

func (p *Product) GetCreated() *time.Time {
	return &p.created_at
}

func (p *Product) SetName(name *string) {
	p.name = *name
}

func (p *Product) SetPrice(price *float64) {
	p.price = *price
}

func (p *Product) SetStock(stock *int) {
	p.stock = *stock
}

// // This is used only if needed, im just keeping this in case we needed this :D
func (p *Product) SetProduct(id *int, name *string, price *float64, stock *int, created_at *time.Time) {
	p.id = *id
	p.name = *name
	p.price = *price
	p.stock = *stock
	p.created_at = *created_at
}
