package main

import "fmt"

func main() {
	fmt.Println(`Welcome to Tax War Calculator(TWC)`)
	var revenue float64
	var expenses float64
	var taxRate float64

	// fmt.Print("Input Revenue: ")
	// fmt.Scan(&revenue)
	revenue = inputUser("Input Revenue: ")

	// fmt.Print("Input Expenses: ")
	// fmt.Scan(&expenses)
	expenses = inputUser("Input Expenses: ")

	// fmt.Print("Input Tax rate: ")
	// fmt.Scan(&taxRate)
	taxRate = inputUser("Input Tax rate: ")

	ebt, profit, ratio := theCalculations(revenue, expenses, taxRate)
	// ebt := revenue - expenses
	// // profit := ebt - (ebt * (taxRate / 100))
	// profit := ebt * (1 - taxRate/100)
	// ratio := ebt / profit

	fmt.Println("Earnings before tax:", ebt)
	fmt.Println("Earnings after tax:", profit)
	fmt.Println("Ratio:", ratio)
}

func inputUser(text string) float64 {
	var userIput float64
	fmt.Print(text)
	fmt.Scan(&userIput)

	return userIput
}

func theCalculations(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit

	return ebt, profit, ratio
}
