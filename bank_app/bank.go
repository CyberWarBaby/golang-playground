package main

import "fmt"

func main() {

	var accbalance float64 = 1000

	fmt.Println("Welcome to CYPH-GO Bank")
	fmt.Println("What will you like to do")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice? ")
	fmt.Scan(&choice)

	var amount float64

	if choice == 1 {
		fmt.Println("Your balance is:", accbalance)
	} else if choice == 2 {
		fmt.Print("Amount you want to deposit: ")
		fmt.Scan(&amount)
		accbalance += amount
		fmt.Println("Your new account balance is:", accbalance)
	} else if choice == 3 {
		fmt.Print("Amount you want to withdraw: ")
		fmt.Scan(&amount)
		accbalance -= amount
		fmt.Println("New account balance:", accbalance)
	} else {
		fmt.Println("Goodbye and stay safe!!")
	}

}
