package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

func (p *payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r *razorpay) pay(amount float32) {
	fmt.Println("Pay using razorpay:", amount)
}

type stripe struct{}

func (s *stripe) pay(amount float32) {
	fmt.Println("Pay using stripe:", amount)

}

func main() {
	// razorpayGw := &razorpay{}
	stripeGw := &stripe{}
	newPayment := payment{
		gateway: stripeGw,
		// gateway: razorpayGw,
	}
	newPayment.makePayment(1000)

}
