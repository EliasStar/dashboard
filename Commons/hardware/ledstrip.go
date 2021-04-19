package hardware

import (
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

	return Ledstrip{dev, ledCount, addBurnerLED}, nil
}

type Ledstrip struct {
	*ws2811.WS2811

	ledCount     uint
	hasBurnerLED bool
}

func (ws *Ledstrip) SetLEDColor(index uint, color uint32) {
	if ws.hasBurnerLED {
		index++
	}

	ws.Leds(0)[index] = color
}

func (ws *Ledstrip) SetLEDColorRGB(index uint, red uint8, green uint8, blue uint8) {
	ws.SetLEDColor(index, uint32(red)<<16|uint32(green)<<8|uint32(blue))
}

func (ws *Ledstrip) SetStripColor(color uint32) {
	leds := ws.Leds(0)

	var index uint
	if ws.hasBurnerLED {
		index++
	}

	for ; index < ws.ledCount; index++ {
		leds[index] = color
	}
}

func (ws *Ledstrip) SetStripColorRGB(red uint8, green uint8, blue uint8) {
	ws.SetStripColor(uint32(red)<<16 | uint32(green)<<8 | uint32(blue))
}

func (ws *Ledstrip) GetLEDCount() uint {
	return ws.ledCount
}
