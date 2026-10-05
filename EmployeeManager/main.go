package main

import "fmt"

type Employee struct {
	Name   string
	Salary float64
}

func (e *Employee) increaseSalary(amount float64) {
	e.Salary += amount
}

func (e *Employee) decreaseSalary(amount float64) {
	e.Salary -= amount
}

func (e Employee) displayInfo() {
	fmt.Println("Employee Name:", e.Name)
	fmt.Println("Employee Salary:", e.Salary)
}

func (e Employee) annualSalary() float64 {
	return e.Salary * 12
}

func main() {
	employee := Employee{
		Name:   "John Doe",
		Salary: 15000,
	}
	employee.displayInfo()

	employee.increaseSalary(3000)
	fmt.Println("Salary after increase:", employee.Salary)
	employee.decreaseSalary(1000)
	fmt.Println("Salary after decrease:", employee.Salary)

	annual := employee.annualSalary()
	fmt.Println("Annual Salary:", annual)
}
