package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
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
	v := validator.New()

	e.GET("/", hello)
	// GET  products?pageSize=4
	e.GET("/products", func(c echo.Context) error {
		fmt.Printf("pageSize: %s\n", c.QueryParam("pageSize"))
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
					break
				}
			}
		}

		if product == nil {
			return c.JSON(http.StatusNotFound, "Product Not Found")
		}
		return c.JSON(http.StatusOK, product)
	})

	e.POST("/products", func(c echo.Context) error {
		type Body struct {
			Name string `json:"product_name" validate:"required,min=4"`
			// Vendor string `json:"vendor" validate:"min=5,max=10"`
			//Email string `json:"email" validate:"required_with=vendor,email"`
			//Website string `json:"website" validate:"url"`
			//Country string `json:"country" validate:"len=2"`
			//DefaultDeviceIp string `json:"defaul_device_ip" validate:"ip"`
		}
		var reqBody Body
		if err := c.Bind(&reqBody); err != nil {
			return err
		}
		if err := v.Struct(reqBody); err != nil {
			return err
		}

		product := map[int]string{
			len(products) + 1: reqBody.Name,
		}
		products = append(products, product)
		return c.JSON(http.StatusOK, product)
	})

	e.Logger.Print(fmt.Sprintf("Listen on port %s", port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
}
