package main

import (
	"errors"
	"flag"
	"log"
	"time"

	hw "github.com/EliasStar/DashboardUtils/Commons/hardware"
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
		for _, p := range allPins() {
			p.Write(false)
			p.Mode(false)
		}
	}

	for _, p := range allPins() {
		p.Mode(true)
	}

	if pinName := flag.Arg(0); pinName != "" {
		pin, err := parsePin(pinName)
		util.FatalIfErr(err)

		switch action {
		case "press":
			pin.Write(true)

		case "release":
			pin.Write(false)

		case "toggle":
			val, err := pin.Read()
			util.FatalIfErr(err)

			pin.Write(!val)
			time.Sleep(time.Duration(msToggle) * time.Millisecond)
			pin.Write(val)

		default:
			log.Fatal("possible actions: press, release, toggle")
		}
	} else {
		log.Fatal("screen [<flags>] {power|menu|plus|minus|source}")
	}
}

func allPins() []hw.Pin {
	return []hw.Pin{
		hw.ScreenPowerPin,
		hw.ScreenMenuPin,
		hw.ScreenPlusPin,
		hw.ScreenMinusPin,
		hw.ScreenSourcePin,
	}
}

func parsePin(pin string) (out hw.Pin, err error) {
	switch pin {
	case "Power", "power", "POWER":
		out = hw.ScreenPowerPin

	case "Menu", "menu", "MENU":
		out = hw.ScreenMenuPin

	case "Plus", "plus", "PLUS":
		out = hw.ScreenPlusPin

	case "Minus", "minus", "MINUS":
		out = hw.ScreenMinusPin

	case "Source", "source", "SOURCE":
		out = hw.ScreenSourcePin

	default:
		err = errors.New("possible pin names: power, menu, plus, minus, source")
	}

	return
}
