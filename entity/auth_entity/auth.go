package auth_entity

import "context"

type Auth interface {
	Register(ctx context.Context, input RegisterInput) (output RegisterOutput, err error)
}
