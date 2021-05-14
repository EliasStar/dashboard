package schedule

type Action string

const (
	ActionRead   Action = "read"
	ActionWrite  Action = "write"
	ActionRemove Action = "remove"
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
		ActionWrite,
		ActionRemove,
	}
}
