package main

type PaymentGateway struct {
	logs []Transaction
}

func NewPaymentGateway() *PaymentGateway {
	return &PaymentGateway{}
}

func (pg *PaymentGateway) Process(method PaymentMethod, amount float64, action string) Transaction {
	var message string
	if action == "pay" {
		message = method.Pay(amount)
	} else {
		message = method.Refund(amount)
	}

	transaction := Transaction{
		ID:     GenerateID(),
		Method: method.GetMethodName(),
		Amount: amount,
		Action: action,
		Status: "success",
		Note:   message,
	}

	pg.logs = append(pg.logs, transaction)
	Log(message)
	return transaction
}
