package hardware

import (
	"image/color"

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

func (ws *Ledstrip) LEDs() []uint32 {
	if ws.hasBurnerLED {
		return ws.Leds(0)[1:]
	}

	return ws.Leds(0)
}

func (ws *Ledstrip) LED(index int) *uint32 {
	return &ws.LEDs()[index]
}

func (ws *Ledstrip) SetStrip(color uint32) {
	leds := ws.LEDs()

	for i := 0; i < len(leds); i++ {
		leds[i] = color
	}
}

func (ws *Ledstrip) SetStripColor(c color.Color) {
	r, g, b, _ := c.RGBA()
	ws.SetStrip(r<<16 | g<<8 | b)
}

func (ws *Ledstrip) SetStripRGB(red byte, green byte, blue byte) {
	ws.SetStripColor(color.RGBA{red, green, blue, 0xff})
}
