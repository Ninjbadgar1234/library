package server

import (
	"log"
	"os"

	"example.com/book/src/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Init() {
	app := fiber.New()
	app.Use(recover.New())

	controller.SetCategoryControllers(app)
	controller.SetBookControllers(app)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
