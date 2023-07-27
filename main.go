package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"taxcalc/entity"
	"taxcalc/utils"
)

func parseInput(inputStr string) (int, string, float64, error) {

	matches := regexp.MustCompile(utils.INPUT_PATTERN).FindStringSubmatch(inputStr)
	if len(matches) != 4 {
		fmt.Println(utils.INVALID_INPUT)
		return 0, "", 0.0, errors.New(utils.INVALID_INPUT)
	}

	quantity, err := strconv.Atoi(matches[1])
	if err != nil {
		fmt.Println(utils.INVALID_INPUT)
		return 0, "", 0.0, err
	}

	name := matches[2]
	price, err := strconv.ParseFloat(matches[3], 64)
	if err != nil {
		fmt.Println(utils.INVALID_INPUT)
		return 0, "", 0.0, err
	}

	return quantity, name, price, nil
}

func getInputAndProcess() []entity.Item {
	var items []entity.Item

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the shopping basket items in the format: <quantity> <name> at <price>")
	fmt.Println("Press 'Enter' on a new line to finish input.")
	fmt.Println()

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		if input == "" {
			break
		}

		quantity, name, price, err := parseInput(input)
		if err != nil {
			continue
		}

		item := entity.Item{
			Name:     name,
			Price:    price,
			Quantity: quantity,
		}

		if strings.Contains(input, utils.IMPORTED) {
			item.IsImported = true
		}

		if utils.IsExemptedItem(name) {
			item.IsSalesTaxExempt = true
		}

		items = append(items, item)
	}
	return items
}

// generate receipt
func generateReceipt(itms []entity.Item) entity.Receipt {
	var (
		items          []entity.Item
		totalImportTax float64
		totalSalesTax  float64
		totalPrice     float64
	)

	for _, item := range itms {

		importTax := utils.CalculateImportDutyTax(item) * float64(item.Quantity)
		salesTax := utils.CalculateSalesTax(item) * float64(item.Quantity)

		// add to relevant entities
		item.Price = item.Price + (importTax + salesTax)
		totalImportTax += importTax
		totalSalesTax += salesTax
		totalPrice += item.Price

		items = append(items, item)
	}

	return entity.Receipt{
		Items:      items,
		TotalTax:   totalSalesTax + totalImportTax,
		TotalPrice: totalPrice,
	}
}

func main() {
	items := getInputAndProcess()

	receipt := generateReceipt(items)

	fmt.Println("Receipt:")
	receipt.PrintReceipt()
}
