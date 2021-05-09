package main

import (
	"context"
	"encoding/gob"
	"errors"
	"flag"
	"fmt"
	"image/color"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/EliasStar/DashboardUtils/Commons/command"
	. "github.com/EliasStar/DashboardUtils/Commons/command/ledstrip"
	hw "github.com/EliasStar/DashboardUtils/Commons/hardware"
	"github.com/EliasStar/DashboardUtils/Commons/util"
	cl "github.com/EliasStar/DashboardUtils/Commons/util/color"
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
		gob.Register(LedstripCmd{})
		gob.Register(LedstripRst{})

		util.PanicIfErr(gob.NewEncoder(con).Encode(&cmd))
		util.PanicIfErr(gob.NewDecoder(con).Decode(&rst))
	} else {
		strip, err := hw.NewLedstrip(misc.LedstripDataPin, misc.LedstripLength, misc.LedstripHasBurnerLED)
		util.PanicIfErr(err)

		util.PanicIfErr(strip.Init())
		defer strip.Fini()

		ctx := context.Background()
		ctx = context.WithValue(ctx, misc.LedstripContextKey, strip)

		rst = cmd.Execute(ctx)
	}

	if !rst.IsOK() {
		log.Fatal(rst.Err())
	}

	ledRst, ok := rst.(LedstripRst)
	if ok {
		fmt.Println(ledRst)
	}
}

func parseCommand() (cmd command.Command) {
	set := flag.NewFlagSet("all", flag.ContinueOnError)
	ledFilter := set.String("leds", "", "filters `leds` using print custom pages syntax")
	util.PanicIfErr(set.Parse(os.Args[2:]))

	leds, err := parseLEDFilter(*ledFilter)
	util.PanicIfErr(err)

	switch os.Args[1] {
	case "read":
		cmd = LedstripCmd{
			Animation: AnimationRead,
			LEDs:      leds,
		}

	case "write":
		set := flag.NewFlagSet("write", flag.ContinueOnError)
		anim := set.String("anim", "write", "custom `animation` controls how colors are written")
		animLen := set.Duration("anim-len", 5*time.Second, "animation length as duration")
		util.PanicIfErr(set.Parse(os.Args[2:]))

		animation := parseAnimation(*anim)
		colors := parseColors(set.Args())

		cmd = LedstripCmd{
			Animation:       animation,
			LEDs:            leds,
			Colors:          colors,
			AnimationLength: *animLen,
		}

	default:
		log.Panic("ledstrip {read|write} [-leds=<leds>] [-anim=<animation>] [-anim-len=<duration>] [<color>...]")
	}

	return
}

func parseLEDFilter(ledFilter string) (leds []uint, err error) {
	ledFilter = strings.ReplaceAll(ledFilter, " ", "")
	ranges := strings.Split(ledFilter, ",")

	for _, v := range ranges {
		if strings.Contains(v, "-") {
			ledstrings := strings.Split(v, "-")
			if len(ledstrings) != 2 {
				err = errors.New("led range malformed: " + v)
				return
			}

			first, e := strconv.ParseUint(ledstrings[0], 10, 0)
			if err != nil {
				err = e
				return
			}

			last, e := strconv.ParseUint(ledstrings[1], 10, 0)
			if err != nil {
				err = e
				return
			}

			if last-first <= 1 {
				err = errors.New("led range out of bounds: " + v)
				return
			}

			for i := 0; i < int(last-first); i++ {
				leds = append(leds, uint(first)+uint(i))
			}
		} else {
			led, e := strconv.ParseUint(v, 10, 0)
			if err != nil {
				err = e
				return
			}

			leds = append(leds, uint(led))
		}
	}

	return
}

func parseAnimation(anim string) (animation LedstripAnimation) {
	switch anim {
	case "flush":
		animation = AnimationFlush

	case "reverseflush":
		animation = AnimationFlushReverse

	case "write":
		animation = AnimationWrite

	default:
		log.Fatal("possible animations: flush, reverseflush, write")
	}

	return
}

func parseColors(colorStrings []string) (colors []color.Color) {
	for _, c := range colorStrings {
		if strings.HasPrefix(c, "0x") {
			c = strings.TrimPrefix(c, "0x")
		} else if strings.HasPrefix(c, "#") {
			c = strings.TrimPrefix(c, "#")
		}

		col, err := strconv.ParseUint(c, 16, 32)
		if err != nil {
			log.Fatal("possible color syntax: 0xRRGGBB, #RRGGBB, RRGGBB")
		}

		color := cl.RGBA32{Color: uint32(col)}
		colors = append(colors, color)
	}

	return
}
