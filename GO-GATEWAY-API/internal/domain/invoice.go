package domain

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusPending  Status = "pending"
	StatusApproved Status = "approved"
	StatusRejected Status = "rejected"
)

type Invoice struct {
	ID             string
	AccountId      string
	Amount         float64
	Status         Status
	Description    string
	PaymentType    string
	CardLastDigits string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type CreditCard struct {
	Number         string
	CVV            string
	ExpireMonth    int
	ExpireYear     int
	CardHolderName string
}

func NewInvoice(
	accountId string,
	amount float64,
	description string,
	paymentType string,
	card CreditCard,
) (*Invoice, error) {
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}
	if accountId == "" {
		return nil, ErrAccountNotFound
	}
	if description == "" {
		return nil, ErrInvalidDescription
	}
	if card.Number == "" || len(card.Number) < 16 || card.CVV == "" || card.ExpireMonth == 0 || card.ExpireYear == 0 || card.CardHolderName == "" {
		return nil, ErrInvalidCreditCard
	}

	lastDigits := card.Number[len(card.Number)-4:]

	return &Invoice{
		ID:             uuid.New().String(),
		AccountId:      accountId,
		Amount:         amount,
		Status:         StatusPending,
		Description:    description,
		PaymentType:    paymentType,
		CardLastDigits: lastDigits,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil
}

func (i *Invoice) Process() error {
	if i.Amount > 10000 {
		return nil
	}

	randomSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	var newStatus Status
	if randomSource.Float64() <= 0.7 {
		newStatus = StatusApproved
	} else {
		newStatus = StatusRejected
	}

	i.Status = newStatus
	return nil
}

func (i *Invoice) UpdateStatus(newStatus Status) error {
	if newStatus != StatusPending {
		return ErrInvalidStatus
	}

	i.Status = newStatus
	i.UpdatedAt = time.Now()
	return nil
}
