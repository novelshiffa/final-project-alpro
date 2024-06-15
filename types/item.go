package types

import (
	"fmt"
	"strings"
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

func (items *Items) IsColumn(columnName string) bool {
	lowerCasedColumnName := strings.ToLower(columnName)
	return lowerCasedColumnName == "id" || lowerCasedColumnName == "name" || lowerCasedColumnName == "category" || lowerCasedColumnName == "price" || lowerCasedColumnName == "stock"
}

func (items *Items) getMaxCharOnName() int {
	// Find Max Algorithm
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

func (p *Items) FilterBy(columnName string, value string) Items {
	columnName = strings.ToLower(columnName)

	var items Items
	columnFilters := map[string](func(item Item) bool){
		"id":       func(item Item) bool { return fmt.Sprintf("%d", item.Id) == value },
		"name":     func(item Item) bool { return item.Name == value },
		"category": func(item Item) bool { return item.Category == value },
		"price":    func(item Item) bool { return fmt.Sprintf("%d", item.Price) == value },
		"stock":    func(item Item) bool { return fmt.Sprintf("%d", item.Stock) == value },
	}

	filterFunc, exists := columnFilters[columnName]

	if !exists || !p.IsColumn(columnName) {
		panic("Undefined column.")
	}

	for i := 0; i < p.Length; i++ {
		if filterFunc(p.Items[i]) {
			items.Items[items.Length] = p.Items[i]

			items.Length++
		}
	}

	return items
}

func (p *Items) SortBy(columnName string, ascending bool) Items {
	// Selection Sort
	/*
		Id       int
		Name     string
		Category string
		Price    int
		Stock    int
	*/

	if !p.IsColumn(columnName) {
		panic("Undefined column name.")
	}

	columnName = strings.ToLower(columnName)
	var items Items

	for i := 0; i < p.Length; i++ {
		items.Items[i] = p.Items[i]
	}

	items.Length = p.Length

	for i := 0; i < items.Length-1; i++ {
		key := i

		for j := i + 1; j < p.Length; j++ {
			switch columnName {
			case "id":
				if items.Items[j].Id < items.Items[key].Id && ascending || items.Items[j].Id > items.Items[key].Id && !ascending {
					key = j
				}
			case "name":
				if items.Items[j].Name < items.Items[key].Name && ascending || items.Items[j].Name > items.Items[key].Name && !ascending {
					key = j
				}
			case "category":
				if items.Items[j].Category < items.Items[key].Category && ascending || items.Items[j].Category > items.Items[key].Category && !ascending {
					key = j
				}
			case "price":
				if items.Items[j].Price < items.Items[key].Price && ascending || items.Items[j].Price > items.Items[key].Price && !ascending {
					key = j
				}
			case "stock":
				if items.Items[j].Stock < items.Items[key].Stock && ascending || items.Items[j].Stock > items.Items[key].Stock && !ascending {
					key = j
				}
			}

		}

		items.Items[i], items.Items[key] = items.Items[key], items.Items[i]
	}

	return items
}
