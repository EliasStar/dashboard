package main

import (
	"dashboard/pins"
	"flag"
	"log"
	"time"
)

func main() {
	var action string
	var msToggle uint

	flag.StringVar(&action, "action", "toggle", "`action` to occur: {press|release|toggle}")
	flag.StringVar(&action, "a", "toggle", "`action` to occur: {press|release|toggle}")
	flag.UintVar(&msToggle, "toggleDelay", 250, "`ms` between pressing and releasing on toggle")
	flag.UintVar(&msToggle, "t", 250, "`ms` between pressing and releasing on toggle")

	flag.Parse()

	initGPIO()
	defer closeGPIO()

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
		log.Fatal("display [<flags>] {power|menu|plus|minus|source}")
	}
}

func initGPIO() {
	for _, p := range pins.All() {
		p.Mode(true)
	}
}

func closeGPIO() {
	for _, p := range pins.All() {
		p.Write(false)
		p.Mode(false)
	}
}
