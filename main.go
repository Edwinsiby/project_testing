package main

import (
	"fmt"

	"test/ecommerce"
	"test/repository"
)

func main() {
	repository.ConnectDB()

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

	err := ecommerce.NewUserManager().Signup("Edwin", "pass@123")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("user created")
	}

	err = ecommerce.NewUserManager().Login("Edwin", "pass@123")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("user logedin")
	}
}
