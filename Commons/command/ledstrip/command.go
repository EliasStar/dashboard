package ledstrip

import (
	"context"
	"image/color"
	"time"

	cmd "github.com/EliasStar/DashboardUtils/Commons/command"
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
	strip, ok := ctx.Value(misc.LedstripContextKey).(*hardware.Ledstrip)
	if !ok {
		return false
	}

	for _, v := range l.LEDs {
		if int(v) >= strip.Length() {
			return false
		}
	}

	a := l.Animation.IsValid()
	b := 0 <= l.AnimationLength
	c := l.AnimationLength.Minutes() <= 5

	d := l.Animation == AnimationRead && len(l.Colors) == 0
	e := len(l.Colors) == 1 || (len(l.Colors) > 1 && len(l.LEDs) == len(l.Colors))

	return a && b && c && (d || e)
}

func (l LedstripCmd) Execute(ctx context.Context) cmd.Result {
	strip, ok := ctx.Value(misc.LedstripContextKey).(*hardware.Ledstrip)
	if !ok {
		return cmd.ErrorRst("ledstrip not initialized")
	}

	switch l.Animation {
	case AnimationRead:
		var colors []color.Color

		if len(l.LEDs) == 0 {
			colors = strip.GetStripColors()
		} else {
			colors = strip.GetLEDColors(l.LEDs)
		}

		return LedstripRst(colors)

	case AnimationWrite:
		if len(l.LEDs) == 0 {
			strip.SetStripColor(l.Colors[0])
		} else if len(l.Colors) == 1 {
			strip.SetLEDColor(l.LEDs, l.Colors[0])
		} else {
			strip.SetLEDColors(l.LEDs, l.Colors)
		}

	case AnimationFlush:
		if len(l.LEDs) == 0 {
			for i := 0; i < strip.Length(); i++ {
				strip.SetSingleLEDColor(uint(i), l.Colors[0])
				time.Sleep(l.AnimationLength / time.Duration(strip.Length()))
			}
		} else if len(l.Colors) == 1 {
			for i := 0; i < len(l.LEDs); i++ {
				strip.SetSingleLEDColor(l.LEDs[i], l.Colors[0])
				time.Sleep(l.AnimationLength / time.Duration(len(l.LEDs)))
			}
		} else {
			for i := 0; i < len(l.LEDs); i++ {
				strip.SetSingleLEDColor(l.LEDs[i], l.Colors[i])
				time.Sleep(l.AnimationLength / time.Duration(len(l.LEDs)))
			}
		}

	case AnimationFlushReverse:
		if len(l.LEDs) == 0 {
			for i := strip.Length() - 1; i >= 0; i-- {
				strip.SetSingleLEDColor(uint(i), l.Colors[0])
				time.Sleep(l.AnimationLength / time.Duration(strip.Length()))
			}
		} else if len(l.Colors) == 1 {
			for i := len(l.LEDs) - 1; i >= 0; i-- {
				strip.SetSingleLEDColor(l.LEDs[i], l.Colors[0])
				time.Sleep(l.AnimationLength / time.Duration(len(l.LEDs)))
			}
		} else {
			for i := len(l.LEDs) - 1; i >= 0; i-- {
				strip.SetSingleLEDColor(l.LEDs[i], l.Colors[i])
				time.Sleep(l.AnimationLength / time.Duration(len(l.LEDs)))
			}
		}
	}

	return cmd.NewResultFromError(strip.Render())
}
