package main

import (
	"fmt"

	"test/ecommerce"
)

func main() {
	cart := ecommerce.NewCart()
	productManager := ecommerce.NewProductManager()

	productManager.AddProduct(ecommerce.Product{Name: "Product 1", Price: 10.99})
	productManager.AddProduct(ecommerce.Product{Name: "Product 2", Price: 19.99})

	cart.AddItem("Product 1")
	cart.AddItem("Product 2")

	items := cart.GetItems()
	fmt.Println("Cart Items:")
	for _, item := range items {
		fmt.Println("- " + item)
	}

	products := productManager.GetProducts()
	fmt.Println("\nAvailable Products:")
	for _, product := range products {
		fmt.Printf("- %s ($%.2f)\n", product.Name, product.Price)
	}
}
