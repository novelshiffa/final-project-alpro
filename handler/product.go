package handler

import (
	"github.com/novelshiffa/final-project-alpro/types"
)

func ProductHandler(products *types.Products) bool {
	var stopLoop bool
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
		case 1:
			backToHome = AddProductHandler(products)
		case 4:
			backToHome = true
		}
	}, func() {})

	return backToHome
}

func AddProductHandler(products *types.Products) bool {
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
			backToProducts = true
		}
	}, func() {
		products.ShowInTable()
	})

	return backToProducts
}
