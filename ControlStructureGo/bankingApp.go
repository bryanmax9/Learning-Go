package main

import (

	"fmt"
	"example.com/bankingapp/fileops"

)

const accountBalanceFile = "balance.txt"


func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("----------------")
		panic("Can't continue, sorry.")

	}

	fmt.Println("Welcome to Go Bank!")
	for {

		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is ", accountBalance)
		case 2:
			fmt.Print("Your Deposit: $")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid Amount. Must be greater than 0")
				continue
			}
			accountBalance += depositAmount

			fmt.Println("Balance Updated!")
			fmt.Println("Balance: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance,accountBalanceFile)
		case 3:
			var withdraw float64
			fmt.Println("How much you wanna Withdraw?")
			fmt.Print("Enter Value: $")
			fmt.Scan(&withdraw)

			if withdraw <= 0 || withdraw > accountBalance {
				fmt.Println("Invalid Amount")
				continue
			}

			accountBalance -= withdraw

			fmt.Println("you just withdrawn: $", withdraw)
			fmt.Println("New Balance: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our Bank")
			return
		}
	}
}


