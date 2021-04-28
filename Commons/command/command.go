package command

import "context"

type ContextKey uint

type Command interface {
	IsValid(ctx context.Context) bool
	Execute(ctx context.Context) Result
}
