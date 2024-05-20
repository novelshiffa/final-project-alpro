package main

import (
	"fmt"

	"github.com/novelshiffa/final-project-alpro/handler"
	"github.com/novelshiffa/final-project-alpro/types"
)

// Global Variable Declarations -- Start

var products types.Products
var transactions types.Transactions

// Global Variable Declarations -- End

func main() {
	products = types.Products{
		Items: [types.NMAX]types.Product{
			{Id: 1, Name: "Product1", Category: "Category1", Price: 100, Stock: 10},
			{Id: 2, Name: "Product2", Category: "Category2", Price: 200, Stock: 20},
		},
		Length: 2,
	}
	var stopLoop bool
	var menu types.Menu

	menu.DefaultSelectedColor = "blue"
	menu.Items[0] = types.NewText("[1] Product")
	menu.Items[1] = types.NewText("[2] Transactions")
	menu.Items[2] = types.NewText("[3] Exit")

	menu.Length = 3
	menu.SetSelected(0)

	var selected int

	var cls bool = true

	menu.Listen(&selected, &stopLoop, &cls, func() {
		switch selected {
		case 0:
			stopLoop = !handler.ProductHandler(&products)
		case 1:
			stopLoop = !handler.TransactionHandler()
		case 2:
			stopLoop = true
			fmt.Println("Thank you. Good bye!")
		}
	}, func() {})

}
