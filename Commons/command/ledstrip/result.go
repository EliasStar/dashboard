package ledstrip

import "image/color"

type LedstripRst struct {
	Colors []color.Color
	Error  error
}

func (l LedstripRst) OK() bool {
	return l.Error == nil
}

func (l LedstripRst) Err() error {
	return l.Error
}
