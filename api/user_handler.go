package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/svdunaev/hotel-reservation/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	user := types.User{
		FirstName: "John",
		LastName:  "Doe",
	}
	return c.JSON(user)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("John Doe")
}
