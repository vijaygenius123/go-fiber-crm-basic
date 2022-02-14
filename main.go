package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"go-fiber-crm-basic/database"
	"go-fiber-crm-basic/lead"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/leads", lead.GetLeads)
	app.Get("/api/leads/:id", lead.GetLead)
	app.Post("/api/leads", lead.CreateLead)
	app.Delete("/api/leads/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("Failed to connect to DB")
	}
	fmt.Println("Connected To DB")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("DB Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)

	err := app.Listen(8000)
	if err != nil {
		panic("Couldn't start server")
	}
	defer func(DBConn *gorm.DB) {
		err := DBConn.Close()
		if err != nil {
			panic("Error closing DB connection")
		}
	}(database.DBConn)
}
