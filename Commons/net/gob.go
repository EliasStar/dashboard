package net

import (
	"encoding/gob"

	"github.com/EliasStar/DashboardUtils/Commons/command"
	"github.com/EliasStar/DashboardUtils/Commons/command/display"
	"github.com/EliasStar/DashboardUtils/Commons/command/launch"
	"github.com/EliasStar/DashboardUtils/Commons/command/ledstrip"
	"github.com/EliasStar/DashboardUtils/Commons/command/schedule"
	"github.com/EliasStar/DashboardUtils/Commons/command/screen"
	"github.com/EliasStar/DashboardUtils/Commons/util/color"
)

func InitGOBBasic() {
	gob.Register(command.ErrorRst(""))
	gob.Register(command.OKRst{})
}

func InitGOBLedstrip() {
	gob.Register(color.RGB{})
	gob.Register(color.RGBA32{})

	gob.Register(ledstrip.LedstripCmd{})
	gob.Register(ledstrip.LedstripRst{})
}

func InitGOBScreen() {
	gob.Register(screen.ScreenCmd{})
	gob.Register(screen.ScreenRst(false))
}

func InitGOBFull() {
	InitGOBBasic()
	InitGOBLedstrip()
	InitGOBScreen()

	gob.Register(display.DisplayCmd{})

	gob.Register(launch.LaunchCmd{})
	gob.Register(launch.LaunchRst(""))

	gob.Register(schedule.ScheduleCmd{})
}
