package color

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
