package usuario

import "github.com/gofiber/fiber"

// Rotas define aas rotas do recurso Usuario
func Rotas(c *fiber.Group) {
	c.Get("/usuarios", ObterUsuario)
	c.Get("/usuarios/:id", ObterUsuarioPorID)
	c.Post("/usuarios", CriarUsuario)
	c.Put("/usuarios/:id", AtualizarUsuario)
	c.Delete("/usuarios/:id", ExcluirUsuario)
}
