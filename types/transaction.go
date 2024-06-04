package types

import (
	"fmt"
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

func (t *Transactions) Edit(id int, newTransaction Transaction) bool {
	var index int = t.FindById(id)

	if index != -1 {
		t.Items[index] = newTransaction

		return true
	}

	return false
}

func (t *Transactions) Delete(id int) bool {
	var index int = t.FindById(id)

	if index != -1 {
		for i := index; i < t.Length; i++ {
			t.Items[index] = t.Items[index+1]
		}

		return true
	}

	return false
}
