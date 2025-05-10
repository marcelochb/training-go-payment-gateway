package service

import (
	"github.com/marcelochb/training-go-payment-gateway/internal/domain"
	"github.com/marcelochb/training-go-payment-gateway/internal/dto"
	repository "github.com/marcelochb/training-go-payment-gateway/internal/repository/invoice"
)

type InvoiceService struct {
	invoiceRepository repository.InvoiceRepository
	accountService    AccountService
}

func NewInvoiceService(invoiceRepository repository.InvoiceRepository, accountService AccountService) *InvoiceService {
	return &InvoiceService{invoiceRepository: invoiceRepository, accountService: accountService}
}

func (s *InvoiceService) Create(input *dto.InvoiceDtoInput) (*dto.InvoiceDtoOutput, error) {
	accountOutput, err := s.accountService.FindByAPIKey(input.APIKey)
	if err != nil {
		return nil, err
	}
	invoice, err := dto.ToInvoiceEntity(input, accountOutput.ID)
	if err != nil {
		return nil, err
	}

	if err := invoice.Process(); err != nil {
		return nil, err
	}

	if invoice.Status == domain.InvoiceStatusApproved {
		_, err = s.accountService.UpdateBalance(input.APIKey, invoice.Amount)
		if err != nil {
			return nil, err
		}
	}

	err = s.invoiceRepository.Save(invoice)
	if err != nil {
		return nil, err
	}

	output := dto.FromInvoiceEntity(invoice)
	return &output, nil
}

func (s *InvoiceService) GetByID(id, apikey string) (*dto.InvoiceDtoOutput, error) {
	invoice, err := s.invoiceRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	accountOutput, err := s.accountService.FindByAPIKey(apikey)
	if err != nil {
		return nil, err
	}

	if invoice.AccountID != accountOutput.ID {
		return nil, domain.ErrUnauthorizedAccess
	}

	output := dto.FromInvoiceEntity(invoice)
	return &output, nil
}

func (s *InvoiceService) ListByAccount(accountID string) ([]*dto.InvoiceDtoOutput, error) {
	invoices, err := s.invoiceRepository.FindByAccountID(accountID)
	if err != nil {
		return nil, err
	}

	output := make([]*dto.InvoiceDtoOutput, len(invoices))
	for i, invoice := range invoices {
		output[i] = &dto.InvoiceDtoOutput{
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
	return output, nil
}

func (s *InvoiceService) ListByAccountAPIKey(apikey string) ([]*dto.InvoiceDtoOutput, error) {
	accountOutput, err := s.accountService.FindByAPIKey(apikey)
	if err != nil {
		return nil, err
	}

	return s.ListByAccount(accountOutput.ID)
}
