package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var expectedReturn float64
	var years float64

	fmt.Print("Investment Amount: ")
	utility("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return: ")
	fmt.Scan(&expectedReturn)

	futureValue, futureRealValue := calculteFutureValues(investmentAmount, expectedReturn, years)
	// futureValue := investmentAmount * math.Pow(1+expectedReturn/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// fmt.Println(futureRealValue)
	// fmt.Println(futureValue)

	formattedvalue := fmt.Sprintf("%.2f", futureValue)
	formattedvaluereal := fmt.Sprintf("%.2f", futureRealValue)
	fmt.Println("Future Value:", formattedvaluereal)
	fmt.Println("Future Real Value:", formattedvalue)
}

func utility(text string) {
	fmt.Print(text)
}

func calculteFutureValues(investmentAmount, expectedReturn, years float64) (fv float64, frv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturn/100, years)
	frv = fv / math.Pow(1+inflationRate/100, years)
	// return fv, frv
	return
}
