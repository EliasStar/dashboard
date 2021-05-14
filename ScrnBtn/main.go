package main

import (
	"context"
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/EliasStar/DashboardUtils/Commons/command"
	sn "github.com/EliasStar/DashboardUtils/Commons/command/screen"
	nt "github.com/EliasStar/DashboardUtils/Commons/net"
	"github.com/EliasStar/DashboardUtils/Commons/util"
	"github.com/EliasStar/DashboardUtils/Commons/util/misc"
)

func main() {
	cmd := parseCommand()
	var rst command.Result

	con, conErr := net.Dial("tcp", "127.0.0.1:"+misc.DashDPort)
	if conErr == nil {
		defer con.Close()

		nt.InitGOBBasic()
		nt.InitGOBScreen()

		util.PanicIfErr(gob.NewEncoder(con).Encode(&cmd))
		util.PanicIfErr(gob.NewDecoder(con).Decode(&rst))
	} else {
		for _, b := range sn.ScreenButtons() {
			util.PanicIfErr(b.Pin().Mode(true))
		}

		rst = cmd.Execute(context.Background())
	}

	if !rst.IsOK() {
		log.Panic(rst)
	}

	snRst, ok := rst.(sn.ScreenRst)
	if ok {
		fmt.Println(snRst)
	}
}

func parseCommand() (cmd command.Command) {
	if len(os.Args) < 3 {
		log.Panic("screen {read|press|release|toggle|reset} [-delay=<ms>] [{power|menu|plus|minus|source}]")
	}

	switch os.Args[1] {
	case "read", "press", "release":
		cmd = sn.ScreenCmd{
			Action: sn.ScreenAction(os.Args[1]),
			Button: parseButton(os.Args[2]),
		}

	case "toggle":
		set := flag.NewFlagSet("toggle", flag.ContinueOnError)
		delay := set.Duration("delay", 250*time.Millisecond, "delay between pressing and releasing on toggle")
		util.PanicIfErr(set.Parse(os.Args[2:]))

		cmd = sn.ScreenCmd{
			Action:      sn.ActionToggle,
			Button:      parseButton(set.Arg(0)),
			ToggleDelay: *delay,
		}

	case "reset":
		for _, b := range sn.ScreenButtons() {
			b.Pin().Write(false)
			b.Pin().Mode(false)
		}
		os.Exit(0)

	default:
		log.Panic("screen {read|press|release|toggle|reset} [-delay=<ms>] [{power|menu|plus|minus|source}]")
	}

	return
}

func parseButton(pin string) (button sn.ScreenButton) {
	switch pin {
	case "power":
		button = sn.ButtonPower

	case "menu":
		button = sn.ButtonMenu

	case "plus":
		button = sn.ButtonPlus

	case "minus":
		button = sn.ButtonMinus

	case "source":
		button = sn.ButtonSource

	default:
		log.Panic("possible pin names: power, menu, plus, minus, source")
	}

	return
}
