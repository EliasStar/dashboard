package pins

import (
	"errors"
	"os/exec"
	"strconv"
)

type Pin uint

const (
	Power  Pin = 17
	Menu   Pin = 27
	Plus   Pin = 22
	Minus  Pin = 23
	Source Pin = 24
	Data   Pin = 18
)

func (p Pin) Mode(out bool) {
	val := "in"
	if out {
		val = "out"
	}

	exec.Command("gpio", "-g", "mode", p.String(), val).Run()
}

func (p Pin) Write(value bool) {
	val := "0"
	if value {
		val = "1"
	}

	exec.Command("gpio", "-g", "write", p.String(), val).Run()
}

func (p Pin) Read() (value bool, err error) {
	out, err := exec.Command("gpio", "-g", "read", p.String()).Output()
	value = string(out[0]) == "1"
	return
}

func (p Pin) String() string {
	return strconv.Itoa(int(p))
}

func All() []Pin {
	return []Pin{Power, Menu, Plus, Minus, Source}
}

func From(pin string) (out Pin, err error) {
	switch pin {
	case "Power", "power", "POWER":
		out = Power

	case "Menu", "menu", "MENU":
		out = Menu

	case "Plus", "plus", "PLUS":
		out = Plus

	case "Minus", "minus", "MINUS":
		out = Minus

	case "Source", "source", "SOURCE":
		out = Source

	default:
		err = errors.New("possible pin names: power, menu, plus, minus, source")
	}

	return
}
