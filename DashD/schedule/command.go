package schedule

import (
	"context"
	"strings"

	"github.com/EliasStar/Dashboard/DashD/command"
	"github.com/EliasStar/Dashboard/DashD/launch"
)

type Command struct {
	Action         Action
	CronExpression string
	Command        launch.Command
}

func (c Command) IsValid(ctx context.Context) bool {
	x := c.Action.IsValid()
	y := c.Command.IsValid(ctx)
	z := len(strings.Fields(c.CronExpression)) == 5

	if c.Action != ActionWrite {
		z = z || c.CronExpression == ""
	}

	return x && y && z

}

func (c Command) Execute(ctx context.Context) command.Result {
	// TODO ScheduleCmd Execute
	return nil
}
