package ledstrip

type LedstripAnimation string

const (
	AnimationRead  LedstripAnimation = "read"
	AnimationWrite LedstripAnimation = "write"

	AnimationFlush        LedstripAnimation = "flush_start_end"
	AnimationFlushReverse LedstripAnimation = "flush_end_start"
)

func (s LedstripAnimation) IsValid() bool {
	for _, a := range LedstripAnimations() {
		if a == s {
			return true
		}
	}

	return false
}

func LedstripAnimations() []LedstripAnimation {
	return []LedstripAnimation{
		AnimationRead,
		AnimationWrite,
		AnimationFlushReverse,
		AnimationFlush,
	}
}
