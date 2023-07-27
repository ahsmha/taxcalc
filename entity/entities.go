package entity

import (
	"fmt"
)

type Item struct {
	Name             string
	Price            float64
	Quantity         int
	IsImported       bool
	IsSalesTaxExempt bool
}

type Receipt struct {
	Items      []Item
	TotalTax   float64
	TotalPrice float64
}

func (r Receipt) PrintReceipt() {
	for _, item := range r.Items {
		fmt.Printf("%d %s: %.2f\n", 1, item.Name, item.Price)
	}
	fmt.Printf("Sales Taxes: %.2f\n", r.TotalTax)
	fmt.Printf("Total: %.2f\n", r.TotalPrice)
}
