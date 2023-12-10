package app

import (
	"helpdesk/app/role"
	"helpdesk/app/user"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	api := app.Group("/api")

	v1 := api.Group("/v1")

	users := v1.Group("/users")
	users.Get("/:userId", user.Show)
	users.Patch("/:userId", user.Update)
	users.Get("/", user.Index)
	users.Post("/", user.Store)

	roles := v1.Group("/roles")
	roles.Get("/", role.Index)
}
