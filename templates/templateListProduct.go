package templates

import (
	"database/sql"
	"fmt"
	"inventory_app/handlers"
	"inventory_app/helper"

	tm "github.com/buger/goterm"
)

type productTemplate struct {
	productHandler handlers.ProductHandler
	db             *sql.DB
}

func NewProductHandler(productHandler handlers.ProductHandler, db *sql.DB) *productTemplate {
	return &productTemplate{productHandler, db}
}

func (p *productTemplate) ListProduct() {
	helper.ClearScreen()
	totals := tm.NewTable(0, 10, 5, ' ', 0)
	products, err := p.productHandler.GetProduct()
	if err != nil {
		panic(err)
	}
	if len(products) == 0 {
		fmt.Println("Data kosong")
	} else {
		fmt.Fprintf(totals, "\t\tProduct List\n")
		fmt.Fprintf(totals, "ID\tProduct\tPrice\tStock\tCreated_at\n")
		for _, v := range products {
			date := *v.GetCreated()
			conv := date.Format("2006-01-02 15:04:05")
			dateC := conv[:19]
			fmt.Fprintf(totals, "%v\t%v\t%.f\t%v\t%v\t\n", *v.GetID(), *v.GetProductName(), *v.GetPrice(), *v.GetStock(), dateC)
		}
	}
	fmt.Println(totals)
	helper.BackHandler()
	Menu(p.db)
}
