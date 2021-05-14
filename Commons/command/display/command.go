package display

import (
	"context"
	"net/url"
	"os/exec"

	cmd "github.com/EliasStar/DashboardUtils/Commons/command"
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

func (d DisplayCmd) Execute(ctx context.Context) cmd.Result {
	c, ok := ctx.Value(misc.DisplayContextKey).(*exec.Cmd)
	if !ok {
		return cmd.ErrorRst("display not initialized")
	}

	if d.Action == ActionGet {
		return DisplayRst(c.Args[1])
	}

	if c.Process != nil {
		c.Process.Kill()
		c.Process.Release()
	}

	if d.Action == ActionSet {
		c = exec.Command(misc.DashDBrowser, d.URL)
		c.Start()
	}

	return cmd.OKRst{}
}
