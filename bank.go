package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)

	os.WriteFile(balanceFile, []byte(balanceText), 644)
}

func getBalanceFromFile() (float64, error) {
	byteData, err := os.ReadFile(balanceFile)
	amountAsString := string(byteData)

	if err != nil {
		return 1000, errors.New("Failed to find balance file.")
	}

	balance, err := strconv.ParseFloat(amountAsString, 64)
	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value.")
	}
	return balance, nil
}

func main() {
	accountBalance, err := getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------")
	}
	var choice int

	fmt.Println("Welcome to Go Bank!")

interfaceLoop:
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")
		fmt.Print("Enter a choice: ")
		fmt.Scan(&choice)

		fmt.Println("Your choice: ", choice)
		switch choice {
		case 1:
			fmt.Println("Your balance is ", formatMoney(accountBalance))
		case 2:
			accountBalance = depositFunds(accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			accountBalance = withdrawlFunds(accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye!")
			break interfaceLoop
		}
	}

	fmt.Println("Thanks for choosing our bank!")
}

func depositFunds(originalBalance float64) float64 {
	var depositAmount float64

	if depositAmount < 0 {
		fmt.Println("Invalid amount. Amount should be greater than 0")
		return originalBalance
	}

	fmt.Print("Enter an amount to deposit: ")
	fmt.Scan(&depositAmount)

	fmt.Println(formatMoney(depositAmount), " deposited!")
	fmt.Println("New balance is: ", formatMoney(originalBalance+depositAmount))

	return originalBalance + depositAmount
}

func withdrawlFunds(originalBalance float64) float64 {
	var withdrawlAmount float64
	fmt.Print("Enter an amount to withdraw: ")
	fmt.Scan(&withdrawlAmount)

	if withdrawlAmount < 0 {
		fmt.Println("Invalid amount. Amount should be greater than 0")
		return originalBalance
	}

	if withdrawlAmount > originalBalance {
		fmt.Println("Cannot withdraw more than the existing balance. Current balance is ", formatMoney(originalBalance))
		return originalBalance
	}

	fmt.Println(formatMoney(withdrawlAmount), " withdrawn.")
	fmt.Println("Remaining balance is: ", formatMoney(originalBalance-withdrawlAmount))
	return originalBalance - withdrawlAmount
}

func formatMoney(moneyAmount float64) string {
	return fmt.Sprintf("$%.2f", moneyAmount)
}
