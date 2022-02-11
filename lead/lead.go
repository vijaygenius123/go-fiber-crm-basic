package lead

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"go-fiber-crm-basic/database"
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
