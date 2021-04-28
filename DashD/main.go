package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os/exec"

	"github.com/EliasStar/DashboardUtils/Commons/command"
	"github.com/EliasStar/DashboardUtils/Commons/command/display"
	"github.com/EliasStar/DashboardUtils/Commons/command/launch"
	"github.com/EliasStar/DashboardUtils/Commons/command/ledstrip"
	"github.com/EliasStar/DashboardUtils/Commons/command/schedule"
	"github.com/EliasStar/DashboardUtils/Commons/command/screen"
	hw "github.com/EliasStar/DashboardUtils/Commons/hardware"
	"github.com/EliasStar/DashboardUtils/Commons/util"
	"github.com/EliasStar/DashboardUtils/Commons/util/misc"
)

func main() {
	strip, err := hw.NewLedstrip(misc.LedstripDataPin, misc.LedstripLength, misc.LedstripHasBurnerLED)
	util.FatalIfErr(err)

	util.FatalIfErr(strip.Init())
	defer strip.Fini()

	cmd := exec.Command(misc.DashDBrowser)

	ctx := context.Background()
	ctx = context.WithValue(ctx, misc.LedstripContextKey, strip)
	ctx = context.WithValue(ctx, misc.DisplayContextKey, cmd)

	gob.Register(display.DisplayCmd{})
	gob.Register(launch.LaunchCmd{})
	gob.Register(ledstrip.LedstripCmd{})
	gob.Register(schedule.ScheduleCmd{})
	gob.Register(screen.ScreenCmd{})

	listener, err := net.Listen("tcp", "127.0.0.1:"+misc.DashDPort)
	util.FatalIfErr(err)

	defer listener.Close()
	fmt.Println("Listening on:", listener.Addr())

	for {
		con, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println("New Connection:", con.RemoteAddr())
		go handleConnection(con, ctx)
	}
}

func handleConnection(con net.Conn, ctx context.Context) {
	dec := gob.NewDecoder(con)

	var cmd command.Command
	for dec.Decode(&cmd) == nil {
		fmt.Printf("(%v, %T)\n", cmd, cmd)
	}
}
