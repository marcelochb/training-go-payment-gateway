package service

import (
	"github.com/marcelochb/training-go-payment-gateway/internal/domain"
	"github.com/marcelochb/training-go-payment-gateway/internal/dto"
	repository "github.com/marcelochb/training-go-payment-gateway/internal/repository/account"
)

type AccountService struct {
	repository repository.IAccountRepository
}

func NewAccountService(repository repository.IAccountRepository) *AccountService {
	return &AccountService{repository: repository}
}

func (service *AccountService) CreateAccount(input dto.AccountDtoInput) (*dto.AccountDtoOuput, error) {
	account := dto.ToAccountEntity(input)
	existingAccount, err := service.repository.FindByAPIKey(account.APIKey)
	if err != nil && err != domain.ErrAccountNotFound {
		return nil, err
	}
	if existingAccount != nil {
		return nil, domain.ErrDuplicatedAPIKey
	}
	err = service.repository.Save(account)
	if err != nil {
		return nil, err
	}
	output := dto.FromAccountEntity(account)
	return &output, nil
}

func (service *AccountService) UpdateBalance(apikey string, amount float64) (*dto.AccountDtoOuput, error) {
	account, err := service.repository.FindByAPIKey(apikey)
	if err != nil {
		return nil, err
	}
	account.AddBalance(amount)
	err = service.repository.UpdateBalance(account)
	if err != nil {
		return nil, err
	}
	output := dto.FromAccountEntity(account)
	return &output, nil
}

func (service *AccountService) FindByAPIKey(apikey string) (*dto.AccountDtoOuput, error) {
	account, err := service.repository.FindByAPIKey(apikey)
	if err != nil {
		return nil, err
	}
	output := dto.FromAccountEntity(account)
	return &output, nil
}

func (service *AccountService) FindByID(id string) (*dto.AccountDtoOuput, error) {
	account, err := service.repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	output := dto.FromAccountEntity(account)
	return &output, nil
}
