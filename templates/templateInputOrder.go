package templates

import (
	"fmt"
	"inventory_app/helper"
	"inventory_app/models"
	"strings"
	"time"

	tm "github.com/buger/goterm"
)

func (p *productTemplate) InputOrder() {
	totals := tm.NewTable(0, 10, 5, ' ', 0)
	var typeOrder, custName, email, phone, ch string
	helper.ClearScreen()
	fmt.Println("========================")
	fmt.Println("Form Input Order")
	fmt.Println("========================")
	fmt.Println("Tipe Order\n1. Purchase Order\n2. Sales Order")
	fmt.Print("Pilih tipe order: ")
	fmt.Scan(&ch)
	switch ch {
	case "1":
		typeOrder = "PO"
	case "2":
		typeOrder = "SO"
	default:
		fmt.Print("Masukkan pilihan yang benar!")
		helper.BackHandler()
		helper.ClearScreen()
		p.InputOrder()
	}

	fmt.Print("Masukkan Nama Anda: ")
	fmt.Scan(&custName)
	fmt.Print("Masukkan Email Anda: ")
	fmt.Scan(&email)
	fmt.Print("Masukkan Nomor Telpon Anda: ")
	fmt.Scan(&phone)
	idOrder, err := p.productHandler.InsertOrderHandler(typeOrder, custName, email, phone)
	// fmt.Println("Ini kebaca!", lastOrD)
	if err != nil {
		panic(err)
	}
	fmt.Print(idOrder)
	products, err := p.productHandler.GetProduct()
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(totals, "ID\tName\tPrice\tStock")
	for _, v := range products {
		id, name, price, stock, _ := v.GetAll()
		fmt.Fprintf(totals, "%v\t%v\t%.f\t%d\n", *id, *name, *price, *stock)
	}
	helper.ClearScreen()
	var count int
	var product models.Product
	var priceProd float64
	var repeat string

	if typeOrder == "PO" {
		var ODs []models.OrderDetail
		for {
			var OD models.OrderDetail
			var P models.Product
			fmt.Println("====================================\nProducts\n====================================")

			fmt.Println(totals)
			fmt.Println("====================================\nPurchase Order\n====================================")
			item := inputItem()
			fmt.Print("Masukkan harga barang yang ingin diinput: ")
			fmt.Scan(&priceProd)
			productCheck, err := p.productHandler.GetProductByNameAndPrice(item, priceProd)
			if err != nil {
				panic(err)
			}
			fmt.Print(productCheck)
			fmt.Print("Masukkan jumlah yang ingin diinput: ")
			fmt.Scan(&count)
			if *productCheck.GetID() == 0 {
				id, name, price, stock, created_at := productCheck.GetAll()
				P.SetProduct(id, name, price, stock, created_at)
				total := p.productHandler.TotalPrice(P.GetPrice(), count)
				OD.SetDetail(idOrder, *P.GetID(), *P.GetProductName(), *P.GetPrice(), count, total, time.Now())

				ODs = append(ODs, OD)
			}

			fmt.Println("Apakah anda akan memasukkan kembali data yang diinginkan?[y/t]")
			fmt.Scan(&repeat)
			if strings.ToLower(repeat) == "y" {
				continue
			} else if strings.ToLower(repeat) == "t" {
				break
			}
		}
		for _, v := range ODs {
			_, order_id, product_id, product_name, price, quantity, _, _ := v.GetDetail()
			p.productHandler.InsertOrderDetail(*order_id, *product_id, *product_name, *price, *quantity)
		}
	} else if typeOrder == "SO" {
		var SO []models.OrderDetail
		fmt.Println("====================================\nSales Order\n====================================")
		item := inputItem()
		fmt.Print("Masukkan harga barang yang diinput: ")
		fmt.Scan(&priceProd)

	}
	id, name, price, stock, _ := product.GetAll()
	fmt.Printf("%v\t%v\t%.f\t%v", *id, *name, *price, *stock)
	p.productHandler.InsertOrderDetail(idOrder, *id, *name, *price, count)
	total := float64(count) * *price
	p.productHandler.UpdateOrderAfterDetailHandler(count, total, idOrder)
	helper.BackHandler()
	Menu(p.db)
}

func inputItem() string {
	var nameProd string
	fmt.Print("Masukkan nama barang yang ingin diinput: ")
	fmt.Scan(&nameProd)
	return nameProd
}
