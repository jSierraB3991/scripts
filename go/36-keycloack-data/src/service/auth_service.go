package service

import (
	"context"

	"github.com/jdsierrab3991/scripts/36-keycloack-data/src/keycloack"
	"github.com/jdsierrab3991/scripts/36-keycloack-data/src/model"
)

type AuthService struct {
	kCLient keycloack.KeycloackClient
}

func NewAuthService(env model.Enviroment) *AuthService {
	return &AuthService{
		kCLient: *keycloack.NewKeycloackClient(env),
	}
}

func (service *AuthService) SaveUser(ctx context.Context, userName, password, email string) error {
	return service.kCLient.CreateUser(ctx, model.CreateUserParam{
		UserName:  userName,
		Password:  password,
		Email:     email,
		FirstName: "Juancho",
		LastName:  "Rois",
	})
}

func (service *AuthService) Login(ctx context.Context, userName, password string) (*model.JWT, error) {
	return service.kCLient.LoginUser(ctx, userName, password)
}
