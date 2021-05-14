package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
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
	con, err := net.Dial("tcp", os.Args[1]+":port")
	util.PanicIfErr(err)

	defer con.Close()

	util.InitGOBFull()

	enc := gob.NewEncoder(con)
	dec := gob.NewDecoder(con)
	scn := bufio.NewScanner(os.Stdin)

	for scn.Scan() {
		var cmd command.Command

		// TODO DashConnect Input
		in := strings.Split(scn.Text(), " ")

		switch in[0] {
		case "display":
			cmd = display.Command{display.Action(in[1]), in[2]}

		case "launch":
			cmd = launch.Command{in[1], in[2:]}

		case "ledstrip":
			cmd = ledstrip.Command{}

		case "schedule":
			cmd = schedule.Command{}

		case "screen":
			action := screen.Action(in[1])

			val, err := strconv.ParseUint(in[2], 10, 0)
			if err != nil {
				log.Println(err)
				continue
			}
			btn := screen.Button(val)

			var delay time.Duration
			if action == screen.ActionToggle {
				val, err := strconv.ParseUint(in[3], 10, 0)
				if err != nil {
					log.Println(err)
					continue
				}

				delay = time.Microsecond * time.Duration(val)
			}

			cmd = screen.Command{action, btn, delay}

		default:
			continue
		}

		enc.Encode(&cmd)

		var rst command.Result
		dec.Decode(&rst)
		fmt.Printf("%#v\n", rst)
	}
}
