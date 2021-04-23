package color

type RGBA32 struct {
	Color uint32
}

func (c RGBA32) RGBA() (r, g, b, a uint32) {
	a = (c.Color & 0xff000000)
	r = (c.Color & 0x00ff0000) << 8
	g = (c.Color & 0x0000ff00) << 16
	b = (c.Color & 0x000000ff) << 24

	return
}
