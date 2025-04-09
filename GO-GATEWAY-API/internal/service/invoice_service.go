package service

import (
	"github.com/pedrohenriquebl/gateway/internal/domain"
	"github.com/pedrohenriquebl/gateway/internal/dto"
)

type InvoiceServiceInterface interface {
	Create(input dto.CreateInvoiceInput) (*dto.InvoiceOutput, error)
	Save(invoice *dto.CreateInvoiceInput) (*dto.InvoiceOutput, error)
	GetById(id, apiKey string) (*dto.InvoiceOutput, error)
	ListByAccount(accountID string) ([]*dto.InvoiceOutput, error)
	ListByAPIKey(apiKey string) ([]*dto.InvoiceOutput, error)
}

type InvoiceService struct {
	invoiceRepository domain.InvoiceRepository
	AccountService    AccountService
}

func NewInvoiceService(invoiceRepository domain.InvoiceRepository, accountService AccountService) *InvoiceService {
	return &InvoiceService{
		invoiceRepository: invoiceRepository,
		AccountService:    accountService,
	}
}

func (s *InvoiceService) Create(input dto.CreateInvoiceInput) (*dto.InvoiceOutput, error) {
	accountOutput, err := s.AccountService.FindByAPIKey(input.APIKey)
	if err != nil {
		return nil, err
	}

	invoice, err := dto.ToInvoice(input, accountOutput.ID)
	if err != nil {
		return nil, err
	}

	if err := invoice.Process(); err != nil {
		return nil, err
	}

	if invoice.Status == domain.StatusApproved {
		_, err = s.AccountService.UpdateBalance(input.APIKey, invoice.Amount)
		if err != nil {
			return nil, err
		}
	}

	if err := s.invoiceRepository.Save(invoice); err != nil {
		return nil, err
	}

	return dto.FromInvoice(invoice), nil
}

func (s *InvoiceService) GetById(id, apiKey string) (*dto.InvoiceOutput, error) {
	invoice, err := s.invoiceRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	accountOutput, err := s.AccountService.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	if accountOutput.ID != invoice.AccountId {
		return nil, domain.ErrUnauthorizedAccess
	}

	return dto.FromInvoice(invoice), nil
}

func (s *InvoiceService) ListByAccount(accountID string) ([]*dto.InvoiceOutput, error) {
	invoices, err := s.invoiceRepository.FindByAccountID(accountID)
	if err != nil {
		return nil, err
	}

	output := make([]*dto.InvoiceOutput, len(invoices))
	for i, invoice := range invoices {
		output[i] = dto.FromInvoice(invoice)
	}

	return output, nil
}

func (s *InvoiceService) ListByAPIKey(apiKey string) ([]*dto.InvoiceOutput, error) {
	accountOutput, err := s.AccountService.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	invoices, err := s.invoiceRepository.FindByAccountID(accountOutput.ID)
	if err != nil {
		return nil, err
	}

	output := make([]*dto.InvoiceOutput, len(invoices))
	for i, invoice := range invoices {
		output[i] = dto.FromInvoice(invoice)
	}

	return output, nil
}
