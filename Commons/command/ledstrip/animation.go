package ledstrip

type LedstripAnimation string

const (
	AnimationRead  LedstripAnimation = "read"
	AnimationWrite LedstripAnimation = "write"

	AnimationSprinkle     LedstripAnimation = "sprinkle"
	AnimationFlush        LedstripAnimation = "flush_start_end"
	AnimationFlushReverse LedstripAnimation = "flush_end_start"
)

func (s LedstripAnimation) IsValid() bool {
	for _, a := range ScreenActions() {
		if a == s {
			return true
		}
	}

	return false
}

func ScreenActions() []LedstripAnimation {
	return []LedstripAnimation{
		AnimationRead,
		AnimationWrite,
		AnimationSprinkle,
		AnimationFlushReverse,
		AnimationFlush,
	}
}
