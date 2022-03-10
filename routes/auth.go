package routes

import (
	"JWT-GoFiber/Database"
	"JWT-GoFiber/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go/v4"
	"strconv"
	"time"
)

const SecretKey = "secretkey"

//  Time := time.Date(2022, 03, 9, 7, 0, 0, 0, time.UTC)

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
	database.Database.Db.Create(&user)
	return c.JSON(user)
}

func Login(c *fiber.Ctx)error{
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User

	database.Database.Db.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"Message":"User not found",
		})
	}
	
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil{
		c.Status(400)
		return c.JSON(fiber.Map{
			"Message":"Wrong Password",
		})
		// .JSON(err.Error())
	}


	claims := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Time(),
		// Add(time.Hour * 24).Unix(), Bug to handle
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(500)
		c.JSON(fiber.Map{
			"Message":"Could not log in",
		})
	}

	return c.JSON(token)



}