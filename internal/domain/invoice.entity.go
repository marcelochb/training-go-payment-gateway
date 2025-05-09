package domain

import (
	"math/rand/v2"
	"time"
)

type InvoiceStatus string

const (
	InvoiceStatusPending  InvoiceStatus = "pending"
	InvoiceStatusApproved InvoiceStatus = "approved"
	InvoiceStatusRejected InvoiceStatus = "rejected"
)

type InvoiceEntity struct {
	ID             string
	AccountID      string
	Status         InvoiceStatus
	Description    string
	PaymentType    string
	CardLastDigits string
	Amount         float64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type CreditCardEntity struct {
	Number         string
	CVV            string
	ExpireMonth    int
	ExpireYear     int
	CardholderName string
}

func NewInvoiceEntity(accountID string, amount float64, description string, paymentType string, card CreditCardEntity) (*InvoiceEntity, error) {
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}

	lastDigits := card.Number[len(card.Number)-4:]

	return &InvoiceEntity{
		AccountID:      accountID,
		Amount:         amount,
		Description:    description,
		Status:         InvoiceStatusPending,
		PaymentType:    paymentType,
		CardLastDigits: lastDigits,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil
}

func (invoice *InvoiceEntity) Process() error {
	if invoice.Amount > 10000 {
		return nil
	}

	randomSource := rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), uint64(time.Now().UnixNano())))
	var newStatus InvoiceStatus
	if randomSource.Float64() > 0.7 {
		newStatus = InvoiceStatusApproved
	} else {
		newStatus = InvoiceStatusRejected
	}

	invoice.Status = newStatus
	return nil
}

func (invoice *InvoiceEntity) UpdateStatus(newStatus InvoiceStatus) error {
	if newStatus == InvoiceStatusPending {
		return ErrInvalidStatus
	}

	invoice.Status = newStatus
	invoice.UpdatedAt = time.Now()
	return nil
}
