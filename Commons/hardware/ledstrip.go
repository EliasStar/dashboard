package hardware

import (
	"image/color"

	"github.com/EliasStar/DashboardUtils/Commons/util"
	col "github.com/EliasStar/DashboardUtils/Commons/util/color"

	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

func NewLedstrip(pin Pin, ledCount uint, addBurnerLED bool) (*Ledstrip, error) {
	opt := ws2811.DefaultOptions
	channel := &opt.Channels[0]

	channel.GpioPin = int(pin)
	channel.LedCount = int(ledCount)
	channel.Brightness = 255

	if addBurnerLED {
		channel.LedCount++
	}

	dev, err := ws2811.MakeWS2811(&opt)
	if err != nil {
		return nil, err
	}

	return &Ledstrip{dev, addBurnerLED}, nil
}

type Ledstrip struct {
	*ws2811.WS2811

	hasBurnerLED bool
}

func (l *Ledstrip) LEDs() []uint32 {
	if l.hasBurnerLED {
		return l.Leds(0)[1:]
	}

	return l.Leds(0)
}

func (l *Ledstrip) Length() uint {
	return uint(len(l.LEDs()))
}

func (l *Ledstrip) GetSingleLEDColor(index uint) color.Color {
	return col.RGBA32{l.LEDs()[index]}
}

func (l *Ledstrip) GetLEDColors(indicies []uint) []color.Color {
	leds := l.LEDs()
	c := make([]color.Color, len(indicies))

	for i, v := range indicies {
		c[i] = col.RGBA32{leds[v]}
	}

	return c
}

func (l *Ledstrip) GetStripColors() []color.Color {
	leds := l.LEDs()
	c := make([]color.Color, len(leds))

	for i, v := range leds {
		c[i] = col.RGBA32{v}
	}

	return c
}

func (l *Ledstrip) SetSingleLEDColor(index uint, c color.Color) {
	r, g, b, _ := c.RGBA()
	l.LEDs()[index] = r<<16 | g<<8 | b
}

func (l *Ledstrip) SetLEDColor(indicies []uint, c color.Color) {
	leds := l.LEDs()
	r, g, b, _ := c.RGBA()

	for _, v := range indicies {
		leds[v] = r<<16 | g<<8 | b
	}
}

func (l *Ledstrip) SetLEDColors(indicies []uint, c []color.Color) {
	leds := l.LEDs()

	for i := 0; i < util.Min(len(indicies), len(c)); i++ {
		r, g, b, _ := c[i].RGBA()
		leds[indicies[i]] = r<<16 | g<<8 | b
	}
}

func (l *Ledstrip) SetStripColor(c color.Color) {
	leds := l.LEDs()
	r, g, b, _ := c.RGBA()

	for i := 0; i < len(leds); i++ {
		leds[i] = r<<16 | g<<8 | b
	}
}
