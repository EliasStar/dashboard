package ledstrip

import (
	"image/color"
)

type Result []color.Color

func (r Result) IsOK() bool {
	return true
}
