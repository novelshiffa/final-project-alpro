package main

import "fmt"

func main() {
	// var i int
	var choice int

	for {
		fmt.Println("1. Product")
		fmt.Println("2. Transaction")
		fmt.Println("3. Exit")
		fmt.Scanln(&choice)

		if choice == 1 {
			fmt.Println("Choice Product")
		} else if choice == 2 {
			fmt.Println("Choice Transaction")
		} else if choice == 3 {
			break
		}
	}
}
