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
	var subChoice int

	for {
		fmt.Println("Select Context")
		fmt.Println("[1] Product")
		fmt.Println("[2] Transaction")
		fmt.Println("[3] Exit")
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("[1] Add new product")
			fmt.Println("[2] View products")
			fmt.Println("[3] Edit product")
			fmt.Println("[4] Delete product")
			fmt.Println("Press any number except 1-4 to exit")
			fmt.Print("Your choice: ")

			fmt.Scan(&subChoice)

			if subChoice == 1 {
				// call add new product function
			} else if subChoice == 2 {
				// call view products function
			} else if subChoice == 3 {
				// call edit product function
			} else if subChoice == 4 {
				// call delete product function
			} else {
				break
			}

		} else if choice == 2 {
			fmt.Println("[1] Create transaction")
			fmt.Println("[2] View transactions")
			fmt.Println("[3] Edit transaction")
			fmt.Println("[4] Delete transaction")
			fmt.Println("Press any number except 1-4 to exit")
			fmt.Print("Your choice: ")

			fmt.Scan(&subChoice)

			if subChoice == 1 {
				// call create transaction function
			} else if subChoice == 2 {
				// call view transactions function
			} else if subChoice == 3 {
				// call edit transaction function
			} else if subChoice == 4 {
				// call delete transaction function
			} else {
				break
			}
		} else if choice == 3 {
			fmt.Println("Thank you. Good bye.")
			break
		}
	}
}
