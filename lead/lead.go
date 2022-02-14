package lead

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"go-fiber-crm-basic/database"
	"net/http"
)

type Lead struct {
	gorm.Model
	Name    string
	Company string
	Email   string
	phone   int
}

func GetLeads(c *fiber.Ctx) {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	err := c.JSON(leads)
	if err != nil {
		return
	}
}

func CreateLead(c *fiber.Ctx) {
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(http.StatusBadRequest).Send(err)
	}
	db.Create(&lead)
	err := c.JSON(lead)
	if err != nil {
		return
	}
}

func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	err := c.JSON(lead)
	if err != nil {
		return
	}
}

func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)

	if lead.Name == "" {
		c.Status(http.StatusNotFound).Send("No Lead Found")
	}

	db.Delete(&lead)

	c.Status(http.StatusOK).Send("Lead Deleted")
}
