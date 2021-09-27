package simple_factory

import "testing"

func TestNewProduct(t *testing.T) {
	product := NewProduct("milk")
	product.Show()
	product = NewProduct("book")
	product.Show()
}
