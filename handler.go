package main

import "fmt"

func ProductHandler() {
	var choice int

	fmt.Println("[1] Add new product")
	fmt.Println("[2] View products")
	fmt.Println("[3] Edit product")
	fmt.Println("[4] Delete product")
	fmt.Println("Press any number except 1-4 to exit")
	fmt.Print("Your choice: ")

	fmt.Scan(&choice)

	if choice == 1 {
		// call add new product function
	} else if choice == 2 {
		// call view products function
	} else if choice == 3 {
		// call edit product function
	} else if choice == 4 {
		// call delete product function
	} else {
		return
	}
}

func TransactionHandler() {
	var choice int

	fmt.Println("[1] Create transaction")
	fmt.Println("[2] View transactions")
	fmt.Println("[3] Edit transaction")
	fmt.Println("[4] Delete transaction")
	fmt.Println("Press any number except 1-4 to exit")
	fmt.Print("Your choice: ")

	fmt.Scan(&choice)

	if choice == 1 {
		// call create transaction function
	} else if choice == 2 {
		// call view transactions function
	} else if choice == 3 {
		// call edit transaction function
	} else if choice == 4 {
		// call delete transaction function
	} else {
		return
	}
}
