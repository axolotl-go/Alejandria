package main

import (
	"github.com/axolotl-go/Alejandria/db"
	"github.com/axolotl-go/Alejandria/models"
	"github.com/axolotl-go/Alejandria/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	db.Dbconnection()
	db.DB.AutoMigrate(models.Users{})

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Worlds!")

	})

	//Users
	app.Post("/user", routes.CreateUser)
	app.Get("/users", routes.FindUsers)

	app.Listen(":8080")
}
