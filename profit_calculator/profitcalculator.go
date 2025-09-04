package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Println("WELCOME TO MY PROFIT CALCULATION APP")

	// revenue, expenses, taxRate = getuserInput()

	// fmt.Print("Kindly enter your revenue: ")
	// fmt.Scan(&revenue)

	// fmt.Print("Kindly enter your expenses: ")
	// fmt.Scan(&expenses)

	// fmt.Print("Kindly enter your Tax rate: ")
	// fmt.Scan(&taxRate)

	revenue = newGetInfo("Revenue: ")
	expenses = newGetInfo("Expenses: ")
	taxRate = newGetInfo("Tax rate: ")

	// earningbeforetax := revenue - expenses
	// profit := earningbeforetax * (1 - (taxRate / 100))
	// ratio := earningbeforetax / profit
	earningbeforetax, profit, ratio := calculation(revenue, expenses, taxRate)

	fmt.Println("Earnings before tax:", earningbeforetax)
	formattedprofit := fmt.Sprintf("%0.2f", profit)
	fmt.Println("Profit:", formattedprofit)
	formattedRatio := fmt.Sprintf("%0.2f", ratio)
	fmt.Println("Ratio:", formattedRatio)
}

// func getuserInput() (revenue float64, expenses float64, taxRate float64) {
// 	fmt.Print("Kindly enter your revenue: ")
// 	fmt.Scan(&revenue)

// 	fmt.Print("Kindly enter your expenses: ")
// 	fmt.Scan(&expenses)

// 	fmt.Print("Kindly enter your Tax rate: ")
// 	fmt.Scan(&taxRate)

// 	return
// }

func newGetInfo(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		fmt.Print("Invalid correct")
	}
	return userInput
}

func calculation(revenue, expenses, taxRate float64) (earningbeforetax float64, profit float64, ratio float64) {
	earningbeforetax = revenue - expenses
	profit = earningbeforetax * (1 - (taxRate / 100))
	ratio = earningbeforetax / profit

	return
}
