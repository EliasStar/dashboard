package main

import (
	"fmt"
	"log"
	"net"
	"os"

	lg "github.com/EliasStar/DashboardUtils/Commons/log"
	nt "github.com/EliasStar/DashboardUtils/Commons/net"
)

func main() {
	con, err := net.Dial("tcp", os.Args[1]+":"+nt.DashDPort)
	lg.FatalIfErr(err)

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
