package handler

import (
	"fmt"
	"strings"
	"time"

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
	txt.SetColor("green")

	menu.Listen(&selected, &stopLoop, &cls, func() {
		switch selected {
		case 0:
			CreateNewTransaction(t, i)
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
		case 5:
			stopLoop = true
			var goodByeText = types.NewText("さよなら！")
			goodByeText.SetColor("green")
			fmt.Println(goodByeText.Colored)
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

	var rightArrowText types.Text = types.NewText("[→] ")
	rightArrowText.SetColor("blue")

	prompt := types.NewText("Enter transaction time ( Type 1011-01-01 01:01:01 to cancel ): ")
	prompt.SetColor("white")

	InputTime(rightArrowText.Colored+prompt.Colored, &transaction.Time, true, false)

	if transaction.Time.Format("2024-01-01 01:01:01") == "1011-01-01 01:01:01" {
		return
	}

	prompt.SetValue("Enter type of transaction (Type 0 to cancel): ")
	InputTransactionType(rightArrowText.Colored+prompt.Colored, &transaction.Type, true, "0")

	if transaction.Type == "0" || transaction.Type == "" {
		return
	}

	prompt.SetValue("Enter item id (Type 0 to cancel): ")
	InputItem(rightArrowText.Colored+prompt.Colored, &transaction.Item, false, i)

	if transaction.Item.Id == 0 {
		return
	}

	prompt.SetValue("Enter quantity (Type -1 to cancel): ")
	InputInteger(rightArrowText.Colored+prompt.Colored, &transaction.Quantity, false)

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
	menu.Items[3] = types.NewText("[4] Exit program")

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

			InputColumnName("transactions", "Sort by which column?", &column)

			if column == "0" {
				backToTransactions = ViewAllTransactions(t, "/transactions/view")
			} else {
				invalidInputErrText := types.NewText("Please input either Y or N.")
				invalidInputErrText.SetColor("red")

				for {
					fmt.Print(RightArrowedPrompt("Would you like to sort it ascendingly? [Y/N] (0 to cancel) "))
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
		case 1:
			var transactionsCopy types.Transactions
			var column string

			InputColumnName("transactions", "Filter by which column?", &column)

			if column == "0" {
				backToTransactions = ViewAllTransactions(t, title)
			} else {
				column = strings.ToLower(column)
				var temp string
				if column == "id" || column == "itemid" || column == "quantity" {
					var temp2 int
					InputInteger(RightArrowedPrompt("Enter value to filter: "), &temp2, true)
					temp = fmt.Sprintf("%d", temp2)
				} else if column == "time" {
					var timeTemp time.Time
					var dateOnly string

					invalidInputErrText := types.NewText("Please input either Y or N.")
					invalidInputErrText.SetColor("red")

					for !(dateOnly == "y" || dateOnly == "n") {
						fmt.Print(RightArrowedPrompt("Would you like to filter it by date (not date and time)? [Y/N] (0 to cancel) "))
						fmt.Scanln(&dateOnly)

						dateOnly = strings.ToLower(dateOnly)
					}

					InputTime("Enter time to filter: ", &timeTemp, true, dateOnly == "y")

					if dateOnly == "y" {
						temp = timeTemp.Format("2006-01-02")
					} else {
						temp = timeTemp.Format("2006-01-02 15:04:05")
					}

				} else {
					fmt.Print(RightArrowedPrompt("Enter value to filter (0 to cancel): "))
					InputlnString(&temp)
				}

				if column != "0" {
					transactionsCopy = t.FilterBy(column, temp)
					backToTransactions = ViewAllTransactions(&transactionsCopy, title+fmt.Sprintf(" %s='%s'", column, temp))
				} else {
					backToTransactions = ViewAllTransactions(t, title)
				}
			}
		case 2:
			backToTransactions = true
		case 3:
			stopLoop = true

			var goodByeText = types.NewText("さよなら！")
			goodByeText.SetColor("green")
			fmt.Println(goodByeText.Colored)
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
		InputInteger(RightArrowedPrompt("Enter item id (0 to cancel): "), &id, true)

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
	InputTime(
		OldValueFormat(t.Items[index].Time.String())+
			RightArrowedPrompt("Transaction time (Press Enter if you don't want to edit this attribute): "),
		&t.Items[index].Time,
		false,
		false,
	)

	// Input transaction type
	InputTransactionType(
		OldValueFormat(t.Items[index].Type)+
			RightArrowedPrompt("Type of transaction (Press Enter if you don't want to edit this attribute): "),
		&t.Items[index].Type,
		false, "",
	)

	// Input item id
	InputItem(
		OldValueFormat(fmt.Sprint(t.Items[index].Item.Id))+
			RightArrowedPrompt("Enter item id (Press Enter if you don't want to edit this attribute): "),
		&t.Items[index].Item,
		false,
		i,
	)

	// Input quantity
	InputInteger(
		OldValueFormat(fmt.Sprint(t.Items[index].Quantity))+
			RightArrowedPrompt("Enter quantity (Press Enter if you don't want to edit this attribute): "),
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
		InputInteger(RightArrowedPrompt("Enter item id (0 to cancel): "), &id, true)

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
