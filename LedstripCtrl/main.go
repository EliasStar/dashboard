package main

import (
	"errors"
	"flag"
	"log"
	"strconv"
	"strings"

	hw "github.com/EliasStar/DashboardUtils/Commons/hardware"
	lg "github.com/EliasStar/DashboardUtils/Commons/log"
)

func main() {
	var ledIdentifier string

	flag.StringVar(&ledIdentifier, "leds", "", "")
	flag.StringVar(&ledIdentifier, "l", "", "")

	flag.Parse()

	if colorStr := flag.Arg(0); colorStr != "" {
		color, err := parseColor(colorStr)
		lg.FatalIfErr(err)

		strip, err := hw.MakeLedstrip(hw.LedstripDataPin, 62, true)
		lg.FatalIfErr(err)

		lg.FatalIfErr(strip.Init())
		defer strip.Fini()

		strings.ReplaceAll(ledIdentifier, " ", "")
		if ledIdentifier != "" {
			ledIndicies, err := parseLEDs(ledIdentifier)
			lg.PanicIfErr(err)

			leds := strip.LEDs()
			for _, v := range ledIndicies {
				leds[v] = color
			}
		} else {
			strip.SetStrip(color)
		}

		strip.Render()
	} else {
		log.Fatal("ledstrip [<flags>] <color>")
	}
}

func parseColor(colorStr string) (uint32, error) {
	if strings.HasPrefix(colorStr, "0x") {
		colorStr = strings.TrimPrefix(colorStr, "0x")
	} else if strings.HasPrefix(colorStr, "#") {
		colorStr = strings.TrimPrefix(colorStr, "#")
	}

	color, err := strconv.ParseUint(colorStr, 16, 32)
	if err != nil {
		return 0, errors.New("possible color syntax: 0xRRGGBB, #RRGGBB, RRGGBB")
	}

	return uint32(color), nil
}

func parseLEDs(ledIdentifier string) ([]uint, error) {
	ranges := strings.Split(ledIdentifier, ",")

	var leds []uint

	for _, v := range ranges {
		if strings.Contains(v, "-") {
			ledstrings := strings.Split(v, "-")
			if len(ledstrings) != 2 {
				return nil, errors.New("led range malformed: " + v)
			}

			first, err := strconv.ParseUint(ledstrings[0], 10, 0)
			if err != nil {
				return nil, err
			}

			last, err := strconv.ParseUint(ledstrings[1], 10, 0)
			if err != nil {
				return nil, err
			}

			if last-first <= 1 {
				return nil, errors.New("led range out of bounds: " + v)
			}

			ledrange := make([]uint, (last-first)+1)

			for i := 0; i < len(ledrange); i++ {
				ledrange[i] = uint(first) + uint(i)
			}

			leds = append(leds, ledrange...)
		} else {
			led, err := strconv.ParseUint(v, 10, 0)
			if err != nil {
				return nil, err
			}

			leds = append(leds, uint(led))
		}
	}

	return leds, nil
}
