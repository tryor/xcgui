package main

import (
	"log"
	//xcgui "github.com/codyguo/xcgui/xc"
	xcgui "github.com/tryor/xcgui/xc"
)

func main() {
	//xcgui.XC_LoadResource("resource.xml", "")

	hxCGUI := xcgui.XC_LoadLayout("layoutTest.xml", 0)

	//xcgui.XWnd_AdjustLayout(xcgui.HWINDOW(hxCGUI))
	xcgui.XWnd_ShowWindow(xcgui.HWINDOW(hxCGUI), xcgui.SW_SHOW)

	hEdit := xcgui.XEdit_CreateEx(int(10), int(200), int(300), int(50), xcgui.Edit_type_none, hxCGUI)
	log.Println(hEdit)

	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}
