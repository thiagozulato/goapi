package httpresult

import "github.com/gofiber/fiber"

// ErrorResult define um modelo de retorno para StatusCode entre 4xx e 5xx
type ErrorResult struct {
	Message string `json:"message"`
}

// OkResult definie um modelo de retorno quando StatusCode 2xx
type OkResult struct {
	Data interface{} `json:"data"`
}

func newErrorResult(message string) ErrorResult {
	return ErrorResult{
		Message: message,
	}
}

func newOkResult(data interface{}) OkResult {
	return OkResult{
		Data: data,
	}
}

func setStatus(c *fiber.Ctx, statusCode int, response interface{}) {
	if err := c.Status(statusCode).JSON(response); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(newErrorResult("Erro ao gerar Result"))
	}
}

// InternalServerErrorMessage envia statusCode=500 com mensagem customizada
func InternalServerErrorMessage(c *fiber.Ctx, message string) {
	setStatus(c, fiber.StatusInternalServerError, newErrorResult(message))
}

// BadRequestMessage envia statusCode=400 com mensagem customizada
func BadRequestMessage(c *fiber.Ctx, message string) {
	setStatus(c, fiber.StatusBadRequest, newErrorResult(message))
}

// NotFoundMessage envia statusCode=404 com mensagem customizada
func NotFoundMessage(c *fiber.Ctx, message string) {
	setStatus(c, fiber.StatusNotFound, newErrorResult(message))
}

// NotFound envia statusCode=404 com mensagem padrão
func NotFound(c *fiber.Ctx) {
	setStatus(c, fiber.StatusNotAcceptable, newErrorResult("Not Found"))
}

// Unauthorized envia statusCode=401 com mensagem padrão
func Unauthorized(c *fiber.Ctx) {
	setStatus(c, fiber.StatusUnauthorized, newErrorResult("Unauthorized"))
}

// Ok envia statusCode=200
func Ok(c *fiber.Ctx, data interface{}) {
	setStatus(c, fiber.StatusOK, newOkResult(data))
}
