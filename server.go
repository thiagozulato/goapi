package main

import (
	"goapi/config"
	"goapi/context/usuario"
	"goapi/helpers/httpresult"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	api := app.Group("/api")
	v1 := api.Group("/v1")

	usuario.Rotas(v1)

	app.Use(httpresult.NotFound)

	port := config.Get("SERVER_PORT")

	if err := app.Listen(port); err != nil {
		panic(err)
	}
}
