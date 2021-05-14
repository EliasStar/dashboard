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

	"github.com/EliasStar/Dashboard/DashD/command"
	"github.com/EliasStar/Dashboard/DashD/screen"
	"github.com/EliasStar/Dashboard/DashD/util"
)

func main() {
	cmd := parseCommand()
	var rst command.Result

	con, conErr := net.Dial("tcp", "127.0.0.1:"+util.GetPort())
	if conErr == nil {
		defer con.Close()

		util.InitGOBBasic()
		util.InitGOBScreen()

		util.PanicIfErr(gob.NewEncoder(con).Encode(&cmd))
		util.PanicIfErr(gob.NewDecoder(con).Decode(&rst))
	} else {
		for _, b := range screen.Buttons() {
			util.PanicIfErr(b.SetOutput())
		}

		rst = cmd.Execute(context.Background())
	}

	if !rst.IsOK() {
		log.Panic(rst)
	}

	snRst, ok := rst.(screen.Result)
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
		cmd = screen.Command{
			Action: screen.Action(os.Args[1]),
			Button: parseButton(os.Args[2]),
		}

	case "toggle":
		set := flag.NewFlagSet("toggle", flag.ContinueOnError)
		delay := set.Duration("delay", 250*time.Millisecond, "delay between pressing and releasing on toggle")
		util.PanicIfErr(set.Parse(os.Args[2:]))

		cmd = screen.Command{
			Action:      screen.ActionToggle,
			Button:      parseButton(set.Arg(0)),
			ToggleDelay: *delay,
		}

	case "reset":
		for _, b := range screen.Buttons() {
			b.Write(false)
			b.SetInput()
		}
		os.Exit(0)

	default:
		log.Panic("screen {read|press|release|toggle|reset} [-delay=<ms>] [{power|menu|plus|minus|source}]")
	}

	return
}

func parseButton(pin string) (button screen.Button) {
	switch pin {
	case "power":
		button = screen.ButtonPower

	case "menu":
		button = screen.ButtonMenu

	case "plus":
		button = screen.ButtonPlus

	case "minus":
		button = screen.ButtonMinus

	case "source":
		button = screen.ButtonSource

	default:
		log.Panic("possible pin names: power, menu, plus, minus, source")
	}

	return
}
