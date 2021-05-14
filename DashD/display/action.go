package display

type Action string

const (
	ActionSet   Action = "set"
	ActionReset Action = "reset"
	ActionGet   Action = "get"
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
		ActionSet,
		ActionReset,
		ActionGet,
	}
}
