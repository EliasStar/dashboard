package display

import (
	"context"
	"net/url"
	"os/exec"

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
	cmd := exec.Command("chromium-browser", d.URL)

	err := cmd.Start()

	return nil
}
