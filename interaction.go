package main

import "fmt"

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
