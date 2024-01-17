package endpoints

import (
	"context"

	authservice "github.com/Antoha2/auth/services"
	"github.com/go-kit/kit/endpoint"
)

type RegisterRequest struct {
	FirstName string `json:"firstname" gorm:"column:firstname"`
	LastName  string `json:"lastname" gorm:"column:lastname"`
	Username  string `json:"username" gorm:"column:username"`
	Password  string `json:"password"`
}

type RegisterResponse struct {
	UserId int `json:"user_id"`
}

// const (
// 	roleAdmin = "admin"
// 	roleUser  = "user"
// 	roleDev   = "dev"
// )

func MakeRegisterEndpoint(s authservice.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		_ = request.(RegisterRequest)
		// req := request.(RegisterRequest)

		// inputUser := new(helper.User)
		// inputRoles := new(helper.UsersRoles)

		// inputUser.FirstName = req.FirstName
		// inputUser.LastName = req.LastName
		// inputUser.Password = req.Password
		// inputUser.Username = req.Username

		// inputRoles.Roles = append(inputRoles.Roles, roleAdmin)
		// inputRoles.Roles = append(inputRoles.Roles, roleDev) // ?!? !!!!!!!!!!!!!!!!!!!!!!!!!!

		// userId, err := s.CreateUser(inputUser, inputRoles)

		// if err != nil {
		// 	return nil, err
		// }

		userId := 1122

		return RegisterResponse{UserId: userId}, nil
	}
}
