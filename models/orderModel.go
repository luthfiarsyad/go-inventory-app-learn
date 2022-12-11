package models

import "time"

type Order struct {
	id            int
	typeOrder     string
	number        string
	customer_name string
	email         string
	phone         string
	date          time.Time
	quantity      int
	total         float64
	created_at    time.Time
}

func (o *Order) GetID() *int {
	return &o.id
}

// This func for get order of id, typeOrder, number, customerName, email, phone, date, quantity, total, created_at
func (o *Order) GetOrder() (*int, *string, *string, *string, *string, *string, *time.Time, *int, *float64, *time.Time) {
	return &o.id, &o.typeOrder, &o.number, &o.customer_name, &o.email, &o.phone, &o.date, &o.quantity, &o.total, &o.created_at
}

//
func (o *Order) SetOrder(typeOrder string, number string, customer_name string, email string, phone string, date time.Time) {
	o.typeOrder = typeOrder
	o.number = number
	o.customer_name = customer_name
	o.email = email
	o.phone = phone
	o.date = date
}

// func (o *Order) GetType() *string {
// 	return &o.typeOrder
// }

// func (o *Order) GetNumber() *string {
// 	return &o.number
// }

// func (o *Order) GetCustomerName() *string {
// 	return &o.customer_name
// }

// func (o *Order) GetEmail() *string {
// 	return &o.email
// }

// func (o *Order) GetPhone() *string {
// 	return &o.phone
// }

// func (o *Order) GetDate() *time.Time {
// 	return &o.date
// }

// func (o *Order) GetQuantity() *int {
// 	return &o.quantity
// }

// func (o *Order) GetTotal() *float64 {
// 	return &o.total
// }

// func (o *Order) GetCreatedAt() *time.Time {
// 	return &o.created_at
// }

func (o *Order) SetID(id *int) {
	o.id = *id
}

// func (o *Order) SetType(typeOrder *string) {
// 	o.typeOrder = *typeOrder
// }

// func (o *Order) SetNumber(number *string) {
// 	o.typeOrder = *number
// }

// func (o *Order) SetCustomerName(custName *string) {
// 	o.typeOrder = *custName
// }

// func (o *Order) SetEmail(email *string) {
// 	o.typeOrder = *email
// }

// func (o *Order) SetPhone(phone *string) {
// 	o.typeOrder = *phone
// }

// func (o *Order) SetDate(date *time.Time) {
// 	o.date = *date
// }

// func (o *Order) SetQuantity(quantity *int) {
// 	o.quantity = *quantity
// }

// func (o *Order) SetTotal(total *float64) {
// 	o.total = *total
// }

// func (o *Order) SetCreatedAt(created_at *time.Time) {
// 	o.created_at = *created_at
// }
