package command

import "image/color"

type LedstripCmd struct {
	Action string
	LEDs   []uint
	Colors []color.RGBA
}
