package errors

import "errors"

var (
	ErrInvalidEmailOrPassword   = errors.New("Email ou senha inválidos")
	ErrEmailAndPasswordRequired = errors.New("Email e senha são obrigatórios")
)
