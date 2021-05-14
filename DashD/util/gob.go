package util

import (
	"encoding/gob"

	"github.com/EliasStar/Dashboard/DashD/command"
	"github.com/EliasStar/Dashboard/DashD/display"
	"github.com/EliasStar/Dashboard/DashD/launch"
	"github.com/EliasStar/Dashboard/DashD/ledstrip"
	"github.com/EliasStar/Dashboard/DashD/schedule"
	"github.com/EliasStar/Dashboard/DashD/screen"
)

func InitGOBBasic() {
	gob.Register(command.ErrorRst(""))
	gob.Register(command.OKRst{})
}

func InitGOBLedstrip() {
	gob.Register(ledstrip.RGB{})
	gob.Register(ledstrip.RGBA32{})

	gob.Register(ledstrip.Command{})
	gob.Register(ledstrip.Result{})
}

func InitGOBScreen() {
	gob.Register(screen.Command{})
	gob.Register(screen.Result(false))
}

func InitGOBFull() {
	InitGOBBasic()
	InitGOBLedstrip()
	InitGOBScreen()

	gob.Register(display.Command{})

	gob.Register(launch.Command{})
	gob.Register(launch.Result(""))

	gob.Register(schedule.Command{})
}
