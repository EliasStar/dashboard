package display

import (
	"context"
	"net/url"
	"os/exec"

	"github.com/EliasStar/Dashboard/DashD/command"
)

const Browser = "browser"

type ContextKey struct{}

type Command struct {
	Action Action
	URL    string
}

func (c Command) IsValid(ctx context.Context) bool {
	valid := c.Action.IsValid()
	_, err := url.Parse(c.URL)
	return valid && err == nil
}

func (c Command) Execute(ctx context.Context) command.Result {
	cmd, ok := ctx.Value(ContextKey{}).(*exec.Cmd)
	if !ok {
		return command.ErrorRst("display not initialized")
	}

	if c.Action == ActionGet {
		return Result(cmd.Args[1])
	}

	if cmd.Process != nil {
		cmd.Process.Kill()
		cmd.Process.Release()
	}

	if c.Action == ActionSet {
		*cmd = *exec.Command(Browser, c.URL)
		cmd.Start()
	}

	return command.OKRst{}
}
