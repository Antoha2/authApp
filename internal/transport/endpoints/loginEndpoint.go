package endpoints

import (
	"context"
	"log"

	authservice "github.com/Antoha2/auth/internal/services"

	"github.com/go-kit/kit/endpoint"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Roles    []string
}

type LoginResponse struct {
	Token string `json:"token"`
}

func MakeLoginEndpoint(s authservice.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		//req := request.(LoginRequest)

		_ = request.(LoginRequest)
		token := "!!!"
		log.Println(token)
		// token, err := s.GenerateToken(req.Username, req.Password)
		// if err != nil {
		// 	log.Println(err)
		// 	return nil, err
		// }

		return LoginResponse{Token: token}, nil
	}
}
