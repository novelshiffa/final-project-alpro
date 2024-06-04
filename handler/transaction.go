package handler

import (
	"fmt"

	"github.com/novelshiffa/final-project-alpro/types"
	"github.com/novelshiffa/final-project-alpro/utils"
)

func TransactionHandler(t *types.Transactions, i *types.Items) bool {
	var stopLoop bool = false
	var menu types.Menu

	menu.DefaultSelectedColor = "blue"
	menu.Items[0] = types.NewText("[1] Create new transaction")
	menu.Items[1] = types.NewText("[2] View all transactions")
	menu.Items[2] = types.NewText("[3] Edit transaction")
	menu.Items[3] = types.NewText("[4] Delete transaction")
	menu.Items[4] = types.NewText("[5] Back home")
	menu.Items[5] = types.NewText("[6] Exit program")

	menu.Length = 6
	menu.SetSelected(0)

	var backToHome bool = false

	var selected int

	var cls bool = true

	var txt types.Text = types.NewText("Transactions")

	menu.Listen(&selected, &stopLoop, &cls, func() {
		switch selected {
		case 0:
			if CreateNewTransaction(t, i) {
				fmt.Println("OK")
			}

			stopLoop = false
		case 1:
			stopLoop = !ViewAllTransactions(t)
		case 2:
			if EditTransaction(t, i) {
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

func CreateNewTransaction(t *types.Transactions, i *types.Items) bool {
	if t.Length == types.NMAX {
		panic("Penuh")
	}

	var transaction types.Transaction

	InputTime("Transaction time: ", &transaction.Time, true)
	InputTransactionType("Type of transaction: ", &transaction.Type, true)
	InputItem("Enter item id: ", &transaction.Item, true, i)
	InputInteger("Enter quantity: ", &transaction.Quantity, true)

	t.CreateNew(transaction)

	return true
}

func ViewAllTransactions(t *types.Transactions) bool {
	var stopLoop bool
	var menu types.Menu

	menu.DefaultSelectedColor = "blue"
	menu.Items[0] = types.NewText("[1] Back to /items")
	menu.Items[1] = types.NewText("[2] Exit program")

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

func EditTransaction(t *types.Transactions, i *types.Items) bool {
	var id int
	var index int
	var found bool

	var errText = types.NewText("Item not found. Try again.")
	errText.SetColor("red")

	for !found {
		InputInteger("Enter item id (0 to exit): ", &id, true)

		if id == 0 {
			return true
		} else {
			index = t.FindById(id)

			if index != -1 {
				found = true
			} else {
				fmt.Println(errText.Colored)
			}
		}
	}

	// Input transaction time
	var OldValueFormat = func(oldValue string) string {
		text := types.NewText("[!] Current value: " + oldValue + "\n")
		text.SetColor("green")

		return text.Colored
	}

	InputTime(
		OldValueFormat(t.Items[index].Time.String())+
			"Transaction time (Press Enter if you don't want to edit this attribute): ",
		&t.Items[index].Time,
		false,
	)

	// Input transaction type
	InputTransactionType(
		OldValueFormat(t.Items[index].Type)+
			"Type of transaction (Press Enter if you don't want to edit this attribute): ",
		&t.Items[index].Type,
		false,
	)

	// Input item id
	InputItem(
		OldValueFormat(fmt.Sprint(t.Items[index].Item.Id))+
			"Enter item id (Press Enter if you don't want to edit this attribute): ",
		&t.Items[index].Item,
		false,
		i,
	)

	// Input quantity
	InputInteger(
		OldValueFormat(fmt.Sprint(t.Items[index].Quantity))+
			"Enter quantity (Press Enter if you don't want to edit this attribute): ",
		&t.Items[index].Quantity,
		false,
	)

	return true
}

func DeleteTransaction(t *types.Transactions) bool {
	var id int
	var index int
	var found bool

	var errText = types.NewText("Item not found. Try again.")
	errText.SetColor("red")

	for !found {
		InputInteger("Enter item id (0 to exit): ", &id, true)

		if id == 0 {
			return true
		} else {
			index = t.FindById(id)

			if index != -1 {
				found = true
			} else {
				fmt.Println(errText.Colored)
			}
		}
	}

	_, err := t.Delete(index)
	return err == nil || true
}
