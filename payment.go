package main

type PaymentMethod interface {
	Pay(amount float64) string
	Refund(amount float64) string
	Validate() error
	GetMethodName() string
}

type PaymentInfo struct {
	ID     string
	Amount float64
	Method string
	Status string
}
