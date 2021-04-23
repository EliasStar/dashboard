package util

import "log"

func FatalIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func PanicIfErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
