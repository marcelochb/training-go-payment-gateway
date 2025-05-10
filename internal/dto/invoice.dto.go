package dto

import (
	"time"

	"github.com/marcelochb/training-go-payment-gateway/internal/domain"
)

const (
	StatusPending  = string(domain.InvoiceStatusPending)
	StatusApproved = string(domain.InvoiceStatusApproved)
	StatusRejected = string(domain.InvoiceStatusRejected)
)

type InvoiceDtoInput struct {
	APIKey         string
	Amount         float64 `json:"amount"`
	Description    string  `json:"description"`
	PaymentType    string  `json:"payment_type"`
	CardLastDigits string  `json:"card_last_digits"`
	CardNumber     string  `json:"card_number"`
	CVV            string  `json:"cvv"`
	ExpireMonth    int     `json:"expire_month"`
	ExpireYear     int     `json:"expire_year"`
	CardholderName string  `json:"cardholder_name"`
}

type InvoiceDtoOutput struct {
	ID             string    `json:"id"`
	AccountID      string    `json:"account_id"`
	Status         string    `json:"status"`
	Description    string    `json:"description"`
	PaymentType    string    `json:"payment_type"`
	CardLastDigits string    `json:"card_last_digits"`
	Amount         float64   `json:"amount"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func ToInvoiceEntity(dto *InvoiceDtoInput, accountID string) (*domain.InvoiceEntity, error) {
	card := domain.CreditCardEntity{
		Number:         dto.CardNumber,
		CVV:            dto.CVV,
		ExpireMonth:    dto.ExpireMonth,
		ExpireYear:     dto.ExpireYear,
		CardholderName: dto.CardholderName,
	}

	return domain.NewInvoiceEntity(accountID, dto.Amount, dto.Description, dto.PaymentType, card)
}

func FromInvoiceEntity(invoice *domain.InvoiceEntity) InvoiceDtoOutput {
	return InvoiceDtoOutput{
		ID:             invoice.ID,
		AccountID:      invoice.AccountID,
		Status:         string(invoice.Status),
		Description:    invoice.Description,
		PaymentType:    invoice.PaymentType,
		CardLastDigits: invoice.CardLastDigits,
		Amount:         invoice.Amount,
		CreatedAt:      invoice.CreatedAt,
		UpdatedAt:      invoice.UpdatedAt,
	}
}
