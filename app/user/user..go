package user

import (
	"helpdesk/database"
	"helpdesk/utils"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Id         string    `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	FullName   string    `json:"full_name,omitempty" gorm:"not null;type:varchar"`
	Email      string    `json:"email,omitempty" gorm:"unique;not null;type:varchar"`
	Password   string    `json:"-" gorm:"not null;type:varchar"`
	IsActive   bool      `json:"is_active,omitempty"`
	IsArchived bool      `json:"is_archived,omitempty"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func Index(c *fiber.Ctx) error {
	var users = []User{}

	database.DBConn.Find(&users)

	return c.JSON(utils.ApiResponseWithData[[]User]{
		Data:    users,
		Message: "Users retrieved successfully",
	})
}

func Show(c *fiber.Ctx) error {
	id := c.Params("userId")

	var user User

	database.DBConn.First(&user, "id = ?", id)

	if user.Id == "" {
		return c.Status(404).JSON(utils.ApiResponseOnlyMessage{
			Message: "User not found",
		})
	}

	return c.JSON(utils.ApiResponseWithData[User]{
		Data:    user,
		Message: "User retrieved successfully",
	})
}

func Store(c *fiber.Ctx) error {
	var user User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(utils.ApiResponseOnlyMessage{
			Message: err.Error(),
		})
	}

	database.DBConn.Create(&user)

	c.Status(http.StatusCreated)

	return c.JSON(utils.ApiResponseWithData[User]{
		Data:    user,
		Message: "User created successfully",
	})
}

func Update(c *fiber.Ctx) error {
	id := c.Params("userId")

	var userToUpdate User

	database.DBConn.First(&userToUpdate, "id = ?", id)

	if userToUpdate.Id == "" {
		return c.Status(404).JSON(utils.ApiResponseOnlyMessage{
			Message: "User not found",
		})
	}

	var payload map[string]interface{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(utils.ApiResponseOnlyMessage{
			Message: err.Error(),
		})
	}

	database.DBConn.Model(&userToUpdate).Updates(payload)

	return c.JSON(utils.ApiResponseWithData[User]{
		Data:    userToUpdate,
		Message: "User updated successfully",
	})
}
