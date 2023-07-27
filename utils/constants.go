package utils

const (
	// constant
	SALES_TAX  = 10
	IMPORT_TAX = 5
	IMPORTED   = "imported"

	// patterns
	INPUT_PATTERN = `(\d+) ([\w\s]+) at (\d+\.\d{2})`

	// errors
	INVALID_INPUT = "Invalid input format. Please use the format: <quantity> <name> at <price>"
)

var ExemptedItems = []string{"book", "chocolate bar", "chocolates", "headache pills", "pills", "medicine", "food"}
