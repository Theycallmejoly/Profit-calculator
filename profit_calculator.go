package main

import (
	"fmt"
)

func main() {
	// Declare variables
	var salary, expenses, taxRate float64

	// Prompt user for input
	fmt.Print("Enter your salary: ")
	fmt.Scan(&salary)

	fmt.Print("Enter your expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter your tax rate: ")
	fmt.Scan(&taxRate)

	// Calculate Earnings Before Tax (EBT), Profit after Tax, and ratio
	ebt := salary
	profit := salary - (salary * (taxRate / 100))
	ratio := ebt / profit

	// Print the results
	fmt.Printf("Earnings Before Tax (EBT): %.2f\n", ebt)
	fmt.Printf("Profit after Tax: %.2f\n", profit)
	fmt.Printf("Ratio (EBT to Profit): %.2f\n", ratio)
}
