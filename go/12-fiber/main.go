package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jsierrab3991/scripts/12-fiber/database"
	"github.com/jsierrab3991/scripts/12-fiber/lead"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func main() {
	app := fiber.New()

	initDatabase()
	setupRoutes(app)

	app.Listen(":3000")
}

func initDatabase() {
	database.DatabaseConnect()
	database.DbConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrate")
}
