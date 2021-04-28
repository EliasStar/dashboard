package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"

	"github.com/EliasStar/DashboardUtils/Commons/command"
	nt "github.com/EliasStar/DashboardUtils/Commons/net"
	"github.com/EliasStar/DashboardUtils/Commons/util"
	"github.com/EliasStar/DashboardUtils/Commons/util/misc"
)

func main() {
	con, err := net.Dial("tcp", os.Args[1]+":"+misc.DashDPort)
	util.FatalIfErr(err)

	defer con.Close()

	nt.InitGOB()

	enc := gob.NewEncoder(con)
	dec := gob.NewDecoder(con)

	for {
		var cmd command.Command
		// TODO: Input
		enc.Encode(&cmd)

		var rst command.Result
		dec.Decode(&rst)
		fmt.Printf("%#v\n", rst)
	}
}
