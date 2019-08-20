package xcgui

import (
	"github.com/tryor/xcgui/xc"
)

var (
	winStyle xc.Xc_window_style_ = xc.XC_WINDOW_STYLE_DEFAULT
)

type MainWindow struct {
	WindowBase
}

func NewMainWindow(szie Size, title string) (*MainWindow, error) {
	mw := new(MainWindow)

	if err := InitWindow(
		mw,
		nil,
		szie.Width,
		szie.Height,
		title,
		winStyle); err != nil {
		return nil, err
	}

	succeeded := false

	go func() {
		if !succeeded {
			newErrorNoPanic("NewMainWindow")
		}
	}()

	succeeded = true

	return mw, nil
}
