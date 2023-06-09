package auth_entity

import "context"

//go:generate mockery --name Jwt
type Jwt interface {
	GenerateToken(claims interface{}) (token string, err error)
	ParseToken(tokenString string) (claims interface{}, err error)
}

//go:generate mockery --name AuthUseCase
type AuthUseCase interface {
	Register(ctx context.Context, input RegisterInput) (output RegisterOutput, err error)
	Login(ctx context.Context, input LoginInput) (output LoginOutput, err error)
}
