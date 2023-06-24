package ecommerce

type Product struct {
	Name  string
	Price float64
}

type ProductManager struct {
	products []Product
}

func NewProductManager() *ProductManager {
	return &ProductManager{}
}

func (pm *ProductManager) AddProduct(product Product) {
	pm.products = append(pm.products, product)
}

func (pm *ProductManager) GetProducts() []Product {
	return pm.products
}
