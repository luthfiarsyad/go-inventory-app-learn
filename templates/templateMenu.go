package templates

import (
	"database/sql"
	"fmt"
	"inventory_app/handlers"
	"inventory_app/helper"
	"inventory_app/repository"
	"os"
)

func Menu(db *sql.DB) {
	helper.ClearScreen()
	productRepository := repository.NewProductRepository(db)
	orderRepository := repository.NewOrderRepository(db)
	orderDetailRepository := repository.NewOrderDetailRepository(db)
	productHandler := handlers.NewProductHandler(productRepository, orderRepository, orderDetailRepository)

	productTemplate := NewProductHandler(productHandler, db)

	fmt.Println("Menu")
	fmt.Println("=================")
	fmt.Println("1. List Product")
	fmt.Println("2. Input Order")
	fmt.Println("3. View Order")
	fmt.Println("4. Exit")

	var menu int
	fmt.Print("Pilih menu: ")
	fmt.Scanln(&menu)

	switch menu {
	case 1:
		productTemplate.ListProduct()
	case 2:
		productTemplate.InputOrder()
	case 3:
		productTemplate.OrderList()
	case 4:
		os.Exit(0)
	default:
		Menu(db)
	}
}
