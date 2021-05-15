package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"image/color"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/EliasStar/Dashboard/DashD/command"
	"github.com/EliasStar/Dashboard/DashD/display"
	"github.com/EliasStar/Dashboard/DashD/launch"
	"github.com/EliasStar/Dashboard/DashD/ledstrip"
	"github.com/EliasStar/Dashboard/DashD/schedule"
	"github.com/EliasStar/Dashboard/DashD/screen"
	"github.com/EliasStar/Dashboard/DashD/util"
)

func main() {
	con, conErr := net.Dial("tcp", os.Args[1])
	util.PanicIfErr(conErr)

	defer con.Close()
	util.InitGOBFull()

	enc := gob.NewEncoder(con)
	dec := gob.NewDecoder(con)
	scn := bufio.NewScanner(os.Stdin)

	for scn.Scan() {
		cmd := parseCommand(scn.Text())
		if cmd == nil {
			continue
		}

		var rst command.Result

		util.PanicIfErr(enc.Encode(&cmd))
		util.PanicIfErr(dec.Decode(&rst))

		if !rst.IsOK() {
			log.Panic(rst)
		}

		fmt.Println(rst)
	}
}

func parseCommand(input string) (cmd command.Command) {
	in := strings.Fields(input)
	if len(in) < 1 {
		log.Println("too few arguments")
		return
	}

	switch in[0] {
	case "display":
		cmd = display.Command{
			Action: display.Action(in[1]),
			URL:    in[2],
		}

	case "launch":
		cmd = launch.Command{
			Executable: in[1],
			Arguments:  in[2:],
		}

	case "ledstrip":
		if len(in) < 5 {
			log.Println("too few arguments")
			return
		}

		length, err := time.ParseDuration(in[2])
		if err != nil {
			log.Println(err)
			return
		}

		var leds []uint
		for _, v := range strings.Split(in[3], ",") {
			if v == "" {
				return
			}

			led, err := strconv.ParseUint(v, 10, 0)
			if err != nil {
				log.Println(err)
				return
			}

			leds = append(leds, uint(led))
		}

		var colors []color.Color
		for _, v := range strings.Split(in[4], ",") {
			if v == "" {
				return
			}

			col, err := strconv.ParseUint(v, 16, 0)
			if err != nil {
				log.Println(err)
				return
			}

			c := ledstrip.RGBA32{
				Color: uint32(col),
			}

			colors = append(colors, c)
		}

		cmd = ledstrip.Command{
			Animation:       ledstrip.Animation(in[1]),
			AnimationLength: length,
			LEDs:            leds,
			Colors:          colors,
		}

	case "schedule":
		cmd = schedule.Command{}

	case "screen":
		if len(in) < 4 {
			log.Println("too few arguments")
			return
		}

		btn, err := strconv.ParseUint(in[2], 10, 0)
		if err != nil {
			log.Println(err)
			return
		}

		delay, err := time.ParseDuration(in[3])
		if err != nil {
			log.Println(err)
			return
		}

		cmd = screen.Command{
			Action:      screen.Action(in[1]),
			Button:      screen.Button(btn),
			ToggleDelay: delay,
		}
	}

	return
}
