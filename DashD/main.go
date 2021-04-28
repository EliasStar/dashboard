package main

import (
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"net"
	"os/exec"

	"github.com/EliasStar/DashboardUtils/Commons/command"
	hw "github.com/EliasStar/DashboardUtils/Commons/hardware"
	nt "github.com/EliasStar/DashboardUtils/Commons/net"
	"github.com/EliasStar/DashboardUtils/Commons/util"
	"github.com/EliasStar/DashboardUtils/Commons/util/misc"
)

func main() {
	strip, err := hw.NewLedstrip(misc.LedstripDataPin, misc.LedstripLength, misc.LedstripHasBurnerLED)
	util.FatalIfErr(err)

	util.FatalIfErr(strip.Init())
	defer strip.Fini()

	cmd := exec.Command(misc.DashDBrowser)

	ctx := context.Background()
	ctx = context.WithValue(ctx, misc.LedstripContextKey, strip)
	ctx = context.WithValue(ctx, misc.DisplayContextKey, cmd)

	nt.InitGOB()

	listener, err := net.Listen("tcp", "127.0.0.1:"+misc.DashDPort)
	util.FatalIfErr(err)

	defer listener.Close()
	fmt.Println("Listening on:", listener.Addr())

	for {
		con, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConnection(con, ctx)
	}
}

func handleConnection(con net.Conn, ctx context.Context) {
	addr := con.RemoteAddr()
	dec := gob.NewDecoder(con)
	enc := gob.NewEncoder(con)

	fmt.Println("New Connection:", addr)

	var cmd command.Command
	for dec.Decode(&cmd) == nil {
		fmt.Printf("|%v|> Received: %T\n", addr, cmd)

		var rst command.Result
		rst = command.ErrorRst{errors.New("command invalid")}

		if cmd.IsValid(ctx) {
			rst = cmd.Execute(ctx)
		}

		enc.Encode(&rst)
	}
}
