package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const NMAX int = 1000

// Global Variable Declarations -- Start

var products Products
var transactions Transactions

// Global Variable Declarations -- End

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
	// var i int
	var choice int

	for {
		clearTerminal()

		fmt.Println("Select Context")
		fmt.Println("[1] Product")
		fmt.Println("[2] Transaction")
		fmt.Println("[3] Exit")
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			ProductHandler()
		} else if choice == 2 {
			TransactionHandler()
		} else if choice == 3 {
			fmt.Println("Thank you. Good bye.")
			break
		}
	}
}
