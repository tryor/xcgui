package main

import (
	"github.com/tryor/xcgui"

	"github.com/tryor/xcgui/xc"
)

func main() {
	mw, _ := xcgui.NewMainWindow(xcgui.Size{300, 200}, "富文本")

	//hEle := xc.XRichEdit_Create(20, 40, 150, 120, xc.HXCGUI(mw.GetHWindow()))
	//xc.XRichEdit_EnableAutoWrap(hEle, false)

	xc.XEdit_CreateEx(20, 40, 150, 60, xc.Edit_type_none, xc.HXCGUI(mw.GetHWindow()))

	mw.Show()
	mw.Run()
}
