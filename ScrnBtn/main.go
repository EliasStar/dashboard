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
	. "github.com/EliasStar/DashboardUtils/Commons/command/screen"
	"github.com/EliasStar/DashboardUtils/Commons/util"
	"github.com/EliasStar/DashboardUtils/Commons/util/misc"
)

func main() {
	cmd := parseCommand()
	var rst command.Result

	con, conErr := net.Dial("tcp", "127.0.0.1:"+misc.DashDPort)
	if conErr == nil {
		defer con.Close()

		gob.Register(command.ErrorRst{})
		gob.Register(command.OKRst{})
		gob.Register(ScreenCmd{})
		gob.Register(ScreenRst(false))

		util.PanicIfErr(gob.NewEncoder(con).Encode(&cmd))
		util.PanicIfErr(gob.NewDecoder(con).Decode(&rst))
	} else {
		for _, b := range ScreenButtons() {
			util.PanicIfErr(b.Pin().Mode(true))
		}

		rst = cmd.Execute(context.Background())
	}

	if !rst.IsOK() {
		log.Fatal(rst.Err())
	}

	snRst, ok := rst.(ScreenRst)
	if ok {
		fmt.Println(snRst)
	}
}

func parseCommand() (cmd command.Command) {
	switch os.Args[1] {
	case "read", "press", "release":
		cmd = ScreenCmd{
			Action: ScreenAction(os.Args[1]),
			Button: parseButton(os.Args[2]),
		}

	case "toggle":
		set := flag.NewFlagSet("toggle", flag.ContinueOnError)
		delay := set.Duration("delay", 250*time.Millisecond, "delay between pressing and releasing on toggle")
		util.PanicIfErr(set.Parse(os.Args[2:]))

		cmd = ScreenCmd{
			Action:      ActionToggle,
			Button:      parseButton(set.Arg(0)),
			ToggleDelay: *delay,
		}

	case "reset":
		for _, b := range ScreenButtons() {
			b.Pin().Write(false)
			b.Pin().Mode(false)
		}
		os.Exit(0)

	default:
		log.Panic("screen {read|press|release|toggle|reset} [-delay=<ms>] [{power|menu|plus|minus|source}]")
	}

	return
}

func parseButton(pin string) (button ScreenButton) {
	switch pin {
	case "power":
		button = ButtonPower

	case "menu":
		button = ButtonMenu

	case "plus":
		button = ButtonPlus

	case "minus":
		button = ButtonMinus

	case "source":
		button = ButtonSource

	default:
		log.Fatal("possible pin names: power, menu, plus, minus, source")
	}

	return
}
