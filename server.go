package main

import (
	"github.com/gofiber/fiber"
	config "github.com/thiagozulato/goapi/Config"
)

func main() {
	app := fiber.New()

	app.Get("/api/v1", func(ctx *fiber.Ctx) {
		ctx.JSON(&fiber.Map{
			"sucesso": true,
		})
	})

	port := config.Config("PORT")

	if err := app.Listen(port); err != nil {
		panic(err)
	}
}
