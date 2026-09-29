package main

import "fmt"

func main() {
	var name string
	var salary float64
	var transportationAllowance string
	var transportationAmount float64
	var salaryType string
	var salaryWithTransportation float64
	var monthlySaving float64
	var savingGoal float64
	var asratChoice string
	var asratAmount float64
	var trackExpense string
	var expenseAmount float64

	fmt.Println("What is your name?")
	fmt.Scanln(&name)

	fmt.Println("Do you get transportation allowance? (Enter 'Y/y' or 'N/n')")
	fmt.Scanln(&transportationAllowance)
	if transportationAllowance == "Y" || transportationAllowance == "y" {
		fmt.Println("How much is it?")
		fmt.Scanln(&transportationAmount)
		fmt.Println("Is transportation allowance included in your salary? or is it separate? (Enter 'included' or 'separate')")
		fmt.Scanln(&salaryType)

		if salaryType == "included" {
			fmt.Println("How much is your salary?")
			fmt.Scanln(&salaryWithTransportation)
			salary = salaryWithTransportation - transportationAmount
		} else {
			fmt.Println("How much is your salary?")
			fmt.Scanln(&salary)
		}
	} else {
		fmt.Println("How much is your salary?")
		fmt.Scanln(&salary)
	}
	tax := calculateTax(salary)
	pension := calculatePension(salary)

	fmt.Println("Do you take out Asrat? (Enter 'Y/y' or 'N/n')")
	fmt.Scanln(&asratChoice)
	if asratChoice == "Y" || asratChoice == "y" {
		asratAmount = calculateAsrat(salary)
	}

	netSalary := (salary - tax - pension) + transportationAmount - asratAmount

	fmt.Println("How much do you want to save monthly?")
	fmt.Scanln(&monthlySaving)
	fmt.Println("What is your saving goal?")
	fmt.Scanln(&savingGoal)

	fmt.Println("Dear", name, ", your net salary is:", netSalary, "and it will take you", savingGoal/monthlySaving, "months to reach your saving goal of", savingGoal)
	fmt.Println("Your income after saving is: ", netSalary-monthlySaving)

	fmt.Println("Do you want to check Your Expenses? (Enter 'Y/y' or 'N/n')")
	fmt.Scanln(&trackExpense)
	if trackExpense == "Y" || trackExpense == "y" {
		expenseAmount = calculateExpenses()
		fmt.Println("Your Total Expense is: ", expenseAmount)
		if expenseAmount > (netSalary - monthlySaving) {
			fmt.Println("Your expenses exceed your income after saving. Please review your expenses.")
		} else {
			fmt.Println("You are within your budget.")
			fmt.Println("Your remaining income after expenses and saving is: ", (netSalary-monthlySaving)-expenseAmount)
		}
	}
}

func calculateTax(salary float64) float64 {
	if salary <= 2000 {
		return 0
	} else if salary <= 4000 {
		return (salary - 2000) * 0.15
	} else if salary <= 7000 {
		return 300 + (salary-4000)*0.20
	} else if salary <= 10000 {
		return 900 + (salary-7000)*0.25
	} else if salary <= 14000 {
		return 1650 + (salary-10000)*0.30
	} else {
		return 2850 + (salary-14000)*0.35
	}
}

func calculatePension(salary float64) float64 {
	return salary * 0.07
}

func calculateAsrat(salary float64) float64 {
	return salary * 0.1
}

func calculateExpenses() float64 {
	fmt.Println("How many expense categories do you want to enter?")
	var numCategories int
	var totalExpenses float64
	var expenseAmount float64
	var expenseName string

	fmt.Scanln(&numCategories)
	for i := 1; i <= numCategories; i++ {
		fmt.Println("Enter expense category ", i, " name:")
		fmt.Scanln(&expenseName)
		fmt.Println("Enter ", expenseName, " expense amount: ")
		fmt.Scanln(&expenseAmount)
		totalExpenses += expenseAmount
	}

	return totalExpenses
}
