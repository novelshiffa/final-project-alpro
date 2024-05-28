package types

import (
	"errors"
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

func (p *Items) ShowInTable() {
	// Table header
	fmt.Printf("%-5s %-20s %-15s %-10s %-10s\n", "ID", "Name", "Category", "Price", "Stock")
	fmt.Println("--------------------------------------------------------------")

	// Print each item
	for i := 0; i < p.Length; i++ {
		item := p.Items[i]
		fmt.Printf("%-5d %-20s %-15s %-10d %-10d\n", item.Id, item.Name, item.Category, item.Price, item.Stock)
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

func (p *Items) AddNew(item Item) (bool, error) {
	if p.Length == NMAX {
		return false, errors.New("max length reached")
	}

	p.Items[p.Length] = item
	p.Items[p.Length].Id = p.Length + 1
	p.Length++

	return true, nil
}

func (p *Items) FetchAll() {
	for i := 0; i < p.Length; i++ {
		fmt.Println(p.Items[i])
	}
}

func (p *Items) Edit(id int, newItem Item) (bool, error) {
	var index int = p.FindById(id)

	if index != -1 {
		p.Items[index] = newItem

		return true, nil
	}

	return false, errors.New("something went wrong")
}

func (p *Items) Delete(id int) (bool, error) {
	var index int = p.FindById(id)

	if index != -1 {
		for i := index; i < p.Length; i++ {
			p.Items[index] = p.Items[index+1]
		}

		return true, nil
	}

	return false, errors.New("something went wrong")
}
