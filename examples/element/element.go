package main

import (
	"log"
	"syscall"

	xcgui "github.com/tryor/xcgui/xc"
)

func main() {
	hWindow := xcgui.XWnd_Create(0, 0, 600, 600, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.XBtn_SetType(xcgui.XBtn_Create(10, 5, 60, 20, "Close", xcgui.HXCGUI(hWindow)), xcgui.BUTTON_TYPE_CLOSE)

	hele := xcgui.XEle_Create(100, 100, 100, 100, xcgui.HXCGUI(hWindow))

	btn := xcgui.XBtn_Create(200, 200, 200, 40, "测试", xcgui.HXCGUI(hWindow))

	xcgui.XEle_RegEventC(btn, xcgui.XE_BNCLICK, syscall.NewCallback(func(pbHandled *bool) int {
		v := int32(200)
		var pRect xcgui.RECT
		xcgui.XEle_GetRect(hele, &pRect)
		log.Println("1", pRect)
		pRect.Bottom = v + (pRect.Bottom - pRect.Top)
		pRect.Top = v

		log.Println("2", pRect)

		xcgui.XEle_SetRect(hele, &pRect, true, xcgui.XC_ADJUSTLAYOUT_ALL)
		return 0
	}))

	xcgui.XWnd_ShowWindow(hWindow, xcgui.SW_SHOW)

	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}
