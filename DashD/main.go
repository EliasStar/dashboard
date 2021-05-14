package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os/exec"

	"github.com/EliasStar/Dashboard/DashD/command"
	"github.com/EliasStar/Dashboard/DashD/display"
	"github.com/EliasStar/Dashboard/DashD/ledstrip"
	"github.com/EliasStar/Dashboard/DashD/screen"
	"github.com/EliasStar/Dashboard/DashD/util"
)

func main() {
	for _, b := range screen.Buttons() {
		util.PanicIfErr(b.SetOutput())
	}

	strip, err := ledstrip.New(ledstrip.Pin, ledstrip.Length, ledstrip.HasBurnerLED)
	util.PanicIfErr(err)

	util.PanicIfErr(strip.Init())
	defer strip.Fini()

	cmd := exec.Command(display.Browser)

	ctx := context.Background()
	ctx = context.WithValue(ctx, ledstrip.ContextKey{}, strip)
	ctx = context.WithValue(ctx, display.ContextKey{}, cmd)

	util.InitGOBFull()

	listener, err := net.Listen("tcp", "127.0.0.1:"+util.GetPort())
	util.PanicIfErr(err)

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

		var rst command.Result = command.ErrorRst("command invalid")

		if cmd.IsValid(ctx) {
			rst = cmd.Execute(ctx)
		}

		enc.Encode(&rst)
	}
}
