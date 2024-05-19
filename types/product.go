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
	items  [NMAX]Product
	length int
}

func (p *Products) FindById(id int) int {
	// Sequential search algorithm
	for i := 0; i < p.length; i++ {
		if p.items[i].Id == id {
			return i
		}
	}

	return -1
}

func (p *Products) AddNew(product Product) {
	if p.length == NMAX {
		panic("Max length reached.")
	}

	p.items[p.length] = product
	p.length++
}

func (p *Products) FetchAll() {
	for i := 0; i < p.length; i++ {
		fmt.Println(p.items[i])
	}
}

func (p *Products) Edit(id int, newProduct Product) bool {
	var index int = p.FindById(id)

	if index != -1 {
		p.items[index] = newProduct

		return true
	}

	return false
}

func (p *Products) Delete(id int) bool {
	var index int = p.FindById(id)

	if index != -1 {
		for i := index; i < p.length; i++ {
			p.items[index] = p.items[index+1]
		}

		return true
	}

	return false
}
