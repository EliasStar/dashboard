package main

import (
	"errors"
	"flag"
	"image/color"
	"log"
	"strconv"
	"strings"

	hw "github.com/EliasStar/DashboardUtils/Commons/hardware"
	"github.com/EliasStar/DashboardUtils/Commons/util"
	cl "github.com/EliasStar/DashboardUtils/Commons/util/color"
	"github.com/EliasStar/DashboardUtils/Commons/util/misc"
)

func main() {
	var ledIdentifier string

	flag.StringVar(&ledIdentifier, "leds", "", "")
	flag.StringVar(&ledIdentifier, "l", "", "")

	flag.Parse()

	if colorStr := flag.Arg(0); colorStr != "" {
		col, err := parseColor(colorStr)
		util.FatalIfErr(err)

		strip, err := hw.NewLedstrip(misc.LedstripDataPin, misc.LedstripLength, misc.LedstripHasBurnerLED)
		util.FatalIfErr(err)

		util.FatalIfErr(strip.Init())
		defer strip.Fini()

		ledIdentifier = strings.ReplaceAll(ledIdentifier, " ", "")
		if ledIdentifier != "" {
			ledIndicies, err := parseLEDs(ledIdentifier)
			util.PanicIfErr(err)

			strip.SetLEDColor(ledIndicies, col)
		} else {
			strip.SetStripColor(col)
		}

		strip.Render()
	} else {
		log.Fatal("ledstrip [<flags>] <color>")
	}
}

func parseColor(colorStr string) (color.Color, error) {
	if strings.HasPrefix(colorStr, "0x") {
		colorStr = strings.TrimPrefix(colorStr, "0x")
	} else if strings.HasPrefix(colorStr, "#") {
		colorStr = strings.TrimPrefix(colorStr, "#")
	}

	c, err := strconv.ParseUint(colorStr, 16, 32)
	if err != nil {
		return nil, errors.New("possible color syntax: 0xRRGGBB, #RRGGBB, RRGGBB")
	}

	col := uint32(c) | 0xff000000
	return cl.RGBA32{col}, nil
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
