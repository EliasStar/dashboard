package command

import (
	"image/color"
)

type LedstripCmd struct {
	Action string
	LEDs   []uint
	Colors []color.Color
}

func (l LedstripCmd) IsValid() bool {

	return false
}

func (l LedstripCmd) Execute() (interface{}, error) {

	return nil, nil
}
