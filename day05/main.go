package main

import "fmt"

type Payer interface {
	CalculatePayment() float64
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

type Contractor struct {
	Name        string
	HoursWorked float64
	HourlyRate  float64
}

func (d Developer) CalculatePayment() float64 {
	return d.Salary
}

func (i Intern) CalculatePayment() float64 {
	return i.Stipend
}

func (m Manager) CalculatePayment() float64 {
	return m.Salary + m.Bonus
}

func calculateHourlyPayment(hours, rate float64) float64 {
	return hours * rate
}

func (f Freelancer) CalculatePayment() float64 {
	return calculateHourlyPayment(f.HoursWorked, f.HourlyRate)
}

func (c Contractor) CalculatePayment() float64 {
	return calculateHourlyPayment(c.HoursWorked, c.HourlyRate)
}

func calculateTotalPayments(payers []Payer) float64 {
	total := 0.0
	for _, payer := range payers {
		total += payer.CalculatePayment()
	}
	return total
}

func printPayment(p Payer) {
	fmt.Printf("Payment: %.2f ETB\n", p.CalculatePayment())
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

	contractor := Contractor{
		Name:        "Dawit",
		HoursWorked: 40,
		HourlyRate:  250,
	}
	printPayment(contractor)

	total := calculateTotalPayments([]Payer{dev, manager, intern, freelancer, contractor})
	fmt.Printf("Total Payments: %.2f ETB\n", total)
}
