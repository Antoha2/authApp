package transport

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Antoha2/auth/config"
	authEndpoints "github.com/Antoha2/auth/transport/endpoints"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func (wImpl *webImpl) StartHTTP() error {
	AuthOptions := []httptransport.ServerOption{
		//httptransport.ServerBefore(wImpl.UserIdentify),
	}

	signInHandler := httptransport.NewServer(
		authEndpoints.MakeLoginEndpoint(wImpl.authService),
		decodeMakeLoginRequest,
		encodeResponse,
		AuthOptions...,
	)

	signUpHandler := httptransport.NewServer(
		authEndpoints.MakeRegisterEndpoint(wImpl.authService),
		decodeMakeRegisterRequest,
		encodeResponse,
		AuthOptions...,
	)

	r := mux.NewRouter()

	r.Methods("POST").Path("/auth/sign-in").Handler(signInHandler)
	r.Methods("POST").Path("/auth/sign-up").Handler(signUpHandler)

	wImpl.server = &http.Server{Addr: config.HTTPAddr}
	log.Printf("(auth) Запуск HTTP-сервера на http://127.0.0.1%s\n", wImpl.server.Addr) //:8180

	if err := http.ListenAndServe(wImpl.server.Addr, r); err != nil {
		log.Println(err)
	}

	return nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func decodeMakeLoginRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request authEndpoints.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeMakeRegisterRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request authEndpoints.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
