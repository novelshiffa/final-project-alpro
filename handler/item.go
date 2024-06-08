package handler

import (
	"fmt"
	"strings"

	"github.com/novelshiffa/final-project-alpro/types"
	"github.com/novelshiffa/final-project-alpro/utils"
)

func ItemHandler(items *types.Items) bool {
	var stopLoop bool = false
	var menu types.Menu

	menu.DefaultSelectedColor = "blue"
	menu.Items[0] = types.NewText("[1] Add new item")
	menu.Items[1] = types.NewText("[2] View all items")
	menu.Items[2] = types.NewText("[3] Edit item")
	menu.Items[3] = types.NewText("[4] Delete item")
	menu.Items[4] = types.NewText("[5] Back home")
	menu.Items[5] = types.NewText("[6] Exit program")

	menu.Length = 6
	menu.SetSelected(0)

	var backToHome bool = false

	var selected int

	var cls bool = true

	var txt types.Text = types.NewText("Items")

	menu.Listen(&selected, &stopLoop, &cls, func() {
		switch selected {
		case 0:
			AddNewItem(items)
			stopLoop = false
		case 1:
			stopLoop = !ViewAllItems(items, "items")
		case 2:
			EditItem(items)
			stopLoop = false
		case 3:
			DeleteItem(items)
			stopLoop = false
		case 4:
			backToHome = true
		}
	}, func() {
		fmt.Println(txt.Colored)
	})

	return backToHome
}

func AddNewItem(items *types.Items) {
	if items.Length == types.NMAX {
		panic("Penuh")
	}

	var p types.Item

	fmt.Print("Item name: ")
	fmt.Scanln(&p.Name)

	InputInteger("Item price: ", &p.Price, true)

	InputInteger("Item stock: ", &p.Stock, true)

	fmt.Print("Item category: ")
	fmt.Scanln(&p.Category)

	items.AddNew(p)
}

func ViewAllItems(items *types.Items, title string) bool {
	var titleText = types.NewText(title)
	titleText.SetColor("green")

	var stopLoop bool
	var menu types.Menu

	menu.DefaultSelectedColor = "blue"
	menu.Items[0] = types.NewText("[1] Sort by column")
	menu.Items[1] = types.NewText("[2] Back to /items")
	menu.Items[2] = types.NewText("[3] Exit program")

	menu.Length = 3
	menu.SetSelected(0)

	var backToItems bool = false
	var selected int
	var cls bool = false

	menu.Listen(&selected, &stopLoop, &cls, func() {
		switch selected {
		case 0:
			var itemsCopy types.Items
			var column string
			var asc string

			var rightArrowText types.Text = types.NewText("[→] ")
			rightArrowText.SetColor("blue")

			var zeroToCancelText types.Text = types.NewText("(0 to cancel) ")
			zeroToCancelText.SetColor("red")

			var prompt types.Text = types.NewText("Sort by which column? ")
			prompt.SetColor("white")

			var invalidInputErrText types.Text = types.NewText("Undefined column name. Try again.")
			invalidInputErrText.SetColor("red")

			var stopInput bool = false
			for !stopInput {
				fmt.Print(rightArrowText.Colored + prompt.Colored + zeroToCancelText.Colored)
				fmt.Scanln(&column)

				if items.IsColumn(column) || column == "0" {
					stopInput = true
				} else {
					fmt.Println(invalidInputErrText.Colored)
				}
			}

			if column == "0" {
				backToItems = ViewAllItems(items, "items")
			} else {
				prompt := types.NewText("Would you like to sort it ascendingly? [Y/N] ")
				prompt.SetColor("white")
				invalidInputErrText := types.NewText("Please input either Y or N.")
				invalidInputErrText.SetColor("red")

				for {
					fmt.Print(rightArrowText.Colored + prompt.Colored + zeroToCancelText.Colored)
					fmt.Scanln(&asc)

					switch strings.ToLower(asc) {
					case "y":
						itemsCopy = items.SortBy(column, true)
					case "n":
						itemsCopy = items.SortBy(column, false)
					case "0":
						backToItems = ViewAllItems(items, "items")
						return
					default:
						fmt.Println(invalidInputErrText.Colored)
						continue
					}

					backToItems = ViewAllItems(&itemsCopy, fmt.Sprintf("items?sortBy=%s&asc=%t", column, asc == "y"))
					break
				}
			}

		case 1:
			backToItems = true
		case 2:
			stopLoop = true
		}
	}, func() {
		utils.ClearTerminal()
		fmt.Println(titleText.Colored)
		items.ShowInTable()
	})

	return backToItems
}

func EditItem(items *types.Items) {
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
			index = items.FindById(id)

			if index != -1 {
				found = true
			} else {
				fmt.Println(errText.Colored)
			}
		}
	}

	var temp string

	fmt.Print("Enter new name (Press Enter if you don't want to edit this attribute): ")
	fmt.Scanln(&temp)

	if temp != "" {
		items.Items[index].Name = temp
		temp = ""
	}

	fmt.Print("Enter new category (Press Enter if you don't want to edit this attribute): ")
	fmt.Scanln(&temp)

	if temp != "" {
		items.Items[index].Category = temp
		temp = ""
	}

	InputInteger("Enter new price (Press Enter if you don't want to edit this attribute): ", &items.Items[index].Price, false)
	InputInteger("Enter new stock (Press Enter if you don't want to edit this attribute): ", &items.Items[index].Stock, false)
}

func DeleteItem(items *types.Items) {
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
			index = items.FindById(id)

			if index != -1 {
				found = true
			} else {
				fmt.Println(errText.Colored)
			}
		}
	}

	items.Delete(index)
}
