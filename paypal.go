package main

import (
	"errors"
	"fmt"
	"regexp"
)

type PayPalPayment struct {
	Email string
	Token string
}

func (p PayPalPayment) Validate() error {
	if matched, _ := regexp.MatchString(`^[\w\.-]+@[\w\.-]+\.\w+$`, p.Email); !matched {
		return errors.New("invalid email address")
	}
	if len(p.Token) < 5 {
		return errors.New("auth token too short")
	}
	return nil
}

func (p PayPalPayment) Pay(amount float64) string {
	return fmt.Sprintf("PayPal payment of $%.2f sent successfully", amount)
}

func (p PayPalPayment) Refund(amount float64) string {
	return fmt.Sprintf("PayPal refund of $%.2f issued", amount)
}

func (p PayPalPayment) GetMethodName() string {
	return "paypal"
}
