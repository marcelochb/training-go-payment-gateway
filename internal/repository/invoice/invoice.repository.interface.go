package repository

import "github.com/marcelochb/training-go-payment-gateway/internal/domain"

type IInvoiceRepository interface {
	Save(invoice *domain.InvoiceEntity) error
	FindByID(id string) (*domain.InvoiceEntity, error)
	FindByAccountID(accountID string) ([]*domain.InvoiceEntity, error)
	UpdateStatus(invoice *domain.InvoiceEntity) error
}
