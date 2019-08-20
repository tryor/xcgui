package xcgui

import (
	"github.com/tryor/xcgui/xc"
)

type Rectangle struct {
	X, Y, Width, Height int
}

func rectangleFromRECT(r xc.RECT) Rectangle {
	return Rectangle{
		X:      int(r.Left),
		Y:      int(r.Top),
		Width:  int(r.Right - r.Left),
		Height: int(r.Bottom - r.Top),
	}
}

func (r Rectangle) Left() int {
	return r.X
}

func (r Rectangle) Top() int {
	return r.Y
}

func (r Rectangle) Right() int {
	return r.X + r.Width - 1
}

func (r Rectangle) Bottom() int {
	return r.Y + r.Height - 1
}

func (r Rectangle) Location() xc.POINT {
	return xc.POINT{r.X, r.Y}
}

func (r *Rectangle) SetLocation(p xc.POINT) Rectangle {
	r.X = p.X
	r.Y = p.Y

	return *r
}

func (r Rectangle) Size() Size {
	return Size{r.Width, r.Height}
}

func (r *Rectangle) SetSize(s Size) Rectangle {
	r.Width = s.Width
	r.Height = s.Height

	return *r
}

func (r Rectangle) toRECT() xc.RECT {
	return xc.RECT{
		int32(r.X),
		int32(r.Y),
		int32(r.X + r.Width),
		int32(r.Y + r.Height),
	}
}
