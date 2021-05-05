package command

import "context"

type Command interface {
	IsValid(ctx context.Context) bool
	Execute(ctx context.Context) Result
}
