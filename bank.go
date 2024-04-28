package main

import "fmt"

func main() {
	var accountBalance float64 = 1000
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Enter a choice: ")
	fmt.Scan(&choice)

	fmt.Println("Your choice: ", choice)
	switch choice {
	case 1:
		fmt.Println("Your balance is ", formatMoney(accountBalance))
	case 2:
		accountBalance = depositFunds(accountBalance)
	case 3:
		var withdrawlAmount float64
		fmt.Print("Enter an amount to withdraw: ")
		fmt.Scan(&withdrawlAmount)
		fmt.Println(fmt.Sprintf("%2f", withdrawlAmount), " withdrawn.")
	}
}

func depositFunds(originalBalance float64) float64 {
	var depositAmount float64

	fmt.Print("Enter an amount to deposit: ")
	fmt.Scan(&depositAmount)

	fmt.Println(formatMoney(depositAmount), " deposited!")
	fmt.Println("New balance is ", formatMoney(originalBalance+depositAmount))

	return originalBalance + depositAmount
}

func withdrawlFunds(originalBalance float64) float64 {
	var withdrawlAmount float64
	fmt.Print("Enter an amount to withdraw: ")
	fmt.Scan(&withdrawlAmount)

	if withdrawlAmount > originalBalance {
		fmt.Println("Cannot withdraw more than the existing balance. Current balance is ", formatMoney(originalBalance))
		return originalBalance
	}

	fmt.Println(formatMoney(withdrawlAmount), " withdrawn.")
	fmt.Println("Remaining balance is: ", formatMoney(originalBalance-withdrawlAmount))
	return originalBalance - withdrawlAmount
}

func formatMoney(moneyAmount float64) string {
	return fmt.Sprintf("$%2f", moneyAmount)
}
