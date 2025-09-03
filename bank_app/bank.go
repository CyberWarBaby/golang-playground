package main

import (
	"fmt"
	"os"
	"strconv"
)

const accbalanceFile = "balance.txt"

func getBalancefromFile() float64 {
	data, _ := os.ReadFile(accbalanceFile)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}
func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accbalanceFile, []byte(balanceText), 0644)
}

func main() {

	var accbalance float64 = getBalancefromFile()

	fmt.Println("Welcome to CYPH-GO Bank")

	for {
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

			if amount <= 0 {
				fmt.Println("Please enter a valid amount!")
				continue
			}

			accbalance += amount
			fmt.Println("Your new account balance is:", accbalance)
			writeBalanceToFile(accbalance)
		} else if choice == 3 {
			fmt.Print("Amount you want to withdraw: ")

			if amount > accbalance {
				fmt.Print("SORRY! Insufficient balance :( ")
				return
			}

			fmt.Scan(&amount)
			accbalance -= amount
			fmt.Println("New account balance:", accbalance)
			writeBalanceToFile(accbalance)
		} else {
			fmt.Println("Goodbye and stay safe!!")
			fmt.Println("Thanks for choosing us")
			break
		}

	}
}
