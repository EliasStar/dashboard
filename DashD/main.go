package main

import (
	"errors"

	"github.com/EliasStar/DashboardUtils/Commons/log"
)

func main() {
	log.FatalIfErr(errors.New("test"))
}
