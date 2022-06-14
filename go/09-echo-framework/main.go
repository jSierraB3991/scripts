package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hola mundo!")
}

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	products := []map[int]string{{1: "Moviles"}, {2: "Tv"}, {3: "Laptops"}}
	e := echo.New()

	e.GET("/", hello)
	e.GET("/products", func(c echo.Context) error {
		return c.JSON(http.StatusOK, products)
	})
	e.GET("/products/:id", func(c echo.Context) error {
		var product map[int]string
		pId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}

		for _, p := range products {
			for k := range p {
				if pId == k {
					product = p
				}
			}
		}

		if product == nil {
			return c.JSON(http.StatusNotFound, "Product Not Found")
		}
		return c.JSON(http.StatusOK, product)
	})

	e.Logger.Print(fmt.Sprintf("Listen on port %s", port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
}
