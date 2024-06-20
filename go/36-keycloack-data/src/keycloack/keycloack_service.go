package keycloack

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
	"github.com/jdsierrab3991/scripts/36-keycloack-data/src/model"
)

type KeycloackClient struct {
	kc *gocloak.GoCloak

	ClientId     string
	clientSecret string
	realm        string

	userAdmin  string
	pwdAdmin   string
	realmAdmin string
}

func NewKeycloackClient(env model.Enviroment) *KeycloackClient {
	return &KeycloackClient{
		ClientId:     env.ClientId,
		clientSecret: env.ClientSecret,
		realm:        env.Realm,

		userAdmin:  env.UserAdmin,
		pwdAdmin:   env.PwdAdmin,
		realmAdmin: env.RealmAdmin,

		kc: gocloak.NewClient(env.KeyCloakUrl),
	}
}

func (client *KeycloackClient) LoginUser(ctx context.Context, userName, password string) (*model.JWT, error) {

	jwt, err := client.kc.Login(ctx, client.ClientId, client.clientSecret, client.realm, userName, password)
	if err != nil {
		return nil, err
	}

	return &model.JWT{
		AccessToken:  jwt.AccessToken,
		RefreshToken: jwt.RefreshToken,
		ExpiredIn:    jwt.ExpiresIn,
	}, nil
}

func (client *KeycloackClient) CreateUser(ctx context.Context, userParam model.CreateUserParam) error {

	jwt, err := client.kc.LoginAdmin(ctx, client.userAdmin, client.pwdAdmin, client.realmAdmin)
	if err != nil {
		return err
	}

	kUser := gocloak.User{
		Username:  gocloak.StringP(userParam.UserName),
		Email:     gocloak.StringP(userParam.Email),
		FirstName: gocloak.StringP(userParam.FirstName),
		LastName:  gocloak.StringP(userParam.LastName),
		Enabled:   gocloak.BoolP(true),
	}

	userId, err := client.kc.CreateUser(ctx, jwt.AccessToken, client.realm, kUser)
	if err != nil {
		return err
	}

	return client.kc.SetPassword(ctx, jwt.AccessToken, userId, client.realm, userParam.Password, false)
}
