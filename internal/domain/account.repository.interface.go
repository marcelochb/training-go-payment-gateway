package domain

type IAccountRepository interface {
	Save(account *AccountEntity) error
	FindByAPIKey(apiKey string) (*AccountEntity, error)
	FindByID(id string) (*AccountEntity, error)
	UpdateBalance(account *AccountEntity) error
}
