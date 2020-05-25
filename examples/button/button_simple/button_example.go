package main

import (
	"github.com/tryor/xcgui"
	// "github.com/codyguo/xcgui"
)

import (
	"os"

	"github.com/tryor/log4go"
)

var Debuglog log4go.Logger
var Infolog log4go.Logger
var Errorlog log4go.Logger

func fileIsExist(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(err)
	}
	return true
}

func Log_Init() {
	os.MkdirAll("logs/debug", os.ModePerm)
	os.MkdirAll("logs/info", os.ModePerm)
	os.MkdirAll("logs/error", os.ModePerm)
	Debuglog = log4go.NewDefaultLogger(log4go.DEBUG)
	if fileIsExist("log_debug.xml") {
		Debuglog.LoadConfiguration("log_debug.xml")
	}

	Infolog = log4go.NewDefaultLogger(log4go.INFO)
	if fileIsExist("log_info.xml") {
		Infolog.LoadConfiguration("log_info.xml")
	}

	Errorlog = log4go.NewDefaultLogger(log4go.ERROR)
	if fileIsExist("log_error.xml") {
		Errorlog.LoadConfiguration("log_error.xml")
	}
	Debuglog.Debug("init logs ok.")
}

func main() {
	Log_Init()
	Debuglog.Debug("into main ...")
	//time.Sleep(time.Second * 1)

	mw, _ := xcgui.NewMainWindow(xcgui.Size{800, 600}, "测试")

	/*
		mw.SetMinimumSize(xcgui.Size{500, 300})

		button1, _ := xcgui.NewButton(mw, xcgui.Rectangle{300, 180, 100, 25})
		button1.SetText("提示信息")

		button2, _ := xcgui.NewButton(mw, xcgui.Rectangle{300, 235, 100, 30})
		button2.SetText("警告信息")

		button1.OnBtnClick(func() {
			xcgui.MsgBox(mw, "提示", "这是一条提示信息...", xcgui.MsgBoxIconInformation)
		})

		button2.OnBtnClick(func() {
			text := button2.GetText()
			Debuglog.Debug("button2 - text:", text)
			xcgui.MsgBox(mw, "警告", "这是一条警告信息!", xcgui.MsgBoxIconWarning)
		})

		img, _ := xcgui.NewImageFromFile("../../img/open.png")

		button1.SetIcon(img)
		button2.SetIcon(img)

		button1.SetIconAlign(xcgui.BUTTON_ICON_ALIGN_RIGHT)
		button2.SetIconAlign(xcgui.BUTTON_ICON_ALIGN_RIGHT)

		color := xcgui.RGB(255, 0, 0)
		button1.AddBkBorder(xcgui.BUTTON_STATE_LEAVE, color, 255, 2)
		button2.AddBkFill(xcgui.BUTTON_STATE_LEAVE, color, 255)

		img_buy, _ := xcgui.NewImageFromFile("../../img/buy.png")

		button1.AddBkImage(xcgui.BUTTON_STATE_LEAVE, img_buy)
		err := mw.Show()
		if err != nil {
			Debuglog.Debug(err)
		}

		xcgui.DebugToFileInfo("中文写入debug文件.")
	*/
	mw.Show()
	mw.Run()

	mw.Close()

}
