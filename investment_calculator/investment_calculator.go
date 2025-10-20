package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	fmt.Println("Hello investment")
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.6

	fmt.Print("Input the investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Input number of years: ")
	fmt.Scan(&years)

	fmt.Print("Input Expected Return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future real value: %.2f\n", futureRealValue)
	//fmt.Println("Future Value:", futureValue)
	// fmt.Println("Future Real Value:", futureRealValue)
	// fmt.Printf("Future value: %.2f\n", futureValue)
	// fmt.Printf("Future real value: %.2f\n", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)
}
