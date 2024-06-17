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

	var txt types.Text = types.NewText("/items")
	txt.SetColor("green")

	menu.Listen(&selected, &stopLoop, &cls, func() {
		switch selected {
		case 0:
			AddNewItem(items)
			stopLoop = false
		case 1:
			stopLoop = !ViewAllItems(items, "/items/view")
		case 2:
			EditItem(items)
			stopLoop = false
		case 3:
			DeleteItem(items)
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

func AddNewItem(items *types.Items) {
	if items.Length == types.NMAX {
		panic("Penuh")
	}

	var p types.Item

	var rightArrowText types.Text = types.NewText("[→] ")
	rightArrowText.SetColor("blue")

	minusOneToCancel := "(Type -1 to cancel)"
	prompt := types.NewText(fmt.Sprintf("Enter item name %s: ", minusOneToCancel))
	prompt.SetColor("white")

	fmt.Print(rightArrowText.Colored + prompt.Colored)
	InputlnString(&p.Name)

	if p.Name == "-1" {
		return
	}

	prompt.SetValue("Enter item price " + minusOneToCancel + ": ")
	InputInteger(rightArrowText.Colored+prompt.Colored, &p.Price, true)

	if p.Price == -1 {
		return
	}

	prompt.SetValue("Enter item stock " + minusOneToCancel + ": ")
	InputInteger(rightArrowText.Colored+prompt.Colored, &p.Stock, true)

	if p.Stock == -1 {
		return
	}

	prompt.SetValue(fmt.Sprintf("Enter item category %s: ", minusOneToCancel))
	fmt.Print(rightArrowText.Colored + prompt.Colored)
	InputlnString(&p.Category)

	if p.Category == "-1" {
		return
	}

	items.AddNew(p)
}

func ViewAllItems(items *types.Items, title string) bool {
	var titleText = types.NewText(title)
	titleText.SetColor("green")

	var stopLoop bool
	var menu types.Menu

	menu.DefaultSelectedColor = "blue"
	menu.Items[0] = types.NewText("[1] Sort by column")
	menu.Items[1] = types.NewText("[2] Filter by column")
	menu.Items[2] = types.NewText("[3] Back to /items")
	menu.Items[3] = types.NewText("[4] Exit program")

	menu.Length = 4
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

			InputColumnName("items", "Sort by which column?", &column)

			if column == "0" {
				backToItems = ViewAllItems(items, "/items/view")
			} else {
				invalidInputErrText := types.NewText("Please input either Y or N.")
				invalidInputErrText.SetColor("red")

				for {
					fmt.Print(RightArrowedPrompt("Would you like to sort it ascendingly? [Y/N] (0 to cancel) "))
					fmt.Scanln(&asc)

					exitLoop := false // Variable to control the loop

					switch strings.ToLower(asc) {
					case "y":
						itemsCopy = items.SortBy(column, true)
						exitLoop = true
					case "n":
						itemsCopy = items.SortBy(column, false)
						exitLoop = true
					case "0":
						backToItems = ViewAllItems(items, title)
						return
					default:
						fmt.Println(invalidInputErrText.Colored)
					}

					if exitLoop {
						backToItems = ViewAllItems(&itemsCopy, title+fmt.Sprintf(" sortBy=%s&asc=%t", column, asc == "y"))
						return // Exit the loop and the function
					}
				}

			}

		case 1:
			var itemsCopy types.Items
			var column string

			InputColumnName("items", "Filter by which column?", &column)

			if column == "0" {
				backToItems = ViewAllItems(items, title)
			} else {
				column = strings.ToLower(column)
				var temp string
				if column == "id" || column == "stock" || column == "price" {
					var temp2 int
					InputInteger(RightArrowedPrompt("Enter value to filter: "), &temp2, true)
					temp = fmt.Sprintf("%d", temp2)
				} else {
					fmt.Print("Enter value to filter (0 to cancel): ")
					InputlnString(&temp)
				}

				if column != "0" {
					itemsCopy = items.FilterBy(column, temp)
					backToItems = ViewAllItems(&itemsCopy, title+fmt.Sprintf(" %s='%s'", column, temp))
				} else {
					backToItems = ViewAllItems(items, title)
				}
			}
		case 2:
			backToItems = true
		case 3:
			stopLoop = true
			var goodByeText = types.NewText("さよなら！")
			goodByeText.SetColor("green")
			fmt.Println(goodByeText.Colored)
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
		InputInteger(RightArrowedPrompt("Enter item id (0 to cancel) "), &id, true)

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

	fmt.Print(OldValueFormat(items.Items[index].Name))
	fmt.Print(RightArrowedPrompt("Enter new name (Press Enter if you don't want to edit this attribute): "))
	InputlnString(&temp)

	if temp != "" {
		items.Items[index].Name = temp
		temp = ""
	}

	fmt.Print(OldValueFormat(items.Items[index].Category))
	fmt.Print(RightArrowedPrompt("Enter new category (Press Enter if you don't want to edit this attribute): "))
	InputlnString(&temp)

	if temp != "" {
		items.Items[index].Category = temp
		temp = ""
	}

	fmt.Print(OldValueFormat(fmt.Sprintf("%d", items.Items[index].Price)))
	InputInteger(RightArrowedPrompt("Enter new price (Press Enter if you don't want to edit this attribute): "), &items.Items[index].Price, false)

	fmt.Print(OldValueFormat(fmt.Sprintf("%d", items.Items[index].Stock)))
	InputInteger(RightArrowedPrompt("Enter new stock (Press Enter if you don't want to edit this attribute): "), &items.Items[index].Stock, false)
}

func DeleteItem(items *types.Items) {
	var id int
	var index int
	var found bool

	var errText = types.NewText("Item not found. Try again.")
	errText.SetColor("red")

	for !found {
		InputInteger(RightArrowedPrompt("Enter id id (0 to cancel): "), &id, true)

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
