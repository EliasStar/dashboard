package main

import (
	"fmt"
	"log"
	"net"
	"os"

	nt "github.com/EliasStar/DashboardUtils/Commons/net"
	"github.com/EliasStar/DashboardUtils/Commons/util"
)

func main() {
	con, err := net.Dial("tcp", os.Args[1]+":"+nt.DashDPort)
	util.FatalIfErr(err)

	defer con.Close()

	var in string
	for {
		fmt.Scan(&in)

		if _, err := fmt.Fprintln(con, in); err != nil {
			log.Println(err)
			return
		}
	}
}
