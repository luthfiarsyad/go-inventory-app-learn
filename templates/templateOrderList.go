package templates

import (
	"fmt"
	"inventory_app/helper"
	"time"

	tm "github.com/buger/goterm"
)

func (p *productTemplate) OrderList() {
	var pil string
	totals := tm.NewTable(0, 10, 5, ' ', 0)
	helper.ClearScreen()
	// var typeOrder string
	fmt.Println("List Orders")
	// fmt.Println("Masukkan Pilihan Anda:")
	// fmt.Scan(&typeOrder)
	orders, err := p.productHandler.GetOrders()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(totals, "No\tOrder Number\tName\tEmail\tPhone\tDate Order\tQuantity\tTotal\n")
	for _, v := range orders {
		id, typeOrder, number, customerName, email, phone, date, quantity, total, _ := v.GetOrder()
		var t time.Time = *date
		d := t.Format("2006-01-02")
		dateC := d[:10]
		orderNumber := helper.OrderNumberCombiner(*typeOrder, *number)
		fmt.Fprintf(totals, "%v\t%v\t%v\t%v\t%v\t%v\t%v\t%.f\n", *id, orderNumber, *customerName, *email, *phone, dateC, *quantity, *total)
	}
	fmt.Println(totals)
	fmt.Println("Masukkan order number yang diinginkan:")
	fmt.Scan(&pil)
	orderType, orderNum := helper.OrderNumberSeparator(pil)
	id := p.productHandler.GetOrderIdByOrNum(orderType, orderNum)
	fmt.Println("Ini ID di OrderList", id)
	orderDetails := p.productHandler.GetAllDetailByOrderID(id)
	for i, v := range orderDetails {
		_, _, _, product_name, price, quantity, total, _ := v.GetDetail()
		fmt.Println("No\tProduct\tPrice\tQuantity\tTotal")
		fmt.Printf("%d\t%s\t%.f\t%d\t%.f", i, *product_name, *price, *quantity, *total)
	}
	helper.BackHandler()
	// fmt.Scan()
	// Menu(p.db)
}
