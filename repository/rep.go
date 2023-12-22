package repository

import (
	"context"
	"errors"
	"log/slog"
	"time"

	models "github.com/Antoha2/auth"
)

var (
	ErrUserExists   = errors.New("user already exists")
	ErrUserNotFound = errors.New("user not found")
	ErrAppNotFound  = errors.New("app not found")
)

type AuthRepository interface {
	UserSaver(ctx context.Context, email string, passHash []byte) (uid int64, err error)
	UserProvider(ctx context.Context, email string) (models.User, error)
	App(ctx context.Context, appID int) (models.App, error)
}

type RepAuth struct {
	*slog.Logger
	AuthRepository
	TokenTTL time.Duration
}

func New(log *slog.Logger, rep AuthRepository, tokenTTL time.Duration) *RepAuth {
	return &RepAuth{Logger: log,
		AuthRepository: rep,
		TokenTTL:       tokenTTL}
}
