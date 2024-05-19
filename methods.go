package main

import "fmt"

// Methods for products
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

// Methods for transactions
func (t *Transactions) FindById(id int) int {
	low, high := 0, t.length-1
	for low <= high {
		mid := (low + high) / 2
		if t.items[mid].Id == id {
			return mid
		} else if t.items[mid].Id < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func (t *Transactions) AddNew(transaction Transaction) {
	if t.length == NMAX {
		panic("Max length reached.")
	}

	t.items[t.length] = transaction
	t.length++
}

func (t *Transactions) FetchAll() {
	for i := 0; i < t.length; i++ {
		fmt.Println(t.items[i])
	}
}

func (t *Transactions) Edit(id int, newTransaction Transaction) bool {
	var index int = t.FindById(id)

	if index != -1 {
		t.items[index] = newTransaction

		return true
	}

	return false
}

func (t *Transactions) Delete(id int) bool {
	var index int = t.FindById(id)

	if index != -1 {
		for i := index; i < t.length; i++ {
			t.items[index] = t.items[index+1]
		}

		return true
	}

	return false
}
