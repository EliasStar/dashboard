package display

type DisplayAction string

const (
	ActionSet   DisplayAction = "set"
	ActionReset DisplayAction = "reset"
	ActionGet   DisplayAction = "get"
)

func (s DisplayAction) IsValid() bool {
	for _, a := range DisplayActions() {
		if a == s {
			return true
		}
	}

	return false
}

func DisplayActions() []DisplayAction {
	return []DisplayAction{
		ActionSet,
		ActionReset,
		ActionGet,
	}
}
