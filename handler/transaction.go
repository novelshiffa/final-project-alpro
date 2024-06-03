package handler

import (
	"fmt"

	"github.com/novelshiffa/final-project-alpro/types"
	"github.com/novelshiffa/final-project-alpro/utils"
)

func TransactionHandler(t *types.Transactions) bool {
	var stopLoop bool = false
	var menu types.Menu

	menu.DefaultSelectedColor = "blue"
	menu.Items[0] = types.NewText("[1] Create new transaction")
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

	var txt types.Text = types.NewText("Transactions")

	menu.Listen(&selected, &stopLoop, &cls, func() {
		switch selected {
		case 0:
			if CreateNewTransaction(t) {
				fmt.Println("OK")
			}

			stopLoop = false
		case 1:
			stopLoop = !ViewAllTransactions(t)
		case 2:
			if EditTransaction(t) {
				fmt.Println()
			}
			stopLoop = false
		case 3:
			if DeleteTransaction(t) {
				fmt.Println()
			}
			stopLoop = false
		case 4:
			backToHome = true
		}
	}, func() {
		fmt.Println(txt.Colored)
	})

	return backToHome
}

func CreateNewTransaction(t *types.Transactions) bool {
	if t.Length == types.NMAX {
		panic("Penuh")
	}

	// var transaction types.Transaction

	// fmt.Print("Transaction name: ")
	// fmt.Scanln(&p.Name)

	// InputInteger("Item price: ", &p.Price, true)

	// InputInteger("Item stock: ", &p.Stock, true)

	// fmt.Print("Item category: ")
	// fmt.Scanln(&p.Category)

	// _, err := items.AddNew(p)

	return true
}

func ViewAllTransactions(t *types.Transactions) bool {
	var stopLoop bool
	var menu types.Menu

	menu.DefaultSelectedColor = "blue"
	menu.Items[0] = types.NewText("[1] Back to /items")
	menu.Items[1] = types.NewText("[2] Exit")

	menu.Length = 2
	menu.SetSelected(0)

	var backToItems bool = false

	var selected int

	var cls bool = false

	menu.Listen(&selected, &stopLoop, &cls, func() {
		switch selected {
		case 0:
			backToItems = true
		case 2:
			stopLoop = true
		}
	}, func() {
		utils.ClearTerminal()
		t.ShowInTable()
	})

	return backToItems
}

func EditTransaction(t *types.Transactions) bool {
	return true
}

func DeleteTransaction(t *types.Transactions) bool {
	return true
}
