package usuario

// Usuario model
type Usuario struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Senha string `json:"senha"`
	Salt  string `json:"salt"`
	Ativo bool   `json:"ativo"`
}
