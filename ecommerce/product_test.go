package ecommerce_test

import (
	"testing"

	"test/ecommerce"
)

func TestProductManager_AddProduct(t *testing.T) {
	productManager := ecommerce.NewProductManager()

	productManager.AddProduct(ecommerce.Product{Name: "Product 1", Price: 10.99})
	productManager.AddProduct(ecommerce.Product{Name: "Product 2", Price: 19.99})

	products := productManager.GetProducts()
	expected := []ecommerce.Product{
		{Name: "Product 1", Price: 10.99},
		{Name: "Product 2", Price: 19.99},
	}

	if len(products) != len(expected) {
		t.Errorf("Expected %d products, but got %d", len(expected), len(products))
	}

	for i := range expected {
		if products[i].Name != expected[i].Name || products[i].Price != expected[i].Price {
			t.Errorf("Expected product '%s ($%.2f)', but got '%s ($%.2f)'",
				expected[i].Name, expected[i].Price, products[i].Name, products[i].Price)
		}
	}
}
