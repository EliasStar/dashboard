package main

import (
	"flag"
	"log"
	"time"

	"github.com/EliasStar/DashboardUtils/pins"
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
		for _, p := range pins.All() {
			p.Write(false)
			p.Mode(false)
		}
	}

	for _, p := range pins.All() {
		p.Mode(true)
	}

	if pinName := flag.Arg(0); pinName != "" {
		pin, err := pins.From(pinName)
		if err != nil {
			log.Fatal(err)
		}

		switch action {
		case "press":
			pin.Write(true)

		case "release":
			pin.Write(false)

		case "toggle":
			val, err := pin.Read()
			if err != nil {
				log.Fatal(err)
			}

			pin.Write(!val)
			time.Sleep(time.Duration(msToggle) * time.Millisecond)
			pin.Write(val)

		default:
			log.Fatal("possible actions: press, release, toggle")
		}
	} else {
		log.Fatal("screen [<flags>] [{power|menu|plus|minus|source}]")
	}
}
