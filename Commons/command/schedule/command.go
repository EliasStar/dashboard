package schedule

import (
	"context"
	"strings"

	"github.com/EliasStar/DashboardUtils/Commons/command"
	"github.com/EliasStar/DashboardUtils/Commons/command/launch"
)

type ScheduleCmd struct {
	Action         ScheduleAction
	CronExpression string
	Command        launch.LaunchCmd
}

func (s ScheduleCmd) IsValid(ctx context.Context) bool {
	a := s.Action.IsValid()
	b := s.Command.IsValid(ctx)
	c := len(strings.Fields(s.CronExpression)) == 5

	if s.Action != ActionWrite {
		c = c || s.CronExpression == ""
	}

	return a && b && c

}

func (s ScheduleCmd) Execute(ctx context.Context) command.Result {
	// TODO
	return nil
}
