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
	c.JSON(leads)
}

func CreateLead(c *fiber.Ctx) {
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(http.StatusBadRequest).Send(err)
	}
	db.Create(&lead)
	c.JSON(lead)
}
