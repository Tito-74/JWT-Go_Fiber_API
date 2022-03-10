package routes

import (
	"JWT-GoFiber/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14) 
	user := models.User{
		Name:			data["name"],
		Email : 	data["email"],
		Password: password,
	}
	return c.JSON(user)
}