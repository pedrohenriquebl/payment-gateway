package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/pedrohenriquebl/gateway/internal/dto"
	"github.com/pedrohenriquebl/gateway/internal/service"
)

type AccountHandler struct {
	accountService service.AccountServiceInterface
}

func NewAccountHandler(accountService service.AccountServiceInterface) *AccountHandler {
	return &AccountHandler{accountService: accountService}
}

func (handler *AccountHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateAccountInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := handler.accountService.CreateAccount(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (handler *AccountHandler) Get(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-API-Key")
	if apiKey == "" {
		http.Error(w, "API Key is required", http.StatusUnauthorized)
		return
	}

	output, err := handler.accountService.FindByAPIKey(apiKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
