package launch

import (
	"context"
	"os/exec"

	"github.com/EliasStar/Dashboard/DashD/command"
)

type Command struct {
	Executable string
	Arguments  []string
}

func (c Command) IsValid(ctx context.Context) bool {
	_, err := exec.LookPath(c.Executable)
	return err == nil
}

func (c Command) Execute(ctx context.Context) command.Result {
	out, err := exec.Command(c.Executable, c.Arguments...).CombinedOutput()
	if err != nil {
		return command.ErrorRst(err.Error())
	}

	return Result(string(out))
}
