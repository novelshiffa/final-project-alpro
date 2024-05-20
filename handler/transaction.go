package handler

import (
	"github.com/novelshiffa/final-project-alpro/types"
)

func TransactionHandler() bool {
	var stopLoop bool
	var menu types.Menu

	menu.DefaultSelectedColor = "blue"
	menu.Items[0] = types.NewText("[1] Create transaction")
	menu.Items[1] = types.NewText("[2] View all transactions")
	menu.Items[2] = types.NewText("[3] Edit transaction")
	menu.Items[3] = types.NewText("[4] Delete transaction")
	menu.Items[4] = types.NewText("[5] Back home")
	menu.Items[5] = types.NewText("[6] Exit")

	menu.Length = 6
	menu.SetSelected(0)

	var backToHome bool = false

	var selected int

	var cls bool = true

	menu.Listen(&selected, &stopLoop, &cls, func() {
		switch selected {
		case 4:
			backToHome = true
		}
	}, func() {})

	return backToHome
}
