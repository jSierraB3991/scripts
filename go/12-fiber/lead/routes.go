package lead

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jsierrab3991/scripts/12-fiber/database"
)

func GetLeads(c *fiber.Ctx) error {
	db := database.DbConn
	var leads []Lead
	db.Find(&leads)
	return c.JSON(leads)
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DbConn
	var lead Lead
	db.Find(&lead, id)
	if lead.Name == "" {
		return c.Status(404).Send([]byte("Not lead Founbd with id"))
	}
	return c.JSON(lead)
}

func NewLead(c *fiber.Ctx) error {
	db := database.DbConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		return c.Status(500).Send([]byte("The body unwrapper to lead"))
	}
	db.Create(&lead)
	return c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DbConn
	var lead Lead
	db.Find(&lead, id)
	if lead.Name == "" {
		return c.Status(404).Send([]byte("Not lead Founbd with id"))
	}
	db.Delete(&lead)
	return c.Status(204).Send([]byte("Lead sucessfull delete"))
}
