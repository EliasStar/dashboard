package ledstrip

import (
	"context"
	"errors"
	"image/color"
	"time"

	"github.com/EliasStar/DashboardUtils/Commons/command"
	"github.com/EliasStar/DashboardUtils/Commons/hardware"
	"github.com/EliasStar/DashboardUtils/Commons/util/misc"
)

type LedstripCmd struct {
	Animation       LedstripAnimation
	LEDs            []uint
	Colors          []color.Color
	AnimationLength time.Duration
}

func (l LedstripCmd) IsValid(ctx context.Context) bool {
	strip, ok := ctx.Value("strip").(hardware.Ledstrip)
	if !ok {
		return false
	}

	for _, v := range l.LEDs {
		if v >= strip.Length() {
			return false
		}
	}

	a := l.Animation.IsValid()
	b := len(l.Colors) == 1 || (len(l.Colors) > 1 && len(l.LEDs) == len(l.Colors))
	c := 0 <= l.AnimationLength
	d := l.AnimationLength.Minutes() <= 5

	return a && b && c && d
}

func (l LedstripCmd) Execute(ctx context.Context) command.Result {
	strip, ok := ctx.Value(misc.LedstripContextKey).(*hardware.Ledstrip)
	if !ok {
		return LedstripRst{
			command.ErrorRst{errors.New("ledstrip not initialized")},
			nil,
		}
	}

	switch l.Animation {
	case AnimationRead:
		var colors []color.Color

		if len(l.LEDs) == 0 {
			colors = strip.GetStripColors()
		} else {
			colors = strip.GetLEDColors(l.LEDs)
		}

		return LedstripRst{command.ErrorRst{}, colors}

	case AnimationWrite:
		if len(l.LEDs) == 0 {
			strip.SetStripColor(l.Colors[0])
		} else if len(l.Colors) == 1 {
			strip.SetLEDColor(l.LEDs, l.Colors[0])
		} else {
			strip.SetLEDColors(l.LEDs, l.Colors)
		}

	// TODO
	case AnimationSprinkle:
		return LedstripRst{
			command.ErrorRst{errors.New("animation not implemented")},
			nil,
		}

	case AnimationFlush:
		return LedstripRst{
			command.ErrorRst{errors.New("animation not implemented")},
			nil,
		}

	case AnimationFlushReverse:
		return LedstripRst{
			command.ErrorRst{errors.New("animation not implemented")},
			nil,
		}
	}

	return nil
}
