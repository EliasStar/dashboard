package hardware

import (
	"image/color"

	"github.com/EliasStar/DashboardUtils/Commons/util"
	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

func MakeLedstrip(pin Pin, ledCount uint, addBurnerLED bool) (Ledstrip, error) {
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
		return Ledstrip{}, err
	}

	return Ledstrip{dev, addBurnerLED}, nil
}

type Ledstrip struct {
	*ws2811.WS2811

	hasBurnerLED bool
}

func (ws *Ledstrip) GetLEDs() []uint32 {
	if ws.hasBurnerLED {
		return ws.Leds(0)[1:]
	}

	return ws.Leds(0)
}

func (ws *Ledstrip) SetSingleLEDColor(index uint, c color.Color) {
	r, g, b, _ := c.RGBA()
	ws.GetLEDs()[index] = r<<16 | g<<8 | b
}

func (ws *Ledstrip) SetLEDColor(indicies []uint, c color.Color) {
	leds := ws.GetLEDs()
	r, g, b, _ := c.RGBA()

	for _, v := range indicies {
		leds[v] = r<<16 | g<<8 | b
	}
}

func (ws *Ledstrip) SetLEDColors(indicies []uint, c []color.Color) {
	leds := ws.GetLEDs()

	for i := 0; i < util.Min(len(indicies), len(c)); i++ {
		r, g, b, _ := c[i].RGBA()
		leds[indicies[i]] = r<<16 | g<<8 | b
	}
}

func (ws *Ledstrip) SetStripColor(c color.Color) {
	leds := ws.GetLEDs()
	r, g, b, _ := c.RGBA()

	for i := 0; i < len(leds); i++ {
		leds[i] = r<<16 | g<<8 | b
	}
}
