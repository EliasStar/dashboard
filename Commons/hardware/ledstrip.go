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

func (l *Ledstrip) Length() int {
	return len(l.LEDs())
}

func (l *Ledstrip) SetSingleLEDColor(index uint, c color.Color) {
	if int(index) >= l.Length() {
		return
	}

	r, g, b, _ := c.RGBA()
	l.LEDs()[index] = r>>8 | g>>16 | b>>24
}

func (l *Ledstrip) GetSingleLEDColor(index uint) color.Color {
	if int(index) >= l.Length() {
		return nil
	}

	return col.RGBA32{Color: l.LEDs()[index]}
}

func (l *Ledstrip) SetLEDColor(indicies []uint, c color.Color) {
	for _, v := range indicies {
		l.SetSingleLEDColor(v, c)
	}
}

func (l *Ledstrip) SetLEDColors(indicies []uint, c []color.Color) {
	for i := 0; i < util.Min(len(indicies), len(c)); i++ {
		l.SetSingleLEDColor(indicies[i], c[i])
	}
}

func (l *Ledstrip) GetLEDColors(indicies []uint) (c []color.Color) {
	for _, v := range indicies {
		c = append(c, l.GetSingleLEDColor(v))
	}

	return
}

func (l *Ledstrip) SetStripColor(c color.Color) {
	leds := l.LEDs()
	r, g, b, _ := c.RGBA()

	for i := 0; i < len(leds); i++ {
		leds[i] = r>>8 | g>>16 | b>>24
	}
}

func (l *Ledstrip) GetStripColors() (c []color.Color) {
	for _, v := range l.LEDs() {
		c = append(c, col.RGBA32{Color: v})
	}

	return
}
