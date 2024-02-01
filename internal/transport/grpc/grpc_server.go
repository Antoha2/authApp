package transport

// import (
// 	"context"
// 	"errors"

// 	ssov1 "github.com/Antoha2/auth/proto/gen/go/sso"
// 	AuthRepository "github.com/Antoha2/auth/repository"
// 	AuthService "github.com/Antoha2/auth/services"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"
// )

// func Register(gRPCServer *grpc.Server, auth AuthService.AuthService) {
// 	ssov1.RegisterAuthServer(gRPCServer, &serverAPI{service: auth})
// }

// func (s *serverAPI) Login(ctx context.Context, in *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {

// 	if in.Email == "" {
// 		return nil, status.Error(codes.InvalidArgument, "email is required")
// 	}

// 	if in.Password == "" {
// 		return nil, status.Error(codes.InvalidArgument, "password is required")
// 	}

// 	if in.GetAppId() == 0 {
// 		return nil, status.Error(codes.InvalidArgument, "app_id is required")
// 	}

// 	token, err := s.service.Login(ctx, in.GetEmail(), in.GetPassword(), int(in.GetAppId()))
// 	if err != nil {
// 		// Ошибку auth.ErrInvalidCredentials мы создадим ниже
// 		if errors.Is(err, AuthService.ErrInvalidCredentials) {
// 			return nil, status.Error(codes.InvalidArgument, "invalid email or password")
// 		}

// 		return nil, status.Error(codes.Internal, "failed to login")
// 	}

// 	return &ssov1.LoginResponse{Token: token}, nil

// }

// func (s *serverAPI) Register(ctx context.Context, in *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {
// 	if in.Email == "" {
// 		return nil, status.Error(codes.InvalidArgument, "email is required")
// 	}

// 	if in.Password == "" {
// 		return nil, status.Error(codes.InvalidArgument, "password is required")
// 	}

// 	uid, err := s.service.RegisterNewUser(ctx, in.Email, in.Password)
// 	if err != nil {
// 		// Ошибку storage.ErrUserExists мы создадим ниже
// 		if errors.Is(err, AuthRepository.ErrUserExists) {
// 			return nil, status.Error(codes.AlreadyExists, "user already exists")
// 		}
// 		return nil, status.Error(codes.AlreadyExists, "user already exists")
// 	}
// 	return &ssov1.RegisterResponse{UserId: uid}, nil
// }
