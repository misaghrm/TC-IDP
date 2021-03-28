package router

import (
	"github.com/gofiber/fiber/v2"
	"tc-micro-idp/handlers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/")
	api.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON("It Works!")
	})

	challenge := api.Group("/token")
	challenge.Post("/challenge", handlers.ChallengeToken)
	//challenge.Post("/verify")
	//challenge.Post("/refresh")
	//challenge.Get("/logout")
	challenge.Get("/test", handlers.TestToken)

}
