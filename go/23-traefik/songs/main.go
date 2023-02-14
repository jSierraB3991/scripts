package main

import "github.com/labstack/echo/v4"

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, []map[string]string{
			{"title": "In The End", "author": "Linkin Park"},
			{"title": "Dance 4 life", "author": "Tiesto"},
		})
	})
	e.Logger.Fatal(e.Start(":8080"))
}
