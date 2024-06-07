package types

import (
	"fmt"
)

type Item struct {
	Id       int
	Name     string
	Category string
	Price    int
	Stock    int
}

type Items struct {
	Items  [NMAX]Item
	Length int
}

func (items *Items) getMaxCharOnName() int {
	max := 20
	for i := 0; i < items.Length; i++ {
		if len(items.Items[i].Name) > max {
			max = len(items.Items[i].Name)
		}
	}

	return max
}

func (p *Items) ShowInTable() {
	// Get the maximum name length
	nameWidth := p.getMaxCharOnName()

	// Construct the format string dynamically
	headerFormat := fmt.Sprintf("%%-5s %%-%ds %%-15s %%-10s %%-10s\n", nameWidth)
	rowFormat := fmt.Sprintf("%%-5d %%-%ds %%-15s %%-10d %%-10d\n", nameWidth)

	// Table header
	fmt.Printf(headerFormat, "ID", "Name", "Category", "Price", "Stock")
	fmt.Println("--------------------------------------------------------------")

	// Print each item
	for i := 0; i < p.Length; i++ {
		item := p.Items[i]
		fmt.Printf(rowFormat, item.Id, item.Name, item.Category, item.Price, item.Stock)
	}
}

func (p *Items) FindById(id int) int {
	// Sequential search algorithm
	for i := 0; i < p.Length; i++ {
		if p.Items[i].Id == id {
			return i
		}
	}

	return -1
}

func (p *Items) AddNew(item Item) {
	if p.Length == NMAX {
		panic("Something went wrong")
	}

	p.Items[p.Length] = item
	p.Items[p.Length].Id = p.Length + 1
	p.Length++
}

func (p *Items) FetchAll() {
	for i := 0; i < p.Length; i++ {
		fmt.Println(p.Items[i])
	}
}

func (p *Items) Delete(idx int) {
	if idx < 0 || idx >= p.Length {
		panic("Index out of range.")
	}

	for i := idx; i < p.Length; i++ {
		p.Items[i] = p.Items[i+1]
	}

	p.Length--
}
