package main

import "fmt"

// Strategy Interface
type PaymentStrategy interface {
	Pay(amount float32)
}

// Concrete Strategies
type CreditCardPayment struct{}

func (c *CreditCardPayment) Pay(amount float32) {
	fmt.Printf("Paid %f using Credit Card\n", amount)
}

type PaypalPayment struct{}

func (p *PaypalPayment) Pay(amount float32) {
	fmt.Printf("Paid %f using Paypal\n", amount)
}

// Context
type PaymentContext struct {
	strategy PaymentStrategy
}

func (p *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}
func (p *PaymentContext) Pay(amount float32) {
	p.strategy.Pay(amount)
}

// Client code
func main() {
	payment := PaymentContext{}
	payment.SetStrategy(&CreditCardPayment{})
	payment.Pay(22.30)

	payment.SetStrategy(&PaypalPayment{})
	payment.Pay(17.50)
}
