 package main

import "fmt"

// ==========================================
// Interface Definition
// ==========================================

type Payment interface {
	Pay(amount float64)
}

// ==========================================
// UPI Payment
// ==========================================

type UPI struct {
	ID string
}

func (u UPI) Pay(amount float64) {
	fmt.Printf("Paid ₹%.2f using UPI (ID: %s)\n", amount, u.ID)
}

// ==========================================
// Credit Card Payment
// ==========================================

type CreditCard struct {
	CardNumber string
}

func (c CreditCard) Pay(amount float64) {
	fmt.Printf("Paid ₹%.2f using Credit Card (%s)\n", amount, c.CardNumber)
}

// ==========================================
// PayPal Payment
// ==========================================

type PayPal struct {
	Email string
}

func (p PayPal) Pay(amount float64) {
	fmt.Printf("Paid ₹%.2f using PayPal (%s)\n", amount, p.Email)
}

// ==========================================
// Function Accepting Interface
// ==========================================

func MakePayment(p Payment, amount float64) {
	p.Pay(amount)
}

func main() {

	fmt.Println("========================================")
	fmt.Println("      Go Interfaces - Complete Guide")
	fmt.Println("========================================")

	// ----------------------------------------
	// Example 1
	// ----------------------------------------

	upi := UPI{
		ID: "rajnish@upi",
	}

	fmt.Println("\n1. UPI Payment")

	MakePayment(upi, 500)

	// ----------------------------------------
	// Example 2
	// ----------------------------------------

	card := CreditCard{
		CardNumber: "1234-5678-9012-3456",
	}

	fmt.Println("\n2. Credit Card Payment")

	MakePayment(card, 1200)

	// ----------------------------------------
	// Example 3
	// ----------------------------------------

	paypal := PayPal{
		Email: "rajnish@gmail.com",
	}

	fmt.Println("\n3. PayPal Payment")

	MakePayment(paypal, 2000)

	// ----------------------------------------
	// Example 4
	// ----------------------------------------

	var payment Payment

	payment = upi

	fmt.Println("\n4. Interface Variable")

	payment.Pay(100)

	// ----------------------------------------
	// Example 5
	// ----------------------------------------

	fmt.Println("\n5. Type Assertion")

	if u, ok := payment.(UPI); ok {
		fmt.Println("UPI ID:", u.ID)
	}

	// ----------------------------------------
	// Example 6
	// ----------------------------------------

	fmt.Println("\n6. Type Switch")

	checkPaymentType(upi)
	checkPaymentType(card)
	checkPaymentType(paypal)

}

func checkPaymentType(payment Payment) {

	switch p := payment.(type) {

	case UPI:
		fmt.Println("UPI Payment:", p.ID)

	case CreditCard:
		fmt.Println("Credit Card:", p.CardNumber)

	case PayPal:
		fmt.Println("PayPal:", p.Email)

	default:
		fmt.Println("Unknown Payment")

	}

}