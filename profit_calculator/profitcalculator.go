package main

import (
	"errors"
	"fmt"
	"os"
)

// func writeCalcResultToFile(calcResult float64, float64, float64) {
// 	calulationText := fmt.Sprint(calcResult)
// 	os.WriteFile("calculationFile.txt", []byte(calulationText), 0644)
// }

func main() {
	fmt.Println(`Welcome to Tax War Calculator(TWC)`)
	// var revenue float64
	// var expenses float64
	// var taxRate float64
	// var err string

	// fmt.Print("Input Revenue: ")
	// fmt.Scan(&revenue)
	revenue, err1 := userInput("Input Revenue: ")
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	// fmt.Print("Input Expenses: ")
	// fmt.Scan(&expenses)
	expenses, err2 := userInput("Input Expenses: ")
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	// fmt.Print("Input Tax rate: ")
	// fmt.Scan(&taxRate)
	taxRate, err3 := userInput("Input Tax rate: ")
	if err3 != nil {
		fmt.Println(err3)
		return
	}

	ebt, profit, ratio := theCalculations(revenue, expenses, taxRate)
	// ebt := revenue - expenses
	// // profit := ebt - (ebt * (taxRate / 100))
	// profit := ebt * (1 - taxRate/100)
	// ratio := ebt / profit

	fmt.Println("Earnings before tax:", ebt)
	fmt.Println("Earnings after tax:", profit)
	fmt.Println("Ratio:", ratio)
	storeResults(ebt, profit, ratio)
}

func userInput(text string) (float64, error) {
	var userInputInfo float64
	fmt.Print(text)
	fmt.Scan(&userInputInfo)

	if userInputInfo <= 0 {
		return 0, errors.New("Value_Must_Be_Positive")
	}

	return userInputInfo, nil
}

func theCalculations(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit

	return ebt, profit, ratio
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("result.txt", []byte(results), 0644)
}
