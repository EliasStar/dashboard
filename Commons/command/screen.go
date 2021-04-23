package command

import (
	"time"

	"github.com/EliasStar/DashboardUtils/Commons/hardware"
)

type ScreenCmd struct {
	Action      string
	Button      hardware.Pin
	ToggleDelay time.Duration
}
