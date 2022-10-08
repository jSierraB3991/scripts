package routes

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/eliotandelon/gotesting/controllers"
	"gitlab.com/eliotandelon/gotesting/domain/book"
)

func AddRoutes(c *echo.Echo, service book.BookService) {
	api := controllers.NewApi(service)
	apiGroup := c.Group("/api")
	RegisterBookRoutes(apiGroup, api)
}
