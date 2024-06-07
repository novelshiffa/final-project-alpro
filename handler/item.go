package handler

import (
	"fmt"

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
			stopLoop = !ViewAllItems(items)
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

func ViewAllItems(items *types.Items) bool {
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
