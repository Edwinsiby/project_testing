package ecommerce_test

import (
	"testing"

	"test/ecommerce"
)

func TestCart_AddItem(t *testing.T) {
	cart := ecommerce.NewCart()
	cart.AddItem("Item 1")
	cart.AddItem("Item 2")

	items := cart.GetItems()
	expected := []string{"Item 1", "Item 2"}

	if len(items) != len(expected) {
		t.Errorf("Expected %d items, but got %d", len(expected), len(items))
	}

	for i := range expected {
		if items[i] != expected[i] {
			t.Errorf("Expected item '%s', but got '%s'", expected[i], items[i])
		}
	}
}
