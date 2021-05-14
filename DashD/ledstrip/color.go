package ledstrip

type RGB struct {
	R, G, B uint8
}

func (c RGB) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R) << 24
	g = uint32(c.G) << 24
	b = uint32(c.B) << 24
	a = 0xff000000

	return
}

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
