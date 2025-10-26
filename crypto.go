package main

import (
	"errors"
	"fmt"
)

type CryptoPayment struct {
	WalletAddress string
}

func (c CryptoPayment) Validate() error {
	if len(c.WalletAddress) < 10 {
		return errors.New("invalid wallet address (too short)")
	}
	return nil
}

func (c CryptoPayment) Pay(amount float64) string {
	return fmt.Sprintf("Crypto payment of %.2f confirmed on blockchain", amount)
}

func (c CryptoPayment) Refund(amount float64) string {
	return fmt.Sprintf("Crypto refund of %.2f processed", amount)
}

func (c CryptoPayment) GetMethodName() string {
	return "crypto"
}
