package util

import "os"

func GetPort() string {
	content, err := os.ReadFile("port.conf")
	PanicIfErr(err)
	return string(content)
}
