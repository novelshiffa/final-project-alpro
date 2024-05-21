package handler

import (
	"fmt"

	"github.com/novelshiffa/final-project-alpro/types"
	"github.com/novelshiffa/final-project-alpro/utils"
)

func ProductHandler(products *types.Products) bool {
	var stopLoop bool = false
	var menu types.Menu

	menu.DefaultSelectedColor = "blue"
	menu.Items[0] = types.NewText("[1] Add new product")
	menu.Items[1] = types.NewText("[2] View all products")
	menu.Items[2] = types.NewText("[3] Edit product")
	menu.Items[3] = types.NewText("[4] Delete product")
	menu.Items[4] = types.NewText("[5] Back home")
	menu.Items[5] = types.NewText("[6] Exit")

	menu.Length = 6
	menu.SetSelected(0)

	var backToHome bool = false

	var selected int

	var cls bool = true

	menu.Listen(&selected, &stopLoop, &cls, func() {
		switch selected {
		case 0:
			if AddNewProduct(products) {
				fmt.Println("OK")
			}

			stopLoop = false

		case 1:
			stopLoop = !ViewAllProducts(products)
		case 4:
			backToHome = true
		}
	}, func() {})

	return backToHome
}

func AddNewProduct(products *types.Products) bool {
	if products.Length == types.NMAX {
		panic("Penuh")
	}

	utils.ClearTerminal()
	var p types.Product

	fmt.Print("Product name: ")
	fmt.Scanln(&p.Name)

	fmt.Print("Product price: ")
	fmt.Scanln(&p.Price)

	fmt.Print("Product stock: ")
	fmt.Scanln(&p.Stock)

	fmt.Print("Product category: ")
	fmt.Scanln(&p.Category)

	products.Items[products.Length] = p
	products.Length++

	return true
}

func ViewAllProducts(products *types.Products) bool {
	var stopLoop bool
	var menu types.Menu

	menu.DefaultSelectedColor = "blue"
	menu.Items[0] = types.NewText("[1] Back to /products")
	menu.Items[1] = types.NewText("[2] Exit")

	menu.Length = 2
	menu.SetSelected(0)

	var backToProducts bool = false

	var selected int

	var cls bool = false

	menu.Listen(&selected, &stopLoop, &cls, func() {
		switch selected {
		case 0:
			backToProducts = true
		case 2:
			stopLoop = true
		}
	}, func() {
		utils.ClearTerminal()
		products.ShowInTable()
	})

	return backToProducts
}
