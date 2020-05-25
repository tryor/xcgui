package xc

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"sync/atomic"
	"syscall"
	"unicode/utf8"
	"unsafe"
)

const (
	TRUE  = 1
	FALSE = 0
	NULL  = 0
)

const (
	xcDll = "XCGUI.dll"
)

var (
	xcDLL *syscall.DLL
)

var (
	// Functions
	xC_UnicodeToAnsi   *syscall.Proc
	xC_AnsiToUnicode   *syscall.Proc
	xC_DebugToFileInfo *syscall.Proc
	xC_IsHELE          *syscall.Proc
	xC_IsHWINDOW       *syscall.Proc
	xC_IsShape         *syscall.Proc
	xC_IsHXCGUI        *syscall.Proc
	xC_hWindowFromHWnd *syscall.Proc
	xC_IsSViewExtend   *syscall.Proc
	xC_GetObjectType   *syscall.Proc
	xC_GetObjectByID   *syscall.Proc
	//	xC_GetResIDValue         *syscall.Proc
	xC_SetPaintFrequency     *syscall.Proc
	xC_RectInRect            *syscall.Proc
	xC_CombineRect           *syscall.Proc
	xC_ShowLayoutFrame       *syscall.Proc
	xC_EnableDebugFile       *syscall.Proc
	xC_SetLayoutFrameColor   *syscall.Proc
	xC_EnableErrorMessageBox *syscall.Proc
	xInitXCGUI               *syscall.Proc
	xRunXCGUI                *syscall.Proc
	xExitXCGUI               *syscall.Proc

	xC_SetTextRenderingHint    *syscall.Proc
	xDraw_EnableSmoothingMode  *syscall.Proc
	xDraw_EnableWndTransparent *syscall.Proc

	xC_CallUiThread *syscall.Proc

	xC_PostMessage *syscall.Proc
	xC_SendMessage *syscall.Proc
)

type POINT struct {
	X, Y int
}

type RECT struct {
	Left, Top, Right, Bottom int32
}

type SIZE struct {
	CX, CY int32
}

type BorderSize_ struct {
	Left, Top, Right, Bottom int
}

type (
	BOOL    int32
	HRESULT int32
)

// 1.初始化UI库
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.LockOSThread()

	if FileExist(xcDll) {
		xcDLL = syscall.MustLoadDLL(xcDll)
	} else if FileExist("lib/" + xcDll) {
		xcDLL = syscall.MustLoadDLL("lib/" + xcDll)
	} else if FileExist("../lib/" + xcDll) {
		xcDLL = syscall.MustLoadDLL("../lib/" + xcDll)
	} else if FileExist("../../lib/" + xcDll) {
		xcDLL = syscall.MustLoadDLL("../../lib/" + xcDll)
	} else {
		// 下载XCGUI.dll
		//downLoadXCGUIDll()

		if FileExist(xcDll) {
			xcDLL = syscall.MustLoadDLL(xcDll)
		} else {
			panic("xcgui library not found,XCGUI.dll or ./lib/XCGUI.dll or ../lib/XCGUI.dll.")
		}

	}

	// Functions
	xC_UnicodeToAnsi = xcDLL.MustFindProc("XC_UnicodeToAnsi")
	xC_AnsiToUnicode = xcDLL.MustFindProc("XC_AnsiToUnicode")
	xC_DebugToFileInfo = xcDLL.MustFindProc("XC_DebugToFileInfo")
	xC_IsHELE = xcDLL.MustFindProc("XC_IsHELE")
	xC_IsHWINDOW = xcDLL.MustFindProc("XC_IsHWINDOW")
	xC_IsShape = xcDLL.MustFindProc("XC_IsShape")
	xC_IsHXCGUI = xcDLL.MustFindProc("XC_IsHXCGUI")
	xC_hWindowFromHWnd = xcDLL.MustFindProc("XC_hWindowFromHWnd")
	xC_IsSViewExtend = xcDLL.MustFindProc("XC_IsSViewExtend")
	xC_GetObjectType = xcDLL.MustFindProc("XC_GetObjectType")
	xC_GetObjectByID = xcDLL.MustFindProc("XC_GetObjectByID")
	//xC_GetResIDValue = xcDLL.MustFindProc("XC_GetResIDValue")
	xC_SetPaintFrequency = xcDLL.MustFindProc("XC_SetPaintFrequency")
	xC_RectInRect = xcDLL.MustFindProc("XC_RectInRect")
	xC_CombineRect = xcDLL.MustFindProc("XC_CombineRect")
	xC_ShowLayoutFrame = xcDLL.MustFindProc("XC_ShowLayoutFrame")
	xC_EnableDebugFile = xcDLL.MustFindProc("XC_EnableDebugFile")
	xC_SetLayoutFrameColor = xcDLL.MustFindProc("XC_SetLayoutFrameColor")
	xC_EnableErrorMessageBox = xcDLL.MustFindProc("XC_EnableErrorMessageBox")
	xInitXCGUI = xcDLL.MustFindProc("XInitXCGUI")
	xRunXCGUI = xcDLL.MustFindProc("XRunXCGUI")
	xExitXCGUI = xcDLL.MustFindProc("XExitXCGUI")

	xC_SetTextRenderingHint = xcDLL.MustFindProc("XC_SetTextRenderingHint")
	xDraw_EnableSmoothingMode = xcDLL.MustFindProc("XDraw_EnableSmoothingMode")
	xDraw_EnableWndTransparent = xcDLL.MustFindProc("XDraw_EnableWndTransparent")

	xC_CallUiThread = xcDLL.MustFindProc("XC_CallUiThread")
	xC_PostMessage = xcDLL.MustFindProc("XC_PostMessage")
	xC_SendMessage = xcDLL.MustFindProc("XC_SendMessage")

	ret, _, _ := xInitXCGUI.Call(StringToUintPtr("XCGUI Library For Go"))
	// XCGUI的返回值: true 为 1 ，false 为 0
	if ret != TRUE {
		panic("XInitXCGUI call failed.")
	}

	// 默认不生成xcgui_debug.txt
	XC_EnableDebugFile(false)
}

func callUiThread(id uint32) int {
	f, ok := callUiThreadFuns.Load(id)
	if ok {
		callUiThreadFuns.Delete(id)
		cf := f.(*callbackFunc)
		return cf.f(cf.param)
	} else {
		log.Println("callUiThread, callbackFunc is nil")
	}
	return 0
}

var callUiThreadPtr = syscall.NewCallback(callUiThread)
var callUiThreadFuns sync.Map
var callUiThreadId uint32

type callbackFunc struct {
	f     func(data int) int
	param int
}

type callUiThreadType = func(data int) int

//int CALLBACK CallUiThread(int data)
func XC_CallUiThread(f callUiThreadType, lpParam int) int {
	id := atomic.AddUint32(&callUiThreadId, 1)
	callUiThreadFuns.Store(id, &callbackFunc{f, lpParam})
	r, _, _ := xC_CallUiThread.Call(callUiThreadPtr, uintptr(id))
	//r, _, _ := xC_CallUiThread.Call(syscall.NewCallback(f), uintptr(lpParam))
	return int(r)
}

//f func(data int) int
// func XC_CallUiThread2(f uintptr, lpParam int) int {
// 	r, _, _ := xC_CallUiThread.Call(f, uintptr(lpParam))
// 	return int(r)
// }

/* 由于在init已经调用过了，这里只留个函数名字
初始化界面库.

参数:
	pText 保留参数.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XInitXCGUI() bool {
	//	pText := ""

	//	if len(arg) == 1 {
	//		pText = arg[0]
	//	}

	//	ret, _, _ := xInitXCGUI.Call(
	//		StringToUintPtr(pText),
	//		0,
	//		0)
	//	return ret == TRUE
	return true
}

/*
Unicode转换Ansi编码,.

参数:
	[in] pIn 指向待转换的Unicode字符串指针.
	[in] inLen pIn字符数量.
	[out] pOut 指向接收转换后的Ansi字符串缓冲区指针.
	[in] outLen pOut缓冲区大小,字节单位.
返回:
	如果成功,返回写入接收缓冲区字节数量.
*/
func XC_UnicodeToAnsi(pIn *uint16, inLen int, pOut *byte, outLen int) int {
	ret, _, _ := xC_UnicodeToAnsi.Call(
		uintptr(unsafe.Pointer(pIn)),
		uintptr(inLen),
		uintptr(unsafe.Pointer(pOut)),
		uintptr(outLen))

	return int(ret)
}

/*
Ansi转换Unicode编码,.

参数:
	[in] pIn 指向待转换的Ansi字符串指针.
	[in] inLen pIn字符数量.
	[out] pOut 指向接收转换后的Unicode字符串缓冲区指针.
	[in] outLen pOut缓冲区大小,字符wchar_t单位.
返回:
	如果成功,返回写入接收缓冲区字符wchar_t数量.
*/
func XC_AnsiToUnicode(pIn *uint16, inLen int, pOut *byte, outLen int) int {
	ret, _, _ := xC_AnsiToUnicode.Call(
		uintptr(unsafe.Pointer(pIn)),
		uintptr(inLen),
		uintptr(unsafe.Pointer(pOut)),
		uintptr(outLen))

	return int(ret)
}

// 打印调试信息到文件xcgui_debug.txt.
func XC_DebugToFileInfo(pInfo *byte) {
	// fmt.Println("正在打印debug", pInfo)
	// p := (*struct {
	// 	pInfo uintptr
	// 	len   int
	// })(unsafe.Pointer(&pInfo))
	// fmt.Println("p.info:", p.pInfo)

	// xC_DebugToFileInfo.Call(p.pInfo)

	xC_DebugToFileInfo.Call(uintptr(unsafe.Pointer(pInfo)))
}

/*
判断是否为元素句柄.

参数:
	hEle 元素句柄.
返回:
	成功返回TRUE,否则相反.
*/
func XC_IsHELE(hEle HXCGUI) bool {
	ret, _, _ := xC_IsHELE.Call(uintptr(hEle))

	return ret == TRUE
}

/*
判断是否为窗口句柄.

参数:
	hWindow 窗口句柄.
返回:
	成功返回TRUE,否则相反.
*/
func XC_IsHWINDOW(hWindow HXCGUI) bool {
	ret, _, _ := xC_IsHWINDOW.Call(uintptr(hWindow))

	return ret == TRUE
}

/*
判断是否为形状对象.

参数:
	hShape 形状对象句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XC_IsShape(hShape HXCGUI) bool {
	ret, _, _ := xC_IsShape.Call(uintptr(hShape))

	return ret == TRUE
}

/*
判断句柄是否拥有该类型.

参数:
	hXCGUI 炫彩句柄.
	nType 句柄类型.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XC_IsHXCGUI(hXCGUI HXCGUI, nType XC_OBJECT_TYPE) bool {
	ret, _, _ := xC_IsHXCGUI.Call(
		uintptr(hXCGUI),
		uintptr(nType))

	return ret == TRUE
}

/*
通过窗口HWND句柄获取HWINDOW句柄.

参数:
	hWnd 窗口HWND句柄.
返回:
	返回HWINDOW句柄.
*/
func XC_hWindowFromHWnd(hWnd HWND) HWINDOW {
	ret, _, _ := xC_hWindowFromHWnd.Call(uintptr(hWnd))

	return HWINDOW(ret)
}

/*
判断元素是否从滚动视图元素扩展的新元素,包含滚动视图元素.

参数:
	hEle 元素句柄.
返回:
	如果是返回TRUE,否则相反.
*/
func XC_IsSViewExtend(hEle HELE) bool {
	ret, _, _ := xC_IsSViewExtend.Call(uintptr(hEle))

	return ret == TRUE
}

/*
获取句柄类型.

参数:
	hXCGUI 炫彩对象句柄.
返回:
	返回句柄类型.
*/
func XC_GetObjectType(hXCGUI HXCGUI) XC_OBJECT_TYPE {
	ret, _, _ := xC_GetObjectType.Call(uintptr(hXCGUI))

	return XC_OBJECT_TYPE(ret)
}

/*
通过ID获取对象句柄,不包括窗口对象.

参数:
	hWindow 所属窗口句柄,如果不属于任何窗口填NULL.
	nID ID值.
返回:
	成功返回句柄,否则返回NULL.
*/
func XC_GetObjectByID(hWindow HWINDOW, nID int) HXCGUI {
	ret, _, _ := xC_GetObjectByID.Call(
		uintptr(hWindow),
		uintptr(nID))

	if ret == NULL {
		panic("XCGetObjectByID get HXCGUI failed.")
	}

	return HXCGUI(ret)
}

//HWINDOW hWindow,
func XC_PostMessage(hWindow HWINDOW, msg uint32, wParam uintptr, lParam uintptr) bool {
	ret, _, _ := xC_PostMessage.Call(
		uintptr(hWindow),
		uintptr(msg), wParam, lParam)
	return ret != NULL
}

func XC_SendMessage(hWindow HWINDOW, msg uint32, wParam uintptr, lParam uintptr) bool {
	ret, _, _ := xC_SendMessage.Call(
		uintptr(hWindow),
		uintptr(msg), wParam, lParam)
	return ret != NULL
}

/*
获取资源ID整型值.

参数:
	pName 资源ID名称.
返回:
	返回整型值. 注解:int nID=XC_GetResIDValue(L"ID_BUTTON_1");
*/
//func XC_GetResIDValue(pName *uint16) int {
//	ret, _, _ := xC_GetResIDValue.Call(uintptr(unsafe.Pointer(pName)))

//	return int(ret)
//}

/*
设置UI的最小重绘频率.

参数:
	nMilliseconds 重绘最小时间间隔,单位毫秒.
*/
func XC_SetPaintFrequency(nMilliseconds int) {
	xC_SetPaintFrequency.Call(uintptr(nMilliseconds))
}

/*
判断两个矩形是否相交及重叠.

参数:
	pRect1 矩形1.
	pRect2 矩形2.
返回:
	如果两个矩形相交返回TRUE,否则相反.
*/
func XC_RectInRect(pRect1, pRect2 *RECT) bool {
	ret, _, _ := xC_RectInRect.Call(
		uintptr(unsafe.Pointer(pRect1)),
		uintptr(unsafe.Pointer(pRect2)))

	return ret == TRUE
}

/*
组合两个矩形区域.

参数:
	pDest 新的矩形区域.
	pSrc1 源矩形1.
	pSrc2 源矩形2.
*/
func XC_CombineRect(pDest, pSrc1, pSrc2 *RECT) {
	xC_CombineRect.Call(
		uintptr(unsafe.Pointer(pDest)),
		uintptr(unsafe.Pointer(pSrc1)),
		uintptr(unsafe.Pointer(pSrc2)))
}

/*
显示布局对象边界.

参数:
	bShow 是否显示.
*/
func XC_ShowLayoutFrame(bShow bool) {
	xC_ShowLayoutFrame.Call(uintptr(BoolToBOOL(bShow)))
}

/*
是否生成xcgui_debug.txt.

参数:
	bEnable 是否生成.
*/
func XC_EnableDebugFile(bEnable bool) {
	xC_EnableDebugFile.Call(uintptr(BoolToBOOL(bEnable)))
}

/*
设置布局边框颜色.

参数:
	color RGB颜色值.
*/
func XC_SetLayoutFrameColor(color COLORREF) {
	xC_SetLayoutFrameColor.Call(uintptr(color))
}

/*
启用错误弹出,通过该接口可以设置遇到严重错误时不弹出消息提示框.

参数:
	bEnabel 是否启用.
*/
func XC_EnableErrorMessageBox(bEnabel bool) {
	xC_EnableErrorMessageBox.Call(uintptr(BoolToBOOL(bEnabel)))
}

// 运行消息循环,当炫彩窗口数量为0时退出.
func XRunXCGUI() {
	xRunXCGUI.Call()
}

// 退出界面库释放资源.
func XExitXCGUI() {
	xExitXCGUI.Call()
	xcDLL.Release()
}

//设置文本渲染质量.
func XC_SetTextRenderingHint(nType int) {
	xC_SetTextRenderingHint.Call(uintptr(nType))
}

func XDraw_EnableSmoothingMode(hDraw HDRAW, bEnable bool) {
	xDraw_EnableSmoothingMode.Call(uintptr(hDraw), uintptr(BoolToBOOL(bEnable)))
}

func XDraw_EnableWndTransparent(hDraw HDRAW, bEnable bool) {
	xDraw_EnableWndTransparent.Call(uintptr(hDraw), uintptr(BoolToBOOL(bEnable)))
}

func FullPath(path string) (p string) {
	p, _ = filepath.Abs(path)
	return
}

func FileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	} else {
		return true
	}
}

func CallBack(pFun interface{}) uintptr {
	return syscall.NewCallback(pFun)
}

// func CallBackGo(pFunc func()) uintptr {
// 	var pfunc = func() int {
// 		pFunc()

// 		return 0
// 	}

// 	return syscall.NewCallback(pfunc)
// }

// BOOL值转换
func BoolToBOOL(value bool) BOOL {
	if value {
		return 1
	}

	return 0
}

// void * (uintptr)转换为字符串
func UINTptrToString(uintPtr uintptr) string {
	if uintPtr == 0 {
		return ""
	}

	return syscall.UTF16ToString((*[1 << 16]uint16)(unsafe.Pointer(uintPtr))[0:])
}

func UTF8PtrToSting(uintPtr uintptr) string {
	if uintPtr == 0 {
		return ""
	}

	tmpByte := (*[1 << 16]byte)(unsafe.Pointer(uintPtr))[0:]
	for i, v := range tmpByte {
		if v == 0 {
			tmpByte = tmpByte[0:i]
			break
		}
	}

	return UTF8Decode(tmpByte)
}

func UTF8Decode(b []byte) (str string) {
	for len(b) > 0 {
		s, size := utf8.DecodeRune(b)
		str += string(s)
		b = b[size:]
	}

	return
}

func UTF16PtrToString(s *uint16) string {
	if s == nil {
		return ""
	}

	return syscall.UTF16ToString((*[1 << 29]uint16)(unsafe.Pointer(s))[0:])
}

// 字符串转换
func StringToUintPtr(str string) uintptr {
	// fmt.Println("正在转换..", str)
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(str)))
}

func StringToUTF16Ptr(str string) *uint16 {
	return syscall.StringToUTF16Ptr(str)
}

func StringToBytePtr(str string) *byte {
	return syscall.StringBytePtr(str)
}
