package main

import (
	"errors"
	"fmt"
)

func deductExpense(balance *float64, expense float64) (float64, error) {
	if expense <= 0 {
		return *balance, errors.New("expense must be greater than zero")
	}

	if expense > *balance {
		return *balance, errors.New("insufficient balance")
	}

	*balance -= expense
	// nil means there is no error.
	return *balance, nil
}

func addIncome(balance *float64, income float64) (float64, error) {
	if income <= 0 {
		return *balance, errors.New("invalid amount")
	}

	*balance += income
	return *balance, nil
}

func main() {
	balance := 5000.0

	_, err := addIncome(&balance, 2000)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	deductedAmount, err := deductExpense(&balance, 1500)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Final amount after add and deduct = ", deductedAmount)
}
