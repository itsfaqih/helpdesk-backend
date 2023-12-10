package main

import (
	"helpdesk/app"
	"helpdesk/database"
	"helpdesk/database/migrations"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fiberApp := fiber.New()

	database.DBConn = database.Init()

	migrations.RunMigration()

	app.Routes(fiberApp)

	fiberApp.Listen(":3000")
}
