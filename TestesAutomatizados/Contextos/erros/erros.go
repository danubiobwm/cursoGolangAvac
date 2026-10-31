package erros

import "errors"

var (
	ErroUserIDObrigatorio   = errors.New("user id é obrigatório no contexto")
	ErroDeadlineObrigatorio = errors.New("deadline é obrigatório no contexto")
)
