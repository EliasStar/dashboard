package display

import (
	"context"
	"errors"
	"net/url"
	"os/exec"

	"github.com/EliasStar/DashboardUtils/Commons/command"
	"github.com/EliasStar/DashboardUtils/Commons/util/misc"
)

type DisplayCmd struct {
	// TODO: Add url getter and remover
	URL string
}

func (d DisplayCmd) IsValid(ctx context.Context) bool {
	_, err := url.Parse(d.URL)
	return err == nil
}

func (d DisplayCmd) Execute(ctx context.Context) command.Result {
	cmd, ok := ctx.Value(misc.DisplayContextKey).(*exec.Cmd)
	if !ok {
		return command.ErrorRst{errors.New("display not initialized")}
	}

	cmd.Process.Kill()
	cmd.Process.Release()

	cmd = exec.Command(misc.DashDBrowser, d.URL)
	cmd.Start()

	return nil
}
