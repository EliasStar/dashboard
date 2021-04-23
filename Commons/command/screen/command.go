package screen

import (
	"time"
)

type ScreenCmd struct {
	Action      ScreenAction
	Button      ScreenButton
	ToggleDelay time.Duration
}

func (s ScreenCmd) IsValid() bool {
	a := s.Action.IsValid()
	b := s.Button.IsValid()
	c := 0 <= s.ToggleDelay.Milliseconds()
	d := s.ToggleDelay.Milliseconds() <= 5000

	return a && b && c && d
}

func (s ScreenCmd) Execute() (interface{}, error) {
	switch s.Action {
	case ActionIsPressed:
		return s.Button.Pin().Read()

	case ActionPress:
		s.Button.Pin().Write(true)

	case ActionRelease:
		s.Button.Pin().Write(false)

	case ActionToggle:
		s.Button.Pin().Toggle(s.ToggleDelay)
	}

	return nil, nil
}
