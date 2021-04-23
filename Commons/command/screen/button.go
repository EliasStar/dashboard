package screen

import (
	"github.com/EliasStar/DashboardUtils/Commons/hardware"
)

type ScreenButton hardware.Pin

const (
	ButtonPower  ScreenButton = 17
	ButtonMenu   ScreenButton = 27
	ButtonPlus   ScreenButton = 22
	ButtonMinus  ScreenButton = 23
	ButtonSource ScreenButton = 24
)

func (s ScreenButton) IsValid() bool {
	for _, b := range ScreenButtons() {
		if b == s {
			return true
		}
	}

	return false
}

func (s ScreenButton) Pin() hardware.Pin {
	return hardware.Pin(s)
}

func ScreenButtons() []ScreenButton {
	return []ScreenButton{
		ButtonPower,
		ButtonMenu,
		ButtonPlus,
		ButtonMinus,
		ButtonSource,
	}
}
