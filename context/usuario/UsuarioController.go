package usuario

import (
	"fmt"
	"goapi/helpers/httpresult"
	"strconv"

	"github.com/gofiber/fiber"
)

// ObterUsuario obtem todos os usuarios
func ObterUsuario(c *fiber.Ctx) {
	usuarios := []Usuario{
		{
			ID:    1,
			Ativo: true,
			Email: "usuario1@gmail.com",
			Salt:  "@123456",
			Senha: "teste",
		},
		{
			ID:    2,
			Ativo: true,
			Email: "usuario2@gmail.com",
			Salt:  "@654321",
			Senha: "143@tre#44",
		},
	}

	httpresult.Ok(c, usuarios)
}

// CriarUsuario cria um novo usuário
func CriarUsuario(c *fiber.Ctx) {
	usuario := Usuario{}

	if err := c.BodyParser(&usuario); err != nil {
		httpresult.BadRequestMessage(c, "Parâmetros inválidos")
		return
	}

	httpresult.Ok(c, usuario)
}

// ObterUsuarioPorID recupera usuário pelo Identificador
func ObterUsuarioPorID(c *fiber.Ctx) {
	id, _ := strconv.Atoi(c.Params("id"))

	fmt.Printf("ID recuperado %v\n", id)

	httpresult.Ok(c, Usuario{})
}

// ExcluirUsuario utilizando o identificador
func ExcluirUsuario(c *fiber.Ctx) {
	id, _ := strconv.Atoi(c.Params("id"))

	fmt.Printf("ID recuperado %v\n", id)

	httpresult.Ok(c, Usuario{})
}

// AtualizarUsuario .
func AtualizarUsuario(c *fiber.Ctx) {
	id, _ := strconv.Atoi(c.Params("id"))

	fmt.Printf("ID recuperado %v\n", id)

	usuario := Usuario{}

	if err := c.BodyParser(&usuario); err != nil {
		httpresult.BadRequestMessage(c, "Corpo da mensagem inválido")
		return
	}

	if id != usuario.ID {
		httpresult.BadRequestMessage(c, "Não é possível alterar os dados desse usuário")
		return
	}

	httpresult.Ok(c, usuario)
}
