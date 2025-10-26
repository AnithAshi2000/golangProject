package main

import (
	"errors"
	"fmt"
	"regexp"
)

type CreditCardPayment struct {
	CardNumber string
	CVV        string
	Expiry     string
}

func (c CreditCardPayment) Validate() error {
	if matched, _ := regexp.MatchString(`^\d{16}$`, c.CardNumber); !matched {
		return errors.New("invalid card number (must be 16 digits)")
	}
	if matched, _ := regexp.MatchString(`^\d{3}$`, c.CVV); !matched {
		return errors.New("invalid CVV (must be 3 digits)")
	}
	if matched, _ := regexp.MatchString(`^(0[1-9]|1[0-2])/[0-9]{2}$`, c.Expiry); !matched {
		return errors.New("invalid expiry format (MM/YY)")
	}
	return nil
}

func (c CreditCardPayment) Pay(amount float64) string {
	return fmt.Sprintf("Credit Card payment of Rs %.2f approved", amount)
}

func (c CreditCardPayment) Refund(amount float64) string {
	return fmt.Sprintf("Credit Card refund of Rs %.2f processed", amount)
}

func (c CreditCardPayment) GetMethodName() string {
	return "credit"
}
