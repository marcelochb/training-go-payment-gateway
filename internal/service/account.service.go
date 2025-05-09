package service

import (
	"github.com/marcelochb/training-go-payment-gateway/internal/domain"
	"github.com/marcelochb/training-go-payment-gateway/internal/dto"
)

type AccountService struct {
	repository domain.IAccountRepository
}

func NewAccountService(repository domain.IAccountRepository) *AccountService {
	return &AccountService{repository: repository}
}

func (service *AccountService) CreateAccount(input dto.AccountDtoInput) (*dto.AccountDtoOuput, error) {
	account := dto.ToDomain(input)
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
	output := dto.FromDomain(account)
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
	output := dto.FromDomain(account)
	return &output, nil
}

func (service *AccountService) FindByAPIKey(apikey string) (*dto.AccountDtoOuput, error) {
	account, err := service.repository.FindByAPIKey(apikey)
	if err != nil {
		return nil, err
	}
	output := dto.FromDomain(account)
	return &output, nil
}

func (service *AccountService) FindByID(id string) (*dto.AccountDtoOuput, error) {
	account, err := service.repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	output := dto.FromDomain(account)
	return &output, nil
}
