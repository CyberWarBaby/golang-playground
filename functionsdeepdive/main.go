package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	moreNumber := []int{5, 1, 2}
	doubled := transformNumbers(&numbers, triple)
	// for _, v := range number {
	// 	multv := v * 2
	// 	fmt.Println(multv)
	// }
	fmt.Println(doubled)

	transformFn1 := getTransformerFunction(&numbers)
	transformFn2 := getTransformerFunction(&moreNumber)

	transformedNumbers := transformNumbers(&numbers, transformFn1)
	moretransformedNumbers := transformNumbers(&moreNumber, transformFn2)

	fmt.Println(transformedNumbers, moretransformedNumbers)

}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dnumber := []int{}
	for _, value := range *numbers {
		dnumber = append(dnumber, transform(value))
	}
	return dnumber
}

func getTransformerFunction(numbers *[]int) transformFn {

	if (*numbers)[0] == 1 {

		return double

	} else {
		return triple
	}
}
func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
