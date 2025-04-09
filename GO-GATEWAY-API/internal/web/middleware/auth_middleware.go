package middleware

import (
	"net/http"

	"github.com/pedrohenriquebl/gateway/internal/domain"
	"github.com/pedrohenriquebl/gateway/internal/service"
)

type AuthMiddleware struct {
	accountService service.AccountServiceInterface
}

func NewAuthMiddleware(accountService service.AccountServiceInterface) *AuthMiddleware {
	return &AuthMiddleware{
		accountService: accountService,
	}
}

func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")
		if apiKey == "" {
			http.Error(w, "API Key is required", http.StatusUnauthorized)
			return
		}

		_, err := (m.accountService).FindByAPIKey(apiKey)
		if err != nil {
			if err == domain.ErrAccountNotFound {
				http.Error(w, "Invalid API Key", http.StatusUnauthorized)
			}

			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		next.ServeHTTP(w, r)
	})
}
