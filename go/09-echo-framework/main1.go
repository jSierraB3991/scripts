package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hola mundo!")
	})
	e.Logger.Print("Listeng on port 1323")
	e.Logger.Fatal(e.Start(":1323"))
}
