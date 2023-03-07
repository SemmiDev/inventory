package main

import (
	"inventory/handler"
	"inventory/repository"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	itemRepository := repository.NewItemRepository()
	itemHandler := handler.NewItemHandler(itemRepository, app)
	itemHandler.SetupRoutes()

	log.Println("Server is running on port 8080")
	app.Listen(":8080")
}
