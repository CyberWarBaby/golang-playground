package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}

	sum := sumup(1, 10, 15, 9, 6, 12)

	anotherSum := sumup(numbers...)
	fmt.Println(sum)
	fmt.Println(anotherSum)
}

func sumup(numbers ...int) int {

	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
