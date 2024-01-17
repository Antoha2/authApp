package transport

import (
	"context"
	"net/http"

	authService "github.com/Antoha2/auth/services"
)

type webImpl struct {
	authService authService.AuthService
	server      *http.Server
}

func NewWeb(authService authService.AuthService) *webImpl {
	return &webImpl{
		authService: authService,
	}
}

func (wImpl *webImpl) Stop() {

	if err := wImpl.server.Shutdown(context.TODO()); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}
}
