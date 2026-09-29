// Today’s path:

// 🔹 Functions — parameters and return values
// 🔹 Multiple return values
// 🔹 Structs — creating your own data types
// 🔹 Methods
// 🔹 Small practical project: Employee Salary Calculator
// 🔹 Quick quiz/checkpoint

package main

import "fmt"

type Employee struct {
	Name           string
	Salary         float64
	Transportation float64
}

func (e Employee) taxableSalary() float64 {
	return e.Salary - e.Transportation
}

func (e Employee) calculateTax() float64 {

	taxable := e.taxableSalary()
	if taxable <= 2000 {
		return 0
	} else if taxable <= 4000 {
		return (taxable - 2000) * 0.15
	} else if taxable <= 7000 {
		return 300 + (taxable-4000)*0.20
	} else if taxable <= 10000 {
		return 900 + (taxable-7000)*0.25
	} else if taxable <= 14000 {
		return 1650 + (taxable-10000)*0.30
	} else {
		return 2850 + (taxable-14000)*0.35
	}
}
func (e Employee) calculatePension() float64 {
	taxable := e.taxableSalary()
	return taxable * 0.07
}

func (e Employee) calculateTaxAndPension() (float64, float64) {
	return e.calculateTax(), e.calculatePension()
}

func (e Employee) netSalary() float64 {
	tax, pension := e.calculateTaxAndPension()
	fmt.Println("Calculated Tax:", tax)
	fmt.Println("Calculated Pension:", pension)
	return e.Salary - tax - pension
}

func main() {
	employee := Employee{
		Name:           "John Doe",
		Salary:         30000,
		Transportation: 2200,
	}
	net := employee.netSalary()
	fmt.Println("Salary:", employee.Salary)
	fmt.Println("Dear", employee.Name, ", your net salary is:", net)
}
