package types

import "fmt"

type Product struct {
	Id       int
	Name     string
	Category string
	Price    int
	Stock    int
}

type Products struct {
	Items  [NMAX]Product
	Length int
}

func (p *Products) ShowInTable() {
	// Table header
	fmt.Printf("%-5s %-20s %-15s %-10s %-10s\n", "ID", "Name", "Category", "Price", "Stock")
	fmt.Println("--------------------------------------------------------------")

	// Print each product
	for i := 0; i < p.Length; i++ {
		product := p.Items[i]
		fmt.Printf("%-5d %-20s %-15s %-10d %-10d\n", product.Id, product.Name, product.Category, product.Price, product.Stock)
	}
}

func (p *Products) FindById(id int) int {
	// Sequential search algorithm
	for i := 0; i < p.Length; i++ {
		if p.Items[i].Id == id {
			return i
		}
	}

	return -1
}

func (p *Products) AddNew(product Product) {
	if p.Length == NMAX {
		panic("Max length reached.")
	}

	p.Items[p.Length] = product
	p.Length++
}

func (p *Products) FetchAll() {
	for i := 0; i < p.Length; i++ {
		fmt.Println(p.Items[i])
	}
}

func (p *Products) Edit(id int, newProduct Product) bool {
	var index int = p.FindById(id)

	if index != -1 {
		p.Items[index] = newProduct

		return true
	}

	return false
}

func (p *Products) Delete(id int) bool {
	var index int = p.FindById(id)

	if index != -1 {
		for i := index; i < p.Length; i++ {
			p.Items[index] = p.Items[index+1]
		}

		return true
	}

	return false
}
