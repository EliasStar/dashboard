package main

import (
	"encoding/gob"
	"net"
	"os"

	"github.com/EliasStar/DashboardUtils/Commons/command"
	"github.com/EliasStar/DashboardUtils/Commons/command/display"
	"github.com/EliasStar/DashboardUtils/Commons/command/launch"
	"github.com/EliasStar/DashboardUtils/Commons/command/ledstrip"
	"github.com/EliasStar/DashboardUtils/Commons/command/schedule"
	"github.com/EliasStar/DashboardUtils/Commons/command/screen"
	"github.com/EliasStar/DashboardUtils/Commons/util"
	"github.com/EliasStar/DashboardUtils/Commons/util/misc"
)

func main() {
	con, err := net.Dial("tcp", os.Args[1]+":"+misc.DashDPort)
	util.FatalIfErr(err)

	defer con.Close()

	gob.Register(display.DisplayCmd{})
	gob.Register(launch.LaunchCmd{})
	gob.Register(ledstrip.LedstripCmd{})
	gob.Register(schedule.ScheduleCmd{})
	gob.Register(screen.ScreenCmd{})

	enc := gob.NewEncoder(con)

	cmd := command.Command(&screen.ScreenCmd{
		screen.ActionRead,
		screen.ButtonPower,
		0,
	})

	enc.Encode(&cmd)
	enc.Encode(&cmd)
}
