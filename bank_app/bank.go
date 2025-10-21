package main

import "fmt"

func main() {
	var balance float64 = 1000

	fmt.Println("Welcome to GO Bank!")
	fmt.Println("What do yo want to do?")
	fmt.Println("1. check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var userInput int
	fmt.Print("Your choice: ")
	fmt.Scan(&userInput)

	if userInput == 1 {

		fmt.Println(`Your balance is`, balance)

	} else if userInput == 2 {

		fmt.Print("Deposit amount: ")
		var depositAmount float64

		if depositAmount <= 0 {
			fmt.Println("Must be greater than zero")
			return
		}

		fmt.Scan(&depositAmount)
		balance += depositAmount
		fmt.Println("New balance is ", balance)

	} else if userInput == 3 {

		fmt.Print("Withdrawal amount: ")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)
		balance -= withdrawAmount
		fmt.Println("New balance is ", balance)

	} else {
		println("bye")
	}

}
