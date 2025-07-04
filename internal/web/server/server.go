package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/marcelochb/training-go-payment-gateway/internal/service"
	"github.com/marcelochb/training-go-payment-gateway/internal/web/handlers"
)

type Server struct {
	router         *chi.Mux
	server         *http.Server
	accountService *service.AccountService
	port           string
}

func NewServer(accountService *service.AccountService, port string) *Server {
	return &Server{
		router:         chi.NewRouter(),
		accountService: accountService,
		port:           port,
	}
}

func (server *Server) ConfigureRoutes() {
	accountHandler := handlers.NewAccountHandler(server.accountService)
	server.router.Post("/accounts", accountHandler.Create)
	server.router.Get("/accounts", accountHandler.Get)
}

func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}
	return s.server.ListenAndServe()
}
