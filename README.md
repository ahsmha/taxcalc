# Tax Calculator

This is a Golang application that generates a receipt based on the given items and their prices. The application calculates sales taxes and import duty taxes as per the provided requirements and prints the receipt details. This implementation follows OOP principles with structures and methods for tax calculation and receipt printing.

## Requirements

- Go (Golang) installed on your system.

## How to Build

1. Clone this repository to your local machine:

```git clone <repository-url>```

2. Change into the project directory:

```cd taxcalc```

3. Build the application:

```go build```

## How to Run

1. After building the application, you can run it with:

`./taxcalc`

2. Enter the shopping basket items in the format: `<quantity> <name> at <price>`, one item per line.

3. Press 'Enter' on a new line to finish input.

4. The application will calculate the receipt details based on the provided input and print the results.

## Example Input Format

```
1 book at 12.49
1 music CD at 14.99
1 chocolate bar at 0.85
```

## Assumptions

- The data is stored in-memory and not in a database.
- The rounding of taxes is done as per the given rules (rounded up to the nearest 0.05).
- The application assumes that the input data is well-formed and follows the specified format.







