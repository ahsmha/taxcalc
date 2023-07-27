package utils

import (
	"math"
	"strings"
	"taxcalc/entity"
)

// Round a float number to the nearest 0.05
func RoundToNearest0_05(number float64) float64 {
	return math.Ceil(number*20) / 20
}

// if the item is exempted from sales tax
func IsExemptedItem(itemName string) bool {
	for _, exemptedItem := range ExemptedItems {
		if strings.Contains(itemName, exemptedItem) {
			return true
		}
	}
	return false
}

// Calculate the sales tax for an item
func CalculateSalesTax(item entity.Item) float64 {
	if item.IsSalesTaxExempt {
		return 0.0
	}

	salesTax := (item.Price * SALES_TAX) / 100
	return RoundToNearest0_05(salesTax)
}

// Calculate the import duty tax for an item
func CalculateImportDutyTax(item entity.Item) float64 {
	if item.IsImported {
		importDuty := (item.Price * IMPORT_TAX) / 100
		return RoundToNearest0_05(importDuty)
	}
	return 0.0
}
