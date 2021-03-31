package main

import (
	"errors"

	"github.com/EliasStar/DashboardUtils/common"
)

const (
	power  common.Pin = 17
	menu   common.Pin = 27
	plus   common.Pin = 22
	minus  common.Pin = 23
	source common.Pin = 24
)

func allPins() []common.Pin {
	return []common.Pin{power, menu, plus, minus, source}
}

func parsePin(pin string) (out common.Pin, err error) {
	switch pin {
	case "Power", "power", "POWER":
		out = power

	case "Menu", "menu", "MENU":
		out = menu

	case "Plus", "plus", "PLUS":
		out = plus

	case "Minus", "minus", "MINUS":
		out = minus

	case "Source", "source", "SOURCE":
		out = source

	default:
		err = errors.New("possible pin names: power, menu, plus, minus, source")
	}

	return
}
