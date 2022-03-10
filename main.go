package main

import (
	"JWT-GoFiber/Database"
	"JWT-GoFiber/routes"
	"log"
	"github.com/gofiber/fiber/v2"
	
)

func RoutesSetup(app *fiber.App){
	app.Get("/", routes.Hello)
	app.Post("/api/register", routes.Register)
	
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	RoutesSetup(app)

	

	log.Fatal(app.Listen(":8000"))
}