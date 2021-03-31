package main

import (
	"github.com/EliasStar/DashboardUtils/common"

	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

func newLedstrip(pin common.Pin, ledCount uint, addBurnerLED bool) (ledstrip, error) {
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
		return ledstrip{}, err
	}

	return ledstrip{dev, ledCount, addBurnerLED}, nil
}

type ledstrip struct {
	*ws2811.WS2811

	ledCount     uint
	hasBurnerLED bool
}

func (ws *ledstrip) setLEDColor(index uint, color uint32) {
	if ws.hasBurnerLED {
		index++
	}

	ws.Leds(0)[index] = color
}

func (ws *ledstrip) setLEDColorRGB(index uint, red uint8, green uint8, blue uint8) {
	ws.setLEDColor(index, uint32(red)<<16|uint32(green)<<8|uint32(blue))
}

func (ws *ledstrip) setStripColor(color uint32) {
	leds := ws.Leds(0)

	var index uint
	if ws.hasBurnerLED {
		index++
	}

	for ; index < ws.ledCount; index++ {
		leds[index] = color
	}
}

func (ws *ledstrip) setStripColorRGB(red uint8, green uint8, blue uint8) {
	ws.setStripColor(uint32(red)<<16 | uint32(green)<<8 | uint32(blue))
}

func (ws *ledstrip) getLEDCount() uint {
	return ws.ledCount
}
