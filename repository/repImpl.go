package repository

import (
	"context"

	models "github.com/Antoha2/auth"
)

type UserRep interface {
	SaveUser(ctx context.Context, email string, passHash []byte) (uid int64, err error)
	User(ctx context.Context, email string) (models.User, error)
}
