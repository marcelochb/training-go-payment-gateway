package repository

import "github.com/marcelochb/training-go-payment-gateway/internal/domain"

type IAccountRepository interface {
	Save(account *domain.AccountEntity) error
	FindByAPIKey(apiKey string) (*domain.AccountEntity, error)
	FindByID(id string) (*domain.AccountEntity, error)
	UpdateBalance(account *domain.AccountEntity) error
}
