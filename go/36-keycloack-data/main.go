package main

import (
	"context"
	"log"
	"os"

	"github.com/jdsierrab3991/scripts/36-keycloack-data/src/model"
	"github.com/jdsierrab3991/scripts/36-keycloack-data/src/service"
	"github.com/joho/godotenv"
)

func main() {
	log.Println("Start")
	godotenv.Load()
	authService := service.NewAuthService(getEnviroment())
	ctx := context.Background()
	userName := "Trolencio"
	password := "123456789"
	email := "judas3991@gmail.com"
	err := authService.SaveUser(ctx, userName, password, email)
	if err != nil {
		log.Fatal(err)
	}

	jwt, err := authService.Login(ctx, userName, password)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(jwt)
}

func getEnviroment() model.Enviroment {
	return model.Enviroment{
		ClientId:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Realm:        os.Getenv("REALM"),
		UserAdmin:    os.Getenv("USER_ADMIN"),
		PwdAdmin:     os.Getenv("PWD_ADMIN"),
		RealmAdmin:   os.Getenv("REALM_ADMIN"),
		KeyCloakUrl:  os.Getenv("KEYCLOAK_URL"),
	}
}
