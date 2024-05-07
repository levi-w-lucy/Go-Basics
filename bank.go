package main

import (
	"fmt"

	"example.com/bank/fileHandler"
	"github.com/Pallinder/go-randomdata"
)

const balanceFile = "balance.txt"

func main() {
	accountBalance, err := fileHandler.GetFloatFromFile(balanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println("----------")
		panic(err)
	}
	var choice int

	fmt.Println("Welcome to Go Bank!")
	fmt.Println(randomdata.PhoneNumber())

interfaceLoop:
	for {
		presentOptions()
		fmt.Print("Enter a choice: ")
		fmt.Scan(&choice)

		fmt.Println("Your choice: ", choice)
		switch choice {
		case 1:
			fmt.Println("Your balance is ", formatMoney(accountBalance))
		case 2:
			accountBalance = depositFunds(accountBalance)
			fileHandler.WriteFloatToFile(accountBalance, balanceFile)
		case 3:
			accountBalance = withdrawlFunds(accountBalance)
			fileHandler.WriteFloatToFile(accountBalance, balanceFile)
		default:
			fmt.Println("Goodbye!")
			break interfaceLoop
		}
	}

	fmt.Println("Thanks for choosing our bank!")
}
