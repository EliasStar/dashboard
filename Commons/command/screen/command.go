package screen

import (
	"context"
	"time"

	cmd "github.com/EliasStar/DashboardUtils/Commons/command"
)

type ScreenCmd struct {
	Action      ScreenAction
	Button      ScreenButton
	ToggleDelay time.Duration
}

func (s ScreenCmd) IsValid(ctx context.Context) bool {
	a := s.Action.IsValid()
	b := s.Button.IsValid()
	c := 0 <= s.ToggleDelay
	d := s.ToggleDelay.Seconds() <= 5

	return a && b && c && d
}

func (s ScreenCmd) Execute(ctx context.Context) cmd.Result {
	switch s.Action {
	case ActionRead:
		val, err := s.Button.Pin().Read()
		if err != nil {
			return cmd.ErrorRst(err.Error())
		}

		return ScreenRst(val)

	case ActionPress:
		s.Button.Pin().Write(true)

	case ActionRelease:
		s.Button.Pin().Write(false)

	case ActionToggle:
		s.Button.Pin().Toggle(s.ToggleDelay)
	}

	return cmd.OKRst{}
}
