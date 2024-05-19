package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/novelshiffa/final-project-alpro/handler"
	"github.com/novelshiffa/final-project-alpro/types"
)

// Global Variable Declarations -- Start

var products types.Product
var transactions types.Transactions
var menu types.Menu
var stopLoop bool

// Global Variable Declarations -- End

func setStopLoop(val bool) {
	stopLoop = val
}

func clearTerminal() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // for Windows
	} else {
		cmd = exec.Command("clear") // for Unix-like systems
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	var i int = 0

	menu.DefaultSelectedColor = "red"
	menu.Items[0] = types.NewText("[1] Product")
	menu.Items[1] = types.NewText("[2] Transactions")
	menu.Items[2] = types.NewText("[3] Exit")

	menu.Length = 3
	menu.SetSelected(0)

	var selected int

	for !stopLoop {
		clearTerminal()

		menu.ShowAll()

		keyboard.Listen(func(key keys.Key) (stop bool, err error) {
			if key.Code == keys.Up && i > 0 && i <= 2 {
				i--
				menu.SetSelected(i)
			} else if key.Code == keys.Down && i >= 0 && i < 2 {
				i++
				menu.SetSelected(i)
			}

			if key.Code == keys.Enter {
				selected = i
				setStopLoop(true)
				//return false, nil
			}

			return true, nil // Return false to continue listening else stop
		})

		if selected == 0 {
			handler.ProductHandler()
		} else if selected == 1 {
			handler.TransactionHandler()
		} else if selected == 2 {
			fmt.Println("Thank you. Goodbye!")
			break
		}
	}
}
