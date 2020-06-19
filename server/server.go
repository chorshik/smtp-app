package server

import (
	"context"
	"github.com/ebladrocher/smtp-client/server/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Server ...
type Server struct {
	router *mux.Router
	server  *http.Server
	handler *handlers.Handlers
}

// Start ...
func (s *Server) Start() error {
	s.server = &http.Server{
		Addr: "127.0.0.1:8080",
		Handler: s.router,
	}

	s.setHandlers()

	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			panic(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return s.server.Shutdown(ctx)

}

// NewServer ....
func NewServer() *Server {

	newServer := &Server{
		router: mux.NewRouter(),
		handler: handlers.NewHandlers(),
	}

	return newServer
}
