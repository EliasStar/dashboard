package main

import (
	"time"

	"github.com/EliasStar/DashboardUtils/pins"
	"github.com/EliasStar/DashboardUtils/utils"
	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

type ledstrip struct {
	*ws2811.WS2811
}

func (ws *ledstrip) display(color uint32, sleepTime time.Duration) error {
	for i := 0; i < len(ws.Leds(0)); i++ {
		ws.Leds(0)[i] = color

		if err := ws.Render(); err != nil {
			return err
		}

		time.Sleep(sleepTime)
	}

	return nil
}

func main() {
	opt := ws2811.DefaultOptions

	opt.Channels[0].GpioPin = int(pins.Data)
	opt.Channels[0].LedCount = 61
	opt.Channels[0].Brightness = 128
	opt.Channels[0].StripeType = ws2811.WS2812Strip

	dev, err := ws2811.MakeWS2811(&opt)
	utils.FatalIfErr(err)

	utils.FatalIfErr(dev.Init())
	defer dev.Fini()

	strip := ledstrip{dev}

	strip.display(0xff0000, 1000*time.Millisecond)
	strip.display(0x00ff00, 1000*time.Millisecond)
	strip.display(0x0000ff, 1000*time.Millisecond)
	strip.display(0x000000, 500*time.Millisecond)
}
