package main

import "fmt"

type Bank struct {
	AccountNumber string
}

func (b Bank) Pay(amount float64) {
	fmt.Println("Paid", amount, "using Bank Account: ", b.AccountNumber)
}

type TeleBirr struct {
	phoneNumber string
}

func (t TeleBirr) Pay(amount float64) {
	fmt.Println("Paid", amount, "using TeleBirr: ", t.phoneNumber)
}

type Cash struct{}

func (c Cash) Pay(amount float64) {
	fmt.Println("Paid", amount, "using Cash")
}

type PaymentMethod interface {
	Pay(amount float64)
}

func processPayment(p PaymentMethod, amount float64) {
	fmt.Println("Processing payment of amount:", amount)
	p.Pay(amount)
}

func main() {
	bank := Bank{AccountNumber: "12345"}
	teleBirr := TeleBirr{phoneNumber: "0912345678"}
	cash := Cash{}

	payments := []PaymentMethod{
		bank,
		teleBirr,
		cash,
	}

	processAllPayments(payments, 1000)
}

func processAllPayments(payments []PaymentMethod, amount float64) {
	for _, payment := range payments {
		processPayment(payment, amount)
	}
}
