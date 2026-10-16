package main

import "fmt"

type Payer interface {
	GetSalary() float64
}

type Developer struct {
	Name   string
	Salary float64
}
type Intern struct {
	Name    string
	Stipend float64
}

type Manager struct {
	Name   string
	Salary float64
	Bonus  float64
}

type Freelancer struct {
	Name        string
	HoursWorked float64
	HourlyRate  float64
}

func (d Developer) GetSalary() float64 {
	return d.Salary
}

func (i Intern) GetSalary() float64 {
	return i.Stipend
}

func (m Manager) GetSalary() float64 {
	return m.Salary + m.Bonus
}

func (f Freelancer) GetSalary() float64 {
	return f.HoursWorked * f.HourlyRate
}

func printPayment(p Payer) {
	fmt.Printf("Payment: %.2f ETB\n", p.GetSalary())
}

func main() {
	dev := Developer{
		Name:   "Alice",
		Salary: 15000,
	}

	manager := Manager{
		Name:   "Bob",
		Salary: 20000,
		Bonus:  5000,
	}

	intern := Intern{
		Name:    "David",
		Stipend: 5000,
	}

	freelancer := Freelancer{
		Name:        "Charlie",
		HoursWorked: 80,
		HourlyRate:  150,
	}

	printPayment(dev)
	printPayment(manager)
	printPayment(intern)
	printPayment(freelancer)
}
