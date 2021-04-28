package schedule

type ScheduleAction string

const (
	ActionRead   ScheduleAction = "read"
	ActionWrite  ScheduleAction = "write"
	ActionRemove ScheduleAction = "remove"
)

func (s ScheduleAction) IsValid() bool {
	for _, a := range ScheduleActions() {
		if a == s {
			return true
		}
	}

	return false
}

func ScheduleActions() []ScheduleAction {
	return []ScheduleAction{
		ActionRead,
		ActionWrite,
		ActionRemove,
	}
}
