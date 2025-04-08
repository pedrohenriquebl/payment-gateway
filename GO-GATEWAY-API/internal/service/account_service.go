package service

import (
	"github.com/pedrohenriquebl/gateway/internal/domain"
	"github.com/pedrohenriquebl/gateway/internal/dto"
)

type AccountServiceInterface interface {
	CreateAccount(input dto.CreateAccountInput) (*dto.AccountOutput, error)
	UpdateBalance(apiKey string, amount float64) (*dto.AccountOutput, error)
	FindByAPIKey(apiKey string) (*dto.AccountOutput, error)
	FindByID(id string) (*dto.AccountOutput, error)
}

type AccountService struct {
	accountRepository domain.AccountRepository
}

func NewAccountService(accountRepository domain.AccountRepository) *AccountService {
	return &AccountService{accountRepository: accountRepository}
}

func (service *AccountService) CreateAccount(input dto.CreateAccountInput) (*dto.AccountOutput, error) {
	account := dto.ToAccount(input)

	existingAccount, err := service.accountRepository.FindByAPIKey(account.APIKey)
	if err != nil && err != domain.ErrAccountNotFound {
		return nil, err
	}
	if existingAccount != nil {
		return nil, domain.ErrDuplicatedAPIKey
	}
	err = service.accountRepository.Save(account)
	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account)
	return &output, nil
}

func (service *AccountService) UpdateBalance(apiKey string, amount float64) (*dto.AccountOutput, error) {
	account, err := service.accountRepository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	account.AddBalance(amount)
	err = service.accountRepository.UpdateBalance(account)
	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account)
	return &output, nil
}

func (service *AccountService) FindByAPIKey(apiKey string) (*dto.AccountOutput, error) {
	account, err := service.accountRepository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account)
	return &output, nil
}

func (service *AccountService) FindByID(id string) (*dto.AccountOutput, error) {
	account, err := service.accountRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account)
	return &output, nil
}
