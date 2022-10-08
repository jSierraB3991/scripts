package main

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/eliotandelon/gotesting/configuration"
	"gitlab.com/eliotandelon/gotesting/database"
	"gitlab.com/eliotandelon/gotesting/implementation"
	"gitlab.com/eliotandelon/gotesting/routes"
)

func handleHome(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "Hello world"})
}

func main() {
	e := echo.New()
	projectConfiguration, err := configuration.LoadYml("./config.yml")
	if err != nil {
		panic(err)
	}

	db := database.CreateConnection(context.Background(), projectConfiguration)
	mysqlRepository := database.NewRepository(db)
	bookService := implementation.NewBookServiceImplementation(mysqlRepository)

	e.GET("/", handleHome)
	routes.AddRoutes(e, bookService)
	e.Logger.Fatal(e.Start(":8081"))
}
