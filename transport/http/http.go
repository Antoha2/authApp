package transport

import (
	"context"
	"log/slog"
	"net/http"

	authService "github.com/Antoha2/auth/services"
)

type webImpl struct {
	authService authService.AuthService
	server      *http.Server
	log         *slog.Logger
	port        int
}

func NewWeb(authService authService.AuthService, log *slog.Logger, port int) *webImpl {
	return &webImpl{
		authService: authService,
		log:         log,
		port:        port,
	}
}

func (wImpl *webImpl) Stop() {

	if err := wImpl.server.Shutdown(context.TODO()); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}
}
