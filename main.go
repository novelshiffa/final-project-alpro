package main

import (
	"fmt"
	"time"

	"github.com/novelshiffa/final-project-alpro/handler"
	"github.com/novelshiffa/final-project-alpro/types"
)

// Global Variable Declarations -- Start

var items types.Items
var transactions types.Transactions

// Global Variable Declarations -- End

func main() {
	items = types.Items{
		Items: [types.NMAX]types.Item{
			{Id: 1, Name: "Beras Rojo Lele", Category: "Makanan Pokok", Price: 12000, Stock: 100},
			{Id: 2, Name: "Minyak Goreng Bimoli", Category: "Minyak", Price: 14000, Stock: 80},
			{Id: 3, Name: "Gula Pasir Gulaku", Category: "Bahan Pokok", Price: 13500, Stock: 70},
			{Id: 4, Name: "Tepung Terigu Segitiga Biru", Category: "Bahan Pokok", Price: 9000, Stock: 90},
			{Id: 5, Name: "Telur Ayam Kampung", Category: "Protein", Price: 20000, Stock: 60},
			{Id: 6, Name: "Garam Refina", Category: "Bumbu Dapur", Price: 5000, Stock: 50},
			{Id: 7, Name: "Kecap Manis ABC", Category: "Bumbu Dapur", Price: 11000, Stock: 40},
			{Id: 8, Name: "Susu Kental Manis Indomilk", Category: "Minuman", Price: 8500, Stock: 30},
			{Id: 9, Name: "Mie Instan Indomie", Category: "Makanan Instan", Price: 2500, Stock: 200},
			{Id: 10, Name: "Bawang Merah Brebes", Category: "Bumbu Dapur", Price: 30000, Stock: 35},
		},
		Length: 10,
	}

	transactions = types.Transactions{
		Items: [types.NMAX]types.Transaction{
			{Id: 1, Time: time.Date(2024, 6, 1, 14, 0, 0, 0, time.UTC), Type: "incoming", Item: items.Items[0], Quantity: 10},
			{Id: 2, Time: time.Date(2024, 6, 2, 15, 0, 0, 0, time.UTC), Type: "outgoing", Item: items.Items[1], Quantity: 5},
			{Id: 3, Time: time.Date(2024, 6, 3, 16, 0, 0, 0, time.UTC), Type: "incoming", Item: items.Items[1], Quantity: 7},
			{Id: 4, Time: time.Date(2024, 6, 4, 17, 0, 0, 0, time.UTC), Type: "outgoing", Item: items.Items[2], Quantity: 3},
			{Id: 5, Time: time.Date(2024, 6, 5, 18, 0, 0, 0, time.UTC), Type: "incoming", Item: items.Items[3], Quantity: 12},
			{Id: 6, Time: time.Date(2024, 6, 6, 19, 0, 0, 0, time.UTC), Type: "outgoing", Item: items.Items[3], Quantity: 2},
			{Id: 7, Time: time.Date(2024, 6, 7, 20, 0, 0, 0, time.UTC), Type: "incoming", Item: items.Items[3], Quantity: 8},
			{Id: 8, Time: time.Date(2024, 6, 8, 21, 0, 0, 0, time.UTC), Type: "outgoing", Item: items.Items[4], Quantity: 1},
			{Id: 9, Time: time.Date(2024, 6, 9, 22, 0, 0, 0, time.UTC), Type: "incoming", Item: items.Items[4], Quantity: 20},
			{Id: 10, Time: time.Date(2024, 6, 10, 23, 0, 0, 0, time.UTC), Type: "outgoing", Item: items.Items[4], Quantity: 4},
		},
		Length: 10,
	}

	var menu types.Menu

	menu.DefaultSelectedColor = "blue"
	menu.Items[0] = types.NewText("[1] Item")
	menu.Items[1] = types.NewText("[2] Transactions")
	menu.Items[2] = types.NewText("[3] Exit program")

	menu.Length = 3
	menu.SetSelected(0)

	var selected int = 0
	var stopLoop bool = false
	var cls bool = true

	menu.Listen(&selected, &stopLoop, &cls, func() {
		switch selected {
		case 0:
			stopLoop = !handler.ItemHandler(&items)
		case 1:
			stopLoop = !handler.TransactionHandler(&transactions, &items)
		case 2:
			stopLoop = true
			var goodByeText = types.NewText("さよなら！")
			goodByeText.SetColor("green")
			fmt.Println(goodByeText.Colored)
		}
	}, func() {})

}
