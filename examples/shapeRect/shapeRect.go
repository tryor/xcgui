package main

import (
	"log"
	"syscall"

	xcgui "github.com/tryor/xcgui/xc"
)

func main() {
	hWindow := xcgui.XWnd_Create(0, 0, 800, 600, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.CloseBtn(hWindow)

	xcgui.XWnd_RegEventC(hWindow, xcgui.WM_PAINT, syscall.NewCallback(func(hDraw xcgui.HDRAW, pbHandled *xcgui.BOOL) int {
		log.Println("WM_PAINT ", hDraw)
		xcgui.XDraw_EnableSmoothingMode(hDraw, true)
		return 0
	}))

	//int CALLBACK OnWndDrawWindow(HDRAW hDraw,BOOL *pbHandled);

	//xcgui.XDraw_EnableSmoothingMode(xcgui.HDRAW(hWindow), true)

	//    hShapeRect
	hShape := xcgui.XShapeRect_Create(20, 76, 760, 436, xcgui.HXCGUI(hWindow))
	xcgui.XShapeRect_EnableRoundAngle(hShape, true)
	xcgui.XShapeRect_SetRoundAngle(hShape, 28, 28)
	xcgui.XShapeRect_EnableBorder(hShape, false)
	xcgui.XShapeRect_EnableFill(hShape, true)

	xcgui.XShapeRect_SetFillColor(hShape, xcgui.RGB(0, 255, 255), 1)

	xcgui.XWnd_ShowWindow(hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}
