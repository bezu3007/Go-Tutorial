package main

import "fmt"

type Developer struct {
	Name   string
	Salary float64
}
type Manager struct {
	Name   string
	Salary float64
}
type Freelancer struct {
	Name   string
	Salary float64
}

type Payer interface {
	calculatePayment() float64
}

func (d Developer) calculatePayment() float64 {
	return d.Salary
}

func (m Manager) calculatePayment() float64 {
	return m.Salary + 5000 // Bonus for managers
}

func (f Freelancer) calculatePayment() float64 {
	return f.Salary
}

func printPayment(p Payer) {
	fmt.Println("Payment:", p.calculatePayment())
}

func printAllPayments(payers []Payer) {
	for _, p := range payers {
		printPayment(p)
	}
}

func main() {
	dev := Developer{
		Name:   "Alice",
		Salary: 15000,
	}

	manager := Manager{
		Name:   "Bob",
		Salary: 20000,
	}

	freelancer := Freelancer{
		Name:   "Charlie",
		Salary: 10000,
	}

	payers := []Payer{
		dev,
		manager,
		freelancer,
	}

	printAllPayments(payers)
}
