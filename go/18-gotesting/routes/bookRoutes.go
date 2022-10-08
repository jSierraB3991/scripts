package routes

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/eliotandelon/gotesting/controllers"
)

func RegisterBookRoutes(g *echo.Group, api *controllers.API) {
	bookGroup := g.Group("/book")
	bookGroup.GET("", api.ListBooks)
	bookGroup.POST("", api.SaveBook)
	bookGroup.GET("/:id", api.GetBookById)
}
