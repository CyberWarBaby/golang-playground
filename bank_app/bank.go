package main

import (
	"fmt"

	"example.com/bankapp/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accbalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("error!!!")
		fmt.Println(err)
		fmt.Println("--------------")
		// panic("Cant continue sorry!")
	}

	fmt.Println("Welcome to GO Bank!")
	for {
		presentFunction()

		var userInput int
		fmt.Print("Your choice: ")
		fmt.Scan(&userInput)

		// switch userInput{
		// case 1:
		// 	fmt.Println(`Your balance is`, accbalance)
		// case 2:
		// 	fmt.Print("Deposit amount: ")
		// 	var depositAmount float64
		// 	fmt.Scan(&depositAmount)

		// 	if depositAmount <= 0 {
		// 		fmt.Println("Must be greater than zero")
		// 		// return
		// 		continue
		// 	}

		// 	accbalance += depositAmount
		// 	fmt.Println("New balance is ", balance)

		// }

		if userInput == 1 {

			fmt.Println(`Your balance is`, accbalance)

		} else if userInput == 2 {

			fmt.Print("Deposit amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Must be greater than zero")
				// return
				continue
			}

			accbalance += depositAmount
			fmt.Println("New balance is ", accbalance)
			fileops.WriteFloatToFile(accbalance, accountBalanceFile)

		} else if userInput == 3 {

			fmt.Print("Withdrawal amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Withdraw amount must be greater than zero")
				// return
				continue
			}

			if withdrawAmount > accbalance {
				fmt.Println("Insufficient funds(You can't withdraw more than acc bal.)")
				// return
				continue
			}

			accbalance -= withdrawAmount
			fmt.Println("New balance is ", accbalance)
			fileops.WriteFloatToFile(accbalance, accountBalanceFile)

		} else {
			println("bye")
			// return is one way to do it but break is better
			break
		}
	}
}

// func presentFunction() {
// 	fmt.Println("What do yo want to do?")
// 	fmt.Println("1. check balance")
// 	fmt.Println("2. Deposit money")
// 	fmt.Println("3. Withdraw money")
// 	fmt.Println("4. Exit")
// }
