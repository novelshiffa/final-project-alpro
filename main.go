package main

import (
	"fmt"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/novelshiffa/final-project-alpro/handler"
	"github.com/novelshiffa/final-project-alpro/types"
	"github.com/novelshiffa/final-project-alpro/utils"
)

// Global Variable Declarations -- Start

var products types.Product
var transactions types.Transactions
var menu types.Menu
var stopLoop bool

// Global Variable Declarations -- End

func main() {
	menu.DefaultSelectedColor = "blue"
	menu.Items[0] = types.NewText("[1] Product")
	menu.Items[1] = types.NewText("[2] Transactions")
	menu.Items[2] = types.NewText("[3] Exit")

	menu.Length = 3
	menu.SetSelected(0)

	var selected int

	for {
		utils.ClearTerminal()

		menu.ShowAll()

		keyboard.Listen(func(key keys.Key) (stop bool, err error) {
			if key.Code == keys.Up && selected > 0 {
				selected--
				menu.SetSelected(selected)
			} else if key.Code == keys.Down && selected < menu.Length-1 {
				selected++
				menu.SetSelected(selected)
			}

			if key.Code == keys.Enter {
				stopLoop = true
			}

			return true, nil
		})

		if stopLoop {
			switch selected {
			case 0:
				handler.ProductHandler()
			case 1:
				handler.TransactionHandler()
			case 2:
				fmt.Println("Thank you. Goodbye!")
			}
			break
		}
	}

}
