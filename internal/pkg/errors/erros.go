package errors

import "errors"

var (
	ErrInvalidEmailOrPassword   = errors.New("Email ou senha inválidos")
	ErrEmailAndPasswordRequired = errors.New("Email e senha são obrigatórios")
	ErrUserNotFound             = errors.New("Usuário não encontrado")
	ErrUserCreationFailed       = errors.New("Falha ao criar usuário")
)
