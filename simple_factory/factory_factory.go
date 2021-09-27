package simple_factory

import "fmt"

type Product interface {
	Show()
}

func NewProduct(name string) Product {
	switch name {
	case "milk":
		return &Milk{price: 4.5}
	case "book":
		return &Milk{price: 11.1}
	}
	return nil
}

type Milk struct {
	price float64
}

func (m *Milk) Show() {
	fmt.Printf("milk's price is %f\n", m.price)
}

type Book struct {
	price float64
}

func (m *Book) Show() {
	fmt.Printf("book's price is %f\n", m.price)
}
