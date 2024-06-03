package main

import (
	"fmt"

	"github.com/novelshiffa/final-project-alpro/handler"
	"github.com/novelshiffa/final-project-alpro/types"
)

// Global Variable Declarations -- Start

var items types.Items
var transactions types.Transactions

// Global Variable Declarations -- End

func main() {
	items = types.Items{
		Items: [types.NMAX]types.Item{
			{Id: 1, Name: "Item1", Category: "Category1", Price: 100, Stock: 10},
			{Id: 2, Name: "Item2", Category: "Category2", Price: 200, Stock: 20},
		},
		Length: 2,
	}

	var stopLoop bool
	var menu types.Menu

	menu.DefaultSelectedColor = "blue"
	menu.Items[0] = types.NewText("[1] Item")
	menu.Items[1] = types.NewText("[2] Transactions")
	menu.Items[2] = types.NewText("[3] Exit")

	menu.Length = 3
	menu.SetSelected(0)

	var selected int

	var cls bool = true

	menu.Listen(&selected, &stopLoop, &cls, func() {
		switch selected {
		case 0:
			stopLoop = !handler.ItemHandler(&items)
		case 1:
			stopLoop = !handler.TransactionHandler(&transactions)
		case 2:
			stopLoop = true
			fmt.Println("Thank you. Good bye!")
		}
	}, func() {})

}
