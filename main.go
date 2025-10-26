package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Payment Gateway ===")

	fmt.Println("Select Payment Method:")
	fmt.Println("1. Credit Card")
	fmt.Println("2. PayPal")
	fmt.Println("3. Crypto")

	var methodChoice int
	fmt.Print("Enter choice (1-3): ")
	fmt.Scan(&methodChoice)

	var paymentType string
	switch methodChoice {
	case 1:
		paymentType = "credit"
	case 2:
		paymentType = "paypal"
	case 3:
		paymentType = "crypto"
	default:
		fmt.Println("Invalid payment method.")
		return
	}

	var method PaymentMethod
	switch paymentType {
	case "credit":
		var cardNum, cvv, expiry string
		fmt.Print("Enter Card Number (16 digits): ")
		fmt.Scan(&cardNum)
		fmt.Print("Enter CVV (3 digits): ")
		fmt.Scan(&cvv)
		fmt.Print("Enter Expiry (MM/YY): ")
		fmt.Scan(&expiry)
		method = CreditCardPayment{CardNumber: cardNum, CVV: cvv, Expiry: expiry}

	case "paypal":
		var email, token string
		fmt.Print("Enter PayPal Email: ")
		fmt.Scan(&email)
		fmt.Print("Enter Auth Token: ")
		fmt.Scan(&token)
		method = PayPalPayment{Email: email, Token: token}

	case "crypto":
		var wallet string
		fmt.Print("Enter Wallet Address: ")
		fmt.Scan(&wallet)
		method = CryptoPayment{WalletAddress: wallet}
	}

	if err := method.Validate(); err != nil {
		fmt.Println("Validation failed:", err)
		return
	}

	fmt.Println("\nSelect Action:")
	fmt.Println("1. Make Payment")
	fmt.Println("2. Refund")

	var actionChoice int
	fmt.Print("Enter choice (1-2): ")
	fmt.Scan(&actionChoice)

	var action string
	switch actionChoice {
	case 1:
		action = "pay"
	case 2:
		action = "refund"
	default:
		fmt.Println("Invalid action.")
		return
	}

	var amount float64
	fmt.Print("\nEnter amount: ")
	fmt.Scan(&amount)
	if amount <= 0 {
		fmt.Println("Amount must be greater than 0.")
		return
	}

	gateway := NewPaymentGateway()
	transaction := gateway.Process(method, amount, action)

	fmt.Println("\nTransaction Summary")
	fmt.Printf("ID: %s\nMethod: %s\nAction: %s\nAmount: %.2f\nStatus: %s\nNote: %s\n",
		transaction.ID, transaction.Method, transaction.Action, transaction.Amount, transaction.Status, transaction.Note)
}
