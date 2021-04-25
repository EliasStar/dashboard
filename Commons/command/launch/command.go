package launch

import (
	"context"

	"github.com/EliasStar/DashboardUtils/Commons/command"
)

type LaunchCmd struct {
	Executable string
	Arguments  []string
}

func (l LaunchCmd) IsValid(ctx context.Context) bool {

	return false
}

func (l LaunchCmd) Execute(ctx context.Context) command.Result {

	return nil
}
