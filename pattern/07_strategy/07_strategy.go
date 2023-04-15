/*
Паттерн Стратегия относится к поведенческим паттернам.
Он определяет набор алгоритмов схожих по роду деятельности и делает их подменяемыми.
Позволяет подменять алгоритмы без участия клиентов, которые используют эти алгоритмы.

Этот паттерн можно использовать в платежах.
Допустим, мы хотим использовать разные алгоритмы для обработки платежей.
*/

package main

import "fmt"

type PaymentStrategy interface {
	Pay(amount float64)
}

type CreditCardStrategy struct {
	cardNumber     string
	expirationDate string
	cvv            string
}

func (c *CreditCardStrategy) Pay(amount float64) {
	fmt.Printf("Paying %f by card...\n", amount)
}

type PayPalStrategy struct {
	email    string
	password string
}

func (p *PayPalStrategy) Pay(amount float64) {
	fmt.Printf("Paying %f with paypal...\n", amount)
}

type Payment struct {
	strategy PaymentStrategy
}

func (p *Payment) ProcessPayment(amount float64) {
	p.strategy.Pay(amount)
}

func main() {
	creditCard := CreditCardStrategy{
		cardNumber:     "0000 0000 0000 0000",
		expirationDate: "03/29",
		cvv:            "420",
	}
	payment := &Payment{strategy: &creditCard}
	payment.ProcessPayment(1000.0)

	payPal := PayPalStrategy{
		email:    "wb@wb.ru",
		password: "samurai",
	}
	payment = &Payment{strategy: &payPal}
	payment.ProcessPayment(1000.0)
}
