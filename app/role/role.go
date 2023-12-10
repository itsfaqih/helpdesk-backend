package role

import (
	"helpdesk/database"
	"helpdesk/utils"

	"github.com/gofiber/fiber/v2"
)

type Role struct {
	Code        string `json:"code" gorm:"primaryKey;type:varchar"`
	Name        string `json:"name" gorm:"not null;type:varchar"`
	Description string `json:"description" gorm:"not null;type:text;default:''"`
}

func Index(c *fiber.Ctx) error {
	var roles = []Role{}

	database.DBConn.Find(&roles)

	return c.JSON(utils.ApiResponseWithData[[]Role]{
		Data:    roles,
		Message: "Roles retrieved successfully",
	})
}
