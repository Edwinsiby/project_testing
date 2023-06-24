package ecommerce

type Cart struct {
	items []string
}

func NewCart() *Cart {
	return &Cart{}
}

func (c *Cart) AddItem(item string) {
	c.items = append(c.items, item)
}

func (c *Cart) GetItems() []string {
	return c.items
}
