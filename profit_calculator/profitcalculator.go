package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	fmt.Println("WELCOME TO MY PROFIT CALCULATION APP")

	revenue, err1 := newGetInfo("Revenue: ")
	expenses, err2 := newGetInfo("Expenses: ")
	taxRate, err3 := newGetInfo("Tax rate: ")

	if  err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1) //i can use any of the errors since they will all be the same
		return
	}

	// earningbeforetax := revenue - expenses
	// profit := earningbeforetax * (1 - (taxRate / 100))
	// ratio := earningbeforetax / profit
	earningbeforetax, profit, ratio := calculation(revenue, expenses, taxRate)

	fmt.Println("Earnings before tax:", earningbeforetax)
	formattedprofit := fmt.Sprintf("%0.2f", profit)
	fmt.Println("Profit:", formattedprofit)
	formattedRatio := fmt.Sprintf("%0.2f", ratio)
	fmt.Println("Ratio:", formattedRatio)

	storeDatabase(earningbeforetax, profit, ratio)
}

func storeDatabase (earningbeforetax, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", earningbeforetax, profit, ratio)
	os.WriteFile("result.txt", []byte(results), 0644)
}

func calculation(revenue, expenses, taxRate float64) (earningbeforetax float64, profit float64, ratio float64) {
	earningbeforetax = revenue - expenses
	profit = earningbeforetax * (1 - (taxRate / 100))
	ratio = earningbeforetax / profit

	return
}

func newGetInfo(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("value must be positive")
	}
	return userInput, nil
}
