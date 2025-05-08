package domain

type AccountRepository interface {
	Save(account *Account)
	FindByAPIKey(apiKey string) (*Account, error)
}
