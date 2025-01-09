package handlers

import (
	"fmt"
	"roma-x-api/engine/game"
	"roma-x-api/engine/service"

	"github.com/gofiber/fiber/v2"
)

func HandleSpin(symbols map[string]game.Symbol) fiber.Handler {
	gameService := service.NewGameService(symbols)

	return func(c *fiber.Ctx) error {
		var req game.SpinRequest

		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		if err := validateBet(req.Bet); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		response := gameService.GenerateGameOutcome(req.Bet)
		return c.JSON(response)
	}
}

func validateBet(bet float64) error {
	if bet < game.MinBet || bet > game.MaxBet {
		return fmt.Errorf("bet must be between %.2f and %.2f", game.MinBet, float64(game.MaxBet))
	}
	return nil
}
