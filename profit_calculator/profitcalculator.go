package main

import "fmt"

func main() {
	fmt.Println(`Welcome to Tax War Calculator(TWC)`)
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Input Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Input Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Input Tax rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	// profit := ebt - (ebt * (taxRate / 100))
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Println("Earnings before tax:", ebt)
	fmt.Println("Earnings after tax:", profit)
	fmt.Println("Ratio:", ratio)
}
