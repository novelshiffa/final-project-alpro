package types

import (
	"fmt"
	"strings"
	"time"
)

type Transaction struct {
	Id       int
	Time     time.Time
	Type     string
	Item     Item
	Quantity int
}

type Transactions struct {
	Items  [NMAX]Transaction
	Length int
}

func (Transactions *Transactions) IsColumn(columnName string) bool {
	lowerCasedColumnName := strings.ToLower(columnName)
	return lowerCasedColumnName == "id" || lowerCasedColumnName == "time" || lowerCasedColumnName == "type" || lowerCasedColumnName == "itemid" || lowerCasedColumnName == "quantity"
}

func (transactions *Transactions) ShowInTable() {
	fmt.Printf("%-5s %-20s %-10s %-20s %-10s %-10s\n", "ID", "Time", "Type", "Item Id", "Quantity", "Price")
	fmt.Println("-------------------------------------------------------------------------------------------")
	for i := 0; i < transactions.Length; i++ {
		t := transactions.Items[i]
		fmt.Printf("%-5d %-20s %-10s %-20d %-10d %-10.2d\n", t.Id, t.Time.Format("2006-01-02 15:04:05"), t.Type, t.Item.Id, t.Quantity, t.Item.Price)
	}
}

func (t *Transactions) FindById(id int) int {
	// Binary search algorithm
	// TODO: Check if sorted

	low, high := 0, t.Length-1
	for low <= high {
		mid := (low + high) / 2
		if t.Items[mid].Id == id {
			return mid
		} else if t.Items[mid].Id < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func (t *Transactions) CreateNew(transaction Transaction) {
	if t.Length == NMAX {
		panic("Max length reached.")
	}

	t.Items[t.Length] = transaction
	t.Items[t.Length].Id = t.Length + 1
	t.Length++
}

func (t *Transactions) FetchAll() {
	for i := 0; i < t.Length; i++ {
		fmt.Println(t.Items[i])
	}
}

func (t *Transactions) Delete(idx int) {
	if idx < 0 || idx >= t.Length {
		panic("Index out of range.")
	}

	for i := idx; i < t.Length; i++ {
		t.Items[i] = t.Items[i+1]
	}

	t.Length--
}

func (t *Transactions) FilterBy(columnName string, value string) Transactions {
	columnName = strings.ToLower(columnName)

	var transactions Transactions
	columnFilters := map[string](func(transaction Transaction) bool){
		"id": func(transaction Transaction) bool { return fmt.Sprintf("%d", transaction.Id) == value },
		"time": func(transaction Transaction) bool {
			return transaction.Time.Format("2006-01-02 15:04:05") == value || strings.Split(transaction.Time.Format("2006-01-02 15:04:05"), " ")[0] == value
		},
		"type":     func(transaction Transaction) bool { return transaction.Type == value },
		"itemid":   func(transaction Transaction) bool { return fmt.Sprintf("%d", transaction.Item.Id) == value },
		"quantity": func(transaction Transaction) bool { return fmt.Sprintf("%d", transaction.Quantity) == value },
	}

	filterFunc, exists := columnFilters[columnName]

	if !exists || !t.IsColumn(columnName) {
		panic("Undefined column.")
	}

	for i := 0; i < t.Length; i++ {
		if filterFunc(t.Items[i]) {
			transactions.Items[transactions.Length] = t.Items[i]
			transactions.Length++
		}
	}

	return transactions
}

func (t *Transactions) SortBy(columnName string, ascending bool) Transactions {
	// Insertion Sort
	// type Transaction struct {
	// 	Id       int
	// 	Time     time.Time
	// 	Type     string
	// 	Item     Item
	// 	Quantity int
	// }

	if !t.IsColumn(columnName) {
		panic("Undefined column name.")
	}

	columnName = strings.ToLower(columnName)
	var transactions Transactions

	for i := 0; i < t.Length; i++ {
		transactions.Items[i] = t.Items[i]
	}

	transactions.Length = t.Length

	for i := 1; i < t.Length; i++ {
		j := i

		for j > 0 {
			switch columnName {
			case "id":
				if transactions.Items[j].Id < transactions.Items[j-1].Id && ascending || transactions.Items[j].Id > transactions.Items[j-1].Id && !ascending {
					transactions.Items[j-1], transactions.Items[j] = transactions.Items[j], transactions.Items[j-1]
				}
			case "time":
				if (transactions.Items[j].Time.Before(transactions.Items[j-1].Time) && ascending) ||
					(transactions.Items[j].Time.After(transactions.Items[j-1].Time) && !ascending) {
					transactions.Items[j-1], transactions.Items[j] = transactions.Items[j], transactions.Items[j-1]
				}
			case "itemid":
				if transactions.Items[j].Item.Id < transactions.Items[j-1].Item.Id && ascending || transactions.Items[j].Item.Id > transactions.Items[j-1].Item.Id && !ascending {
					transactions.Items[j-1], transactions.Items[j] = transactions.Items[j], transactions.Items[j-1]
				}
			case "type":
				if transactions.Items[j].Type < transactions.Items[j-1].Type && ascending || transactions.Items[j].Type > transactions.Items[j-1].Type && !ascending {
					transactions.Items[j-1], transactions.Items[j] = transactions.Items[j], transactions.Items[j-1]
				}
			case "quantity":
				if transactions.Items[j].Quantity < transactions.Items[j-1].Quantity && ascending || transactions.Items[j].Quantity > transactions.Items[j-1].Quantity && !ascending {
					transactions.Items[j-1], transactions.Items[j] = transactions.Items[j], transactions.Items[j-1]
				}
			}

			j--

		}
	}

	return transactions
}
