package repository

import (
	"context"

	models "github.com/Antoha2/auth"
	"gorm.io/gorm"
)

// var (
// 	ErrUserExists   = errors.New("user already exists")
// 	ErrUserNotFound = errors.New("user not found")
// 	ErrAppNotFound  = errors.New("app not found")
// )

type AuthRepository interface {
	UserSaver(ctx context.Context, email string, passHash []byte) (uid int64, err error)
	UserProvider(ctx context.Context, email string) (models.User, error)
	App(ctx context.Context, appID int) (models.App, error)
}

type RepAuth struct {
	DB *gorm.DB
	AuthRepository
}

func NewRepAuth(dbx *gorm.DB) *RepAuth {
	return &RepAuth{

		DB: dbx,
	}
}
