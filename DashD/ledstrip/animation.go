package ledstrip

type Animation string

const (
	AnimationRead  Animation = "read"
	AnimationWrite Animation = "write"

	AnimationFlush        Animation = "flush_start_end"
	AnimationFlushReverse Animation = "flush_end_start"
)

func (a Animation) IsValid() bool {
	for _, v := range Animations() {
		if a == v {
			return true
		}
	}

	return false
}

func Animations() []Animation {
	return []Animation{
		AnimationRead,
		AnimationWrite,
		AnimationFlushReverse,
		AnimationFlush,
	}
}
