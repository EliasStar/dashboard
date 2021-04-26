package display

import (
	"context"
	"errors"
	"net/url"

	"github.com/EliasStar/DashboardUtils/Commons/command"
)

type DisplayCmd struct {
	URL string
}

func (d DisplayCmd) IsValid(ctx context.Context) bool {
	_, err := url.Parse(d.URL)
	return err == nil
}

func (d DisplayCmd) Execute(ctx context.Context) command.Result {
	cs, ok := ctx.Value("cmdstore").(command.CmdStore)
	if !ok {
		return DisplayRst{errors.New("commandstore not initialized")}
	}

	cmds := cs.Find("chromium-browser")

	for _, v := range cmds {
		cs.Kill(v.Process.Pid)
	}

	_, err := cs.Start("chromium-browser", d.URL)
	if err != nil {
		return DisplayRst{err}
	}

	return nil
}
