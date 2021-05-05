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

	"github.com/EliasStar/DashboardUtils/Commons/command"
	"github.com/EliasStar/DashboardUtils/Commons/command/display"
	"github.com/EliasStar/DashboardUtils/Commons/command/launch"
	"github.com/EliasStar/DashboardUtils/Commons/command/ledstrip"
	"github.com/EliasStar/DashboardUtils/Commons/command/schedule"
	"github.com/EliasStar/DashboardUtils/Commons/command/screen"
	nt "github.com/EliasStar/DashboardUtils/Commons/net"
	"github.com/EliasStar/DashboardUtils/Commons/util"
	"github.com/EliasStar/DashboardUtils/Commons/util/misc"
)

func main() {
	con, err := net.Dial("tcp", os.Args[1]+":"+misc.DashDPort)
	util.PanicIfErr(err)

	defer con.Close()

	nt.InitGOB()

	enc := gob.NewEncoder(con)
	dec := gob.NewDecoder(con)
	scn := bufio.NewScanner(os.Stdin)

	for scn.Scan() {
		var cmd command.Command

		// TODO DashConnect Input
		in := strings.Split(scn.Text(), " ")

		switch in[0] {
		case "display":
			cmd = display.DisplayCmd{display.DisplayAction(in[1]), in[2]}

		case "launch":
			cmd = launch.LaunchCmd{in[1], in[2:]}

		case "ledstrip":
			cmd = ledstrip.LedstripCmd{}

		case "schedule":
			cmd = schedule.ScheduleCmd{}

		case "screen":
			action := screen.ScreenAction(in[1])

			val, err := strconv.ParseUint(in[2], 10, 0)
			if err != nil {
				log.Println(err)
				continue
			}
			btn := screen.ScreenButton(val)

			var delay time.Duration
			if action == screen.ActionToggle {
				val, err := strconv.ParseUint(in[3], 10, 0)
				if err != nil {
					log.Println(err)
					continue
				}

				delay = time.Microsecond * time.Duration(val)
			}

			cmd = screen.ScreenCmd{action, btn, delay}

		default:
			continue
		}

		enc.Encode(&cmd)

		var rst command.Result
		dec.Decode(&rst)
		fmt.Printf("%#v\n", rst)
	}
}
