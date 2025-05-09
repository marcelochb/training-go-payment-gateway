package dto

import (
	"time"

	"github.com/marcelochb/training-go-payment-gateway/internal/domain"
)

type AccountDtoInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AccountDtoOuput struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Balance   float64   `json:"balance"`
	APIKey    string    `json:"api_key,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToDomain(dto AccountDtoInput) *domain.AccountEntity {
	return domain.NewAccount(dto.Name, dto.Email)
}

func FromDomain(account *domain.AccountEntity) AccountDtoOuput {
	return AccountDtoOuput{
		ID:        account.ID,
		Name:      account.Name,
		Email:     account.Email,
		Balance:   account.Balance,
		APIKey:    account.APIKey,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}
}
