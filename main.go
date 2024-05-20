package main

import (
	"fmt"

	"github.com/novelshiffa/final-project-alpro/handler"
	"github.com/novelshiffa/final-project-alpro/types"
)

// Global Variable Declarations -- Start

var products types.Product
var transactions types.Transactions

// Global Variable Declarations -- End

func main() {
	var stopLoop bool
	var menu types.Menu

	menu.DefaultSelectedColor = "blue"
	menu.Items[0] = types.NewText("[1] Product")
	menu.Items[1] = types.NewText("[2] Transactions")
	menu.Items[2] = types.NewText("[3] Exit")

	menu.Length = 3
	menu.SetSelected(0)

	var selected int

	menu.Listen(&selected, &stopLoop, func() {
		switch selected {
		case 0:
			stopLoop = !handler.ProductHandler()
		case 1:
			stopLoop = !handler.TransactionHandler()
		case 2:
			stopLoop = true
			fmt.Println("Thank you. Good bye!")
		}
	})

}
