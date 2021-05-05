package display

import (
	"context"
	"net/url"
	"os/exec"

	. "github.com/EliasStar/DashboardUtils/Commons/command"
	"github.com/EliasStar/DashboardUtils/Commons/util/misc"
)

type DisplayCmd struct {
	Action DisplayAction
	URL    string
}

func (d DisplayCmd) IsValid(ctx context.Context) bool {
	valid := d.Action.IsValid()
	_, err := url.Parse(d.URL)
	return valid && err == nil
}

func (d DisplayCmd) Execute(ctx context.Context) Result {
	cmd, ok := ctx.Value(misc.DisplayContextKey).(*exec.Cmd)
	if !ok {
		return NewErrorRstFromString("display not initialized")
	}

	if d.Action == ActionGet {
		return DisplayRst(cmd.Args[1])
	}

	cmd.Process.Kill()
	cmd.Process.Release()

	if d.Action == ActionSet {
		cmd = exec.Command(misc.DashDBrowser, d.URL)
		cmd.Start()
	}

	return OKRst{}
}
