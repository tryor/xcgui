package xc

import (
	"syscall"
	"unsafe"
)

var (
	// Functions
	xBtn_Create            *syscall.Proc
	xBtn_IsCheck           *syscall.Proc
	xBtn_SetCheck          *syscall.Proc
	xBtn_SetStyle          *syscall.Proc
	xBtn_SetState          *syscall.Proc
	xBtn_GetState          *syscall.Proc
	xBtn_GetStateEx        *syscall.Proc
	xBtn_GetStyle          *syscall.Proc
	xBtn_SetType           *syscall.Proc
	xBtn_GetType           *syscall.Proc
	xBtn_SetGroupID        *syscall.Proc
	xBtn_GetGroupID        *syscall.Proc
	xBtn_SetBindEle        *syscall.Proc
	xBtn_GetBindEle        *syscall.Proc
	xBtn_SetTextAlign      *syscall.Proc
	xBtn_GetTextAlign      *syscall.Proc
	xBtn_SetIconAlign      *syscall.Proc
	xBtn_SetOffset         *syscall.Proc
	xBtn_SetOffsetIcon     *syscall.Proc
	xBtn_SetIconSpace      *syscall.Proc
	xBtn_GetText           *syscall.Proc
	xBtn_SetText           *syscall.Proc
	xBtn_SetIcon           *syscall.Proc
	xBtn_SetIconDisable    *syscall.Proc
	xBtn_AddAnimationFrame *syscall.Proc
	xBtn_EnableAnimation   *syscall.Proc
	xBtn_AddBkBorder       *syscall.Proc
	xBtn_AddBkFill         *syscall.Proc
	xBtn_AddBkImage        *syscall.Proc
	xBtn_GetBkInfoCount    *syscall.Proc
	xBtn_ClearBkInfo       *syscall.Proc
	// xBtn_GetBkInfoManager  *syscall.Proc

	xBtn_SetTypeEx *syscall.Proc
)

func init() {
	xBtn_Create = xcDLL.MustFindProc("XBtn_Create")
	xBtn_IsCheck = xcDLL.MustFindProc("XBtn_IsCheck")
	xBtn_SetCheck = xcDLL.MustFindProc("XBtn_SetCheck")
	xBtn_SetStyle = xcDLL.MustFindProc("XBtn_SetStyle")
	xBtn_SetState = xcDLL.MustFindProc("XBtn_SetState")
	xBtn_GetState = xcDLL.MustFindProc("XBtn_GetState")
	xBtn_GetStateEx = xcDLL.MustFindProc("XBtn_GetStateEx")
	xBtn_GetStyle = xcDLL.MustFindProc("XBtn_GetStyle")
	xBtn_SetType = xcDLL.MustFindProc("XBtn_SetType")
	xBtn_GetType = xcDLL.MustFindProc("XBtn_GetType")
	xBtn_SetGroupID = xcDLL.MustFindProc("XBtn_SetGroupID")
	xBtn_GetGroupID = xcDLL.MustFindProc("XBtn_GetGroupID")
	xBtn_SetBindEle = xcDLL.MustFindProc("XBtn_SetBindEle")
	xBtn_GetBindEle = xcDLL.MustFindProc("XBtn_GetBindEle")
	xBtn_SetTextAlign = xcDLL.MustFindProc("XBtn_SetTextAlign")
	xBtn_GetTextAlign = xcDLL.MustFindProc("XBtn_GetTextAlign")
	xBtn_SetIconAlign = xcDLL.MustFindProc("XBtn_SetIconAlign")
	xBtn_SetOffset = xcDLL.MustFindProc("XBtn_SetOffset")
	xBtn_SetOffsetIcon = xcDLL.MustFindProc("XBtn_SetOffsetIcon")
	xBtn_SetIconSpace = xcDLL.MustFindProc("XBtn_SetIconSpace")
	xBtn_GetText = xcDLL.MustFindProc("XBtn_GetText")
	xBtn_SetText = xcDLL.MustFindProc("XBtn_SetText")
	xBtn_SetIcon = xcDLL.MustFindProc("XBtn_SetIcon")
	xBtn_SetIconDisable = xcDLL.MustFindProc("XBtn_SetIconDisable")
	xBtn_AddAnimationFrame = xcDLL.MustFindProc("XBtn_AddAnimationFrame")
	xBtn_EnableAnimation = xcDLL.MustFindProc("XBtn_EnableAnimation")
	xBtn_AddBkBorder = xcDLL.MustFindProc("XBtn_AddBkBorder")
	xBtn_AddBkFill = xcDLL.MustFindProc("XBtn_AddBkFill")
	xBtn_AddBkImage = xcDLL.MustFindProc("XBtn_AddBkImage")
	xBtn_GetBkInfoCount = xcDLL.MustFindProc("XBtn_GetBkInfoCount")
	xBtn_ClearBkInfo = xcDLL.MustFindProc("XBtn_ClearBkInfo")
	// xBtn_GetBkInfoManager = xcDLL.MustFindProc("XBtn_GetBkInfoManager")

	xBtn_SetTypeEx = xcDLL.MustFindProc("XBtn_SetTypeEx")

}

// 关闭
func CloseBtn(hWindow HWINDOW) {
	XBtn_SetType(XBtn_Create(10, 5, 35, 20, "关闭", HXCGUI(hWindow)),
		BUTTON_TYPE_CLOSE)
}

// 最小化
func MinBtn(hWindow HWINDOW) {
	XBtn_SetType(XBtn_Create(60, 5, 45, 20, "最小化", HXCGUI(hWindow)),
		BUTTON_TYPE_MIN)
}

/*
创建按钮元素

参数:
	x 按钮x坐标
	y 按钮y坐标
	cx 宽度
	cy 高度
	pName 标题 *uint16
	hParent 父是窗口句柄或元素句柄.
返回:
	按钮元素句柄.
*/
func XBtn_Create(x, y, cx, cy int, pName string, hParent HXCGUI) HELE {
	ret, _, _ := xBtn_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		StringToUintPtr(pName),
		// uintptr(unsafe.Pointer(pName)),
		uintptr(hParent))

	return HELE(ret)
}

/*
是否选中状态.

参数:
	hEle 元素句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XBtn_IsCheck(hEle HELE) bool {
	ret, _, _ := xBtn_IsCheck.Call(uintptr(hEle))

	return ret == TRUE
}

/*
设置选中状态.

参数:
	hEle 元素句柄.
	bCheck 是否选中.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XBtn_SetCheck(hEle HELE, bCheck bool) bool {
	ret, _, _ := xBtn_SetCheck.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bCheck)))

	return ret == TRUE
}

/*
设置按钮样式.

参数:
	hEle 元素句柄.
	nStyle 样式.
*/
func XBtn_SetStyle(hEle HELE, nStyle Button_style_) {
	xBtn_SetStyle.Call(
		uintptr(hEle),
		uintptr(nStyle))
}

func XBtn_SetTypeEx(hEle HELE, nStyle Button_type_) {
	xBtn_SetTypeEx.Call(
		uintptr(hEle),
		uintptr(nStyle))
}

/*
设置按钮状态.

参数:
	hEle 元素句柄.
	nState 按钮状态见宏定义.
*/
func XBtn_SetState(hEle HELE, nState Common_state3_) {
	xBtn_SetState.Call(
		uintptr(hEle),
		uintptr(nState))
}

/*
获取按钮状态.

参数:
	hEle 元素句柄.
返回:
	返回按钮状态.
*/
func XBtn_GetState(hEle HELE) Common_state3_ {
	ret, _, _ := xBtn_GetState.Call(uintptr(hEle))

	return Common_state3_(ret)
}

/*
获取按钮状态.

参数:
	hEle 元素句柄.
返回:
	返回按钮状态.
*/
func XBtn_GetStateEx(hEle HELE) Button_state_ {
	ret, _, _ := xBtn_GetStateEx.Call(uintptr(hEle))

	return Button_state_(ret)
}

/*
获取按钮样式.

参数:
	hEle 元素句柄.
返回:
	按钮样式,参见宏定义.
*/
func XBtn_GetStyle(hEle HELE) Button_style_ {
	ret, _, _ := xBtn_GetStyle.Call(uintptr(hEle))

	return Button_style_(ret)
}

/*
设置按钮功能类型.

参数:
	hEle 元素句柄.
	nType 按钮类型,参见宏定义.
*/
func XBtn_SetType(hEle HELE, nType Button_type_) {
	xBtn_SetType.Call(
		uintptr(hEle),
		uintptr(nType))
}

/*
获取按钮功能类型.

参数:
	hEle 元素句柄.
返回:
	按钮类型,参见宏定义.
*/
func XBtn_GetType(hEle HELE) Button_type_ {
	ret, _, _ := xBtn_GetType.Call(uintptr(hEle))

	return Button_type_(ret)
}

/*
设置组ID.

参数:
	hEle 元素句柄.
	nID 组ID.
*/
func XBtn_SetGroupID(hEle HELE, nID int) {
	xBtn_SetGroupID.Call(
		uintptr(hEle),
		uintptr(nID))
}

/*
获取组ID.

参数:
	hEle 元素句柄.
	返回:组ID.
*/
func XBtn_GetGroupID(hEle HELE) int {
	ret, _, _ := xBtn_GetGroupID.Call(uintptr(hEle))

	return int(ret)
}

/*
设置绑定元素.

参数:
	hEle 元素句柄.
	hBindEle 将要绑定的元素.
*/
func XBtn_SetBindEle(hEle, hBindEle HELE) {
	xBtn_SetBindEle.Call(
		uintptr(hEle),
		uintptr(hBindEle))
}

/*
获取绑定的元素.

参数:
	hEle 元素句柄.
返回:
	绑定的元素句柄.
*/
func XBtn_GetBindEle(hEle HELE) HELE {
	ret, _, _ := xBtn_GetBindEle.Call(uintptr(hEle))

	return HELE(ret)
}

/*
设置文本对齐方式.

参数:
	hEle 元素句柄.
	nFlags 对齐方式.
*/
func XBtn_SetTextAlign(hEle HELE, nFlags int) {
	xBtn_SetTextAlign.Call(
		uintptr(hEle),
		uintptr(nFlags))
}

/*
获取文本对齐方式.

参数:
	hEle 元素句柄.
返回:
	文本对齐方式.
*/
func XBtn_GetTextAlign(hEle HELE) int {
	ret, _, _ := xBtn_GetTextAlign.Call(uintptr(hEle))

	return int(ret)
}

/*
设置图标对齐.

参数:
	hEle 元素句柄.
	align 对齐方式.
*/
func XBtn_SetIconAlign(hEle HELE, align Button_icon_align_) {
	xBtn_SetIconAlign.Call(
		uintptr(hEle),
		uintptr(align))
}

/*
设置按钮文本坐标偏移量.

参数:
	hEle 元素句柄.
	x 偏移量.
	y 偏移量.
*/
func XBtn_SetOffset(hEle HELE, x, y int) {
	xBtn_SetOffset.Call(
		uintptr(hEle),
		uintptr(x),
		uintptr(y))
}

/*
设置按钮图标坐标偏移量.

参数:
	hEle 元素句柄.
	x 偏移量.
	y 偏移量.
*/
func XBtn_SetOffsetIcon(hEle HELE, x, y int) {
	xBtn_SetOffsetIcon.Call(
		uintptr(hEle),
		uintptr(x),
		uintptr(y))
}

/*
设置图标与文本间隔大小.

参数:
	hEle 元素句柄.
	size 间隔大小.
*/
func XBtn_SetIconSpace(hEle HELE, size int) {
	xBtn_SetIconSpace.Call(
		uintptr(hEle),
		uintptr(size))
}

/*
获取文本内容.

参数:
	hEle 元素句柄.
	pOut 接收文本内容.
	nOutLen 接收缓冲区长度,字符为单位.
*/
func XBtn_GetText2(hEle HELE, pOut *uint16, nOutLen int) {
	xBtn_GetText.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pOut)),
		uintptr(nOutLen))
}

func XBtn_GetText(hEle HELE) (ret string) {
	r, _, _ := xBtn_GetText.Call(uintptr(hEle))
	if r != 0 {
		ret = UINTptrToString(r)
	}
	return
}

/*
设置文本内容.

参数:
	hEle 元素句柄.
	pName 文本内容. *uint16
*/
func XBtn_SetText(hEle HELE, pName string) {
	xBtn_SetText.Call(
		uintptr(hEle),
		StringToUintPtr(pName))
	// uintptr(unsafe.Pointer(pName)))
}

/*
设置图标.

参数:
	hEle 元素句柄.
	hImage 图片句柄.
*/
func XBtn_SetIcon(hEle HELE, hImage HIMAGE) {
	xBtn_SetIcon.Call(
		uintptr(hEle),
		uintptr(hImage))
}

/*
设置图标禁用状态.

参数:
	hEle 元素句柄.
	hImage 图片句柄.
*/
func XBtn_SetIconDisable(hEle HELE, hImage HIMAGE) {
	xBtn_SetIconDisable.Call(
		uintptr(hEle),
		uintptr(hImage))
}

/*
添加动画帧.

参数:
	hEle 元素句柄.
	hImage 图片句柄
	uElapse 图片帧延时,单位毫秒.
*/
func XBtn_AddAnimationFrame(hEle HELE, hImage HIMAGE, uElapse uint32) {
	xBtn_AddAnimationFrame.Call(
		uintptr(hEle),
		uintptr(hImage),
		uintptr(uElapse))
}

/*
开始或关闭图片动画的播放.

参数:
	hEle 元素句柄.
	bEnable 开始播放动画TRUE,关闭播放动画FALSE.
	bLoopPlay 是否循环播放.
*/
func XBtn_EnableAnimation(hEle HELE, bEnable, bLoopPlay bool) {
	xBtn_EnableAnimation.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)),
		uintptr(BoolToBOOL(bLoopPlay)))
}

/*
添加背景内容边框.

参数:
	hEle 元素句柄.
	nState 按钮状态.
	color RGB颜色.
	alpha 透明度.
	width 线宽.
*/
func XBtn_AddBkBorder(hEle HELE, nState Button_state_, color COLORREF, alpha byte, width int) {
	xBtn_AddBkBorder.Call(
		uintptr(hEle),
		uintptr(nState),
		uintptr(color),
		uintptr(alpha),
		uintptr(width))
}

/*
添加背景内容填充.

参数:
	hEle 元素句柄.
	nState 按钮状态.
	color RGB颜色.
	alpha 透明度.
*/
func XBtn_AddBkFill(hEle HELE, nState Button_state_, color COLORREF, alpha byte) {
	xBtn_AddBkFill.Call(
		uintptr(hEle),
		uintptr(nState),
		uintptr(color),
		uintptr(alpha))
}

/*
添加背景内容图片.

参数:
	hEle 元素句柄.
	nState 按钮状态.
	hImage 图片句柄.
*/
func XBtn_AddBkImage(hEle HELE, nState Button_state_, hImage HIMAGE) {
	xBtn_AddBkImage.Call(
		uintptr(hEle),
		uintptr(nState),
		uintptr(hImage))
}

/*
获取背景内容数量.

参数:
	hEle 元素句柄.
	nState 按钮状态.
返回:
	成功返回背景内容数量,否则返回XC_ID_ERROR.
*/
func XBtn_GetBkInfoCount(hEle HELE, nState Button_state_) int {
	ret, _, _ := xBtn_GetBkInfoCount.Call(
		uintptr(hEle),
		uintptr(nState))

	return int(ret)
}

/*
清空背景内容; 如果背景没有内容,将使用系统默认内容,以便保证背景正确.

参数:
	hEle 元素句柄.
	nState 按钮状态.
*/
func XBtn_ClearBkInfo(hEle HELE, nState Button_state_) {
	xBtn_ClearBkInfo.Call(
		uintptr(hEle),
		uintptr(nState))
}

/*
获取背景内容管理器. 1.8.9.6

参数:
	hEle 元素句柄.
	nState 按钮状态.
返回:
	背景内容管理器.
*/
// func XBtn_GetBkInfoManager(hEle HELE, nState Button_state_) HBKINFOM {
// 	ret, _, _ := xBtn_GetBkInfoManager.Call(
// 		uintptr(hEle),
// 		uintptr(nState))

// 	return HBKINFOM(ret)
// }
