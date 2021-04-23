package main

import (
	"errors"
	"flag"
	"log"
	"time"

	"github.com/EliasStar/DashboardUtils/Commons/command/screen"
	"github.com/EliasStar/DashboardUtils/Commons/util"
)

func main() {
	var action string
	var msToggle uint
	var reset bool

	flag.StringVar(&action, "action", "toggle", "`action` to occur: {press|release|toggle}")
	flag.StringVar(&action, "a", "toggle", "`action` to occur: {press|release|toggle}")

	flag.UintVar(&msToggle, "toggleDelay", 250, "`ms` between pressing and releasing on toggle")
	flag.UintVar(&msToggle, "t", 250, "`ms` between pressing and releasing on toggle")

	flag.BoolVar(&reset, "reset", false, "reset gpios")
	flag.BoolVar(&reset, "r", false, "reset gpios")

	flag.Parse()

	if reset {
		for _, b := range screen.ScreenButtons() {
			b.Pin().Write(false)
			b.Pin().Mode(false)
		}
	}

	for _, b := range screen.ScreenButtons() {
		b.Pin().Mode(true)
	}

	if btnName := flag.Arg(0); btnName != "" {
		btn, err := parseButton(btnName)
		util.FatalIfErr(err)

		switch action {
		case "press":
			err := btn.Pin().Write(true)
			util.FatalIfErr(err)

		case "release":
			err := btn.Pin().Write(false)
			util.FatalIfErr(err)

		case "toggle":
			err := btn.Pin().Toggle(time.Duration(msToggle) * time.Millisecond)
			util.FatalIfErr(err)

		default:
			log.Fatal("possible actions: press, release, toggle")
		}
	} else {
		log.Fatal("screen [<flags>] {power|menu|plus|minus|source}")
	}
}

func parseButton(pin string) (out screen.ScreenButton, err error) {
	switch pin {
	case "Power", "power", "POWER":
		out = screen.ButtonPower

	case "Menu", "menu", "MENU":
		out = screen.ButtonMenu

	case "Plus", "plus", "PLUS":
		out = screen.ButtonPlus

	case "Minus", "minus", "MINUS":
		out = screen.ButtonMinus

	case "Source", "source", "SOURCE":
		out = screen.ButtonSource

	default:
		err = errors.New("possible pin names: power, menu, plus, minus, source")
	}

	return
}
