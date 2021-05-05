package misc

import (
	"github.com/EliasStar/DashboardUtils/Commons/hardware"
)

const (
	LedstripDataPin      hardware.Pin = 18
	LedstripLength       uint         = 62
	LedstripHasBurnerLED bool         = true
)

const (
	DashDPort    = "64586"
	DashDBrowser = "chromium-browser"
)

type ContextKey uint

const (
	LedstripContextKey ContextKey = iota
	DisplayContextKey
)
