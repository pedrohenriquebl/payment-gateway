package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pedrohenriquebl/gateway/internal/service"
	"github.com/pedrohenriquebl/gateway/internal/web/handlers"
	"github.com/pedrohenriquebl/gateway/internal/web/middleware"
)

type Server struct {
	router         *chi.Mux
	server         *http.Server
	accountService service.AccountServiceInterface
	invoiceService service.InvoiceServiceInterface
	port           string
}

func NewServer(accountService service.AccountServiceInterface, invoiceService service.InvoiceServiceInterface, port string) *Server {
	return &Server{
		router:         chi.NewRouter(),
		accountService: accountService,
		invoiceService: invoiceService,
		port:           port,
	}
}

func (server *Server) ConfigureRoutes() {
	accountHandler := handlers.NewAccountHandler(server.accountService)
	invoiceHandler := handlers.NewInvoiceHandler(server.invoiceService)
	authMiddleware := middleware.NewAuthMiddleware(server.accountService)

	server.router.Post("/accounts", accountHandler.Create)
	server.router.Get("/accounts", accountHandler.Get)

	server.router.Group(func(r chi.Router) {
		r.Use(authMiddleware.Authenticate)
		r.Post("/invoice", invoiceHandler.Save)
		r.Get("/invoice", invoiceHandler.ListByAccount)
		r.Get("/invoice/{id}", invoiceHandler.GetById)
	})
}

func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}

	return s.server.ListenAndServe()
}
