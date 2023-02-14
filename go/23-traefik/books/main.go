package main

import "github.com/labstack/echo/v4"

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, []map[string]string{
			{"title": "100 años de soledad", "author": "Gabriel Garcia"},
			{"title": "Don Quijote", "author": "Miguel de Cervantes"},
		})
	})
	e.Logger.Fatal(e.Start(":8080"))
}
