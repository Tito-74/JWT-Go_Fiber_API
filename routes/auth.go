package routes

import (
	database "JWT-GoFiber/Database"
	"JWT-GoFiber/models"

	// "strconv"
	"fmt"
	"time"

	// "logger"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	
)



const SecretKey = "secret"

// type typetime time.Time

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}
	database.Database.Db.Create(&user)
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User

	database.Database.Db.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(404)
		fmt.Printf("Wrong user...\n")
		return c.JSON(fiber.Map{
			"Message": "Wrong Credentials",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(400)
		fmt.Printf("Wrong password...\n")
		return c.JSON(fiber.Map{
			"Message": "Wrong Credentials",
		})
		
	}


	
	token, ExpiresAt, err := CreateJwtToken(&user)
	if err != nil {
		return err

		
	}

	// c.Cookie(cookie)

	return c.JSON(fiber.Map{"token": token, "exp": ExpiresAt})
	
	

}

func CreateJwtToken(user *models.User) (string, int64, error) {
	ExpiresAt := time.Now().Add(time.Hour * 24).Unix()

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["user_id"] = user.Id

	claims["expires"] = ExpiresAt

	t, err := token.SignedString([]byte(SecretKey))

	if err != nil {

		return "", 0, err
	}

	return t, ExpiresAt, nil

}

// cookie = fiber.Cookie{
// 	Name:"jwt",
// 	Value: token, 
// 	Expires: time.Now().Add(time.Hour * 24).Unix(),
// 	HTTPOnly: true,
// }
// c.Cookie(cookie)




	




