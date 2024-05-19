package main

import "time"

type Product struct {
	Id       int
	Name     string
	Category string
	Price    int
	Stock    int
}

type Transaction struct {
	Id       int
	Time     time.Time
	Type     string
	Product  Product
	Quantity int
}

type Products struct {
	items  [NMAX]Product
	length int
}

type Transactions struct {
	items  [NMAX]Transaction
	length int
}
