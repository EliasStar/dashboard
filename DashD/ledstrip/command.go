package ledstrip

import (
	"context"
	"image/color"
	"time"

	"github.com/EliasStar/Dashboard/DashD/command"
)

type ContextKey struct{}

type Command struct {
	Animation       Animation
	LEDs            []uint
	Colors          []color.Color
	AnimationLength time.Duration
}

func (c Command) IsValid(ctx context.Context) bool {
	strip, ok := ctx.Value(ContextKey{}).(*Ledstrip)
	if !ok {
		return false
	}

	for _, v := range c.LEDs {
		if int(v) >= strip.Length() {
			return false
		}
	}

	a := c.Animation == AnimationRead && len(c.Colors) == 0
	b := len(c.Colors) == 1 || (len(c.Colors) > 1 && len(c.LEDs) == len(c.Colors))

	x := c.Animation.IsValid()
	y := 0 <= c.AnimationLength
	z := c.AnimationLength.Minutes() <= 5

	return (a || b) && x && y && z
}

func (c Command) Execute(ctx context.Context) command.Result {
	strip, ok := ctx.Value(ContextKey{}).(*Ledstrip)
	if !ok {
		return command.ErrorRst("ledstrip not initialized")
	}

	switch c.Animation {
	case AnimationRead:
		var colors []color.Color

		if len(c.LEDs) == 0 {
			colors = strip.GetStripColors()
		} else {
			colors = strip.GetLEDColors(c.LEDs)
		}

		return Result(colors)

	case AnimationWrite:
		if len(c.LEDs) == 0 {
			strip.SetStripColor(c.Colors[0])
		} else if len(c.Colors) == 1 {
			strip.SetLEDColor(c.LEDs, c.Colors[0])
		} else {
			strip.SetLEDColors(c.LEDs, c.Colors)
		}

	case AnimationFlush:
		if len(c.LEDs) == 0 {
			for i := 0; i < strip.Length(); i++ {
				strip.SetSingleLEDColor(uint(i), c.Colors[0])
				time.Sleep(c.AnimationLength / time.Duration(strip.Length()))
			}
		} else if len(c.Colors) == 1 {
			for i := 0; i < len(c.LEDs); i++ {
				strip.SetSingleLEDColor(c.LEDs[i], c.Colors[0])
				time.Sleep(c.AnimationLength / time.Duration(len(c.LEDs)))
			}
		} else {
			for i := 0; i < len(c.LEDs); i++ {
				strip.SetSingleLEDColor(c.LEDs[i], c.Colors[i])
				time.Sleep(c.AnimationLength / time.Duration(len(c.LEDs)))
			}
		}

	case AnimationFlushReverse:
		if len(c.LEDs) == 0 {
			for i := strip.Length() - 1; i >= 0; i-- {
				strip.SetSingleLEDColor(uint(i), c.Colors[0])
				time.Sleep(c.AnimationLength / time.Duration(strip.Length()))
			}
		} else if len(c.Colors) == 1 {
			for i := len(c.LEDs) - 1; i >= 0; i-- {
				strip.SetSingleLEDColor(c.LEDs[i], c.Colors[0])
				time.Sleep(c.AnimationLength / time.Duration(len(c.LEDs)))
			}
		} else {
			for i := len(c.LEDs) - 1; i >= 0; i-- {
				strip.SetSingleLEDColor(c.LEDs[i], c.Colors[i])
				time.Sleep(c.AnimationLength / time.Duration(len(c.LEDs)))
			}
		}
	}

	return command.ResultFromError(strip.Render())
}
