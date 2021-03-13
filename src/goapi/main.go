package main

import (
	"goplusreact/src/goapi/database"
	"goplusreact/src/goapi/routes"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {

	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")

}
