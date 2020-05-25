package main

import (
	"log"

	"github.com/tryor/xcgui"

	"github.com/tryor/xcgui/xc"
)

func main() {
	mw, _ := xcgui.NewMainWindow(xcgui.Size{300, 200}, "富文本")

	hEle := xc.XRichEdit_Create(20, 40, 250, 50, xc.HXCGUI(mw.GetHWindow()))
	//xc.XRichEdit_EnableAutoWrap(hEle, false)

	// wb.hWindow = xc.XWnd_Create(
	// 	xc.CW_USEDEFAULT,
	// 	xc.CW_USEDEFAULT,
	// 	width,
	// 	height,
	// 	title,
	// 	// xc.StringToUTF16Ptr(title),
	// 	hwndParent,
	// 	xc.Xc_window_style_(style))

	xc.XEdit_CreateEx(20, 120, 250, 60, xc.Edit_type_none, xc.HXCGUI(mw.GetHWindow()))

	log.Println(hEle)

	mw.Show()
	mw.Run()
}
