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
	CreatedAt time.Time
	UpdatedAt time.Time
	mu        sync.RWMutex
}

func GenerateAPIKey() string {
	randomAPIKeyOnHex := make([]byte, 16)
	rand.Read(randomAPIKeyOnHex)
	return hex.EncodeToString(randomAPIKeyOnHex)
}

func NewAccount(name, email string) *Account {

	account := &Account{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		Balance:   0,
		mu:        sync.RWMutex{},
		APIKey:    GenerateAPIKey(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return account
}

func (account *Account) AddBalance(amount float64) {
	account.mu.Lock()
	defer account.mu.Unlock()
	account.Balance += amount
	account.UpdatedAt = time.Now()
}
