package transport

import (
	"net/http"

	authService "github.com/antoha2/auth/service"
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
