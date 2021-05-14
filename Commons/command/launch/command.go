package launch

import (
	"context"
	"os/exec"

	cmd "github.com/EliasStar/DashboardUtils/Commons/command"
)

type LaunchCmd struct {
	Executable string
	Arguments  []string
}

func (l LaunchCmd) IsValid(ctx context.Context) bool {
	_, err := exec.LookPath(l.Executable)
	return err == nil
}

func (l LaunchCmd) Execute(ctx context.Context) cmd.Result {
	out, err := exec.Command(l.Executable, l.Arguments...).CombinedOutput()
	if err != nil {
		return cmd.ErrorRst(err.Error())
	}

	return LaunchRst(string(out))
}
