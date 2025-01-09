package main

import (
	"roma-x-api/engine/game"
	"roma-x-api/engine/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: customErrorHandler,
	})

	// Add middleware
	app.Use(logger.New())

	// Initialize game symbols
	symbols := game.InitializeSymbols()

	// Setup routes
	setupRoutes(app, symbols)

	app.Listen(":3000")
}

func setupRoutes(app *fiber.App, symbols map[string]game.Symbol) {
	app.Post("/spin", handlers.HandleSpin(symbols))
}

func customErrorHandler(c *fiber.Ctx, err error) error {
	// Return error as JSON
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": err.Error(),
	})
}
