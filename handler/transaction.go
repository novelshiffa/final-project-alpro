package handler

import (
	"fmt"
	"strings"

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

	var txt types.Text = types.NewText("/transactions")

	menu.Listen(&selected, &stopLoop, &cls, func() {
		switch selected {
		case 0:
			fmt.Println("OK")
			stopLoop = false
		case 1:
			stopLoop = !ViewAllTransactions(t, "/transactions/view")
		case 2:
			EditTransaction(t, i)
			stopLoop = false
		case 3:
			DeleteTransaction(t)
			stopLoop = false
		case 4:
			backToHome = true
		}
	}, func() {
		fmt.Println(txt.Colored)
	})

	return backToHome
}

func CreateNewTransaction(t *types.Transactions, i *types.Items) {
	if t.Length == types.NMAX {
		panic("Penuh")
	}

	var transaction types.Transaction

	InputTime("Transaction time: ", &transaction.Time, true)
	InputTransactionType("Type of transaction: ", &transaction.Type, true)
	InputItem("Enter item id: ", &transaction.Item, true, i)
	InputInteger("Enter quantity: ", &transaction.Quantity, true)

	t.CreateNew(transaction)
}

func ViewAllTransactions(t *types.Transactions, title string) bool {
	var titleText = types.NewText(title)
	titleText.SetColor("green")

	var stopLoop bool
	var menu types.Menu

	menu.DefaultSelectedColor = "blue"
	menu.Items[0] = types.NewText("[1] Sort by column")
	menu.Items[1] = types.NewText("[2] Filter by column")
	menu.Items[2] = types.NewText("[3] Back to /transactions")
	menu.Items[3] = types.NewText("[3] Exit program")

	menu.Length = 4
	menu.SetSelected(0)

	var backToTransactions bool = false

	var selected int

	var cls bool = false

	menu.Listen(&selected, &stopLoop, &cls, func() {
		switch selected {
		case 0:
			var transactionsCopy types.Transactions
			var column string
			var asc string

			var rightArrowText types.Text = types.NewText("[→] ")
			rightArrowText.SetColor("blue")

			var zeroToCancelText types.Text = types.NewText("(0 to cancel) ")
			zeroToCancelText.SetColor("red")

			InputColumnName("transactions", "Sort by which column?", &column)

			if column == "0" {
				backToTransactions = ViewAllTransactions(t, "/transactions/view")
			} else {
				prompt := types.NewText("Would you like to sort it ascendingly? [Y/N] ")
				prompt.SetColor("white")
				invalidInputErrText := types.NewText("Please input either Y or N.")
				invalidInputErrText.SetColor("red")

				for {
					fmt.Print(rightArrowText.Colored + prompt.Colored + zeroToCancelText.Colored)
					fmt.Scanln(&asc)

					exitLoop := false // Variable to control the loop

					switch strings.ToLower(asc) {
					case "y":
						transactionsCopy = t.SortBy(column, true)
						exitLoop = true
					case "n":
						transactionsCopy = t.SortBy(column, false)
						exitLoop = true
					case "0":
						backToTransactions = ViewAllTransactions(t, title)
						return
					default:
						fmt.Println(invalidInputErrText.Colored)
					}

					if exitLoop {
						backToTransactions = ViewAllTransactions(&transactionsCopy, title+fmt.Sprintf(" sortBy=%s&asc=%t", column, asc == "y"))
						return // Exit the loop and the function
					}
				}

			}
		case 2:
			backToTransactions = true
		case 3:
			stopLoop = true
		}
	}, func() {
		utils.ClearTerminal()
		fmt.Println(titleText.Colored)
		t.ShowInTable()
	})

	return backToTransactions
}

func EditTransaction(t *types.Transactions, i *types.Items) {
	var id int
	var index int
	var found bool

	var errText = types.NewText("Item not found. Try again.")
	errText.SetColor("red")

	for !found {
		InputInteger("Enter item id (0 to exit): ", &id, true)

		if id == 0 {
			return
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
		text.SetColor("blue")

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
}

func DeleteTransaction(t *types.Transactions) {
	var id int
	var index int
	var found bool

	var errText = types.NewText("Item not found. Try again.")
	errText.SetColor("red")

	for !found {
		InputInteger("Enter item id (0 to exit): ", &id, true)

		if id == 0 {
			return
		} else {
			index = t.FindById(id)

			if index != -1 {
				found = true
			} else {
				fmt.Println(errText.Colored)
			}
		}
	}
}
