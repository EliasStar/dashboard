package screen

import (
	"context"
	"time"

	"github.com/EliasStar/Dashboard/DashD/command"
)

type Command struct {
	Action      Action
	Button      Button
	ToggleDelay time.Duration
}

func (c Command) IsValid(ctx context.Context) bool {
	w := c.Action.IsValid()
	x := c.Button.IsValid()
	y := 0 <= c.ToggleDelay
	z := c.ToggleDelay.Seconds() <= 5

	return w && x && y && z
}

func (c Command) Execute(ctx context.Context) command.Result {
	switch c.Action {
	case ActionRead:
		val, err := c.Button.Read()
		if err != nil {
			return command.ErrorRst(err.Error())
		}

		return Result(val)

	case ActionPress:
		c.Button.Write(true)

	case ActionRelease:
		c.Button.Write(false)

	case ActionToggle:
		c.Button.Toggle(c.ToggleDelay)
	}

	return command.OKRst{}
}
