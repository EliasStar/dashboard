package screen

type Action string

const (
	ActionRead Action = "read"

	ActionPress   Action = "press"
	ActionRelease Action = "release"

	ActionToggle Action = "toggle"
)

func (a Action) IsValid() bool {
	for _, v := range Actions() {
		if a == v {
			return true
		}
	}

	return false
}

func Actions() []Action {
	return []Action{
		ActionRead,
		ActionPress,
		ActionRelease,
		ActionToggle,
	}
}
