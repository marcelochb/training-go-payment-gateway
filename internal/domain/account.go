package domain

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        string
	Name      string
	Email     string
	APIKey    string
	Balance   float64
	Mutex     *sync.Mutex
	CreatedAt time.Time
	UpdatedAt time.Time
}

func generateAPIKey() string {
	key := make([]byte, 16)
	_, err := rand.Read(key)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(key)
}

func NewAccount(name, email string) *Account {

	account := &Account{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		Balance:   0,
		APIKey:    generateAPIKey(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return account
}

func (account *Account) AddBalance(amount float64) {
	account.Mutex.Lock()
	defer account.Mutex.Unlock()
	account.Balance += amount
	account.UpdatedAt = time.Now()
}
