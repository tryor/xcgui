package xc

import (
	"syscall"
	//	"unsafe"
)

var (
	// Functions
	xLayout_Create *syscall.Proc
	//	xLayout_Destroy         *syscall.Proc
	//	xLayout_AddEle          *syscall.Proc
	//	xLayout_AddLayoutObject *syscall.Proc
	//	xLayout_AddShape        *syscall.Proc
	//	xLayout_RemoveChild     *syscall.Proc
	//	xLayout_AdjustLayout    *syscall.Proc
	//	xLayout_GetEle          *syscall.Proc
	//	xLayout_GetParentLayout *syscall.Proc
	//	xLayout_GetParent       *syscall.Proc
	//	xLayout_GetWindow       *syscall.Proc
	//	xLayout_SetID           *syscall.Proc
	//	xLayout_GetID           *syscall.Proc
	xLayout_SetHorizon *syscall.Proc
	//	xLayout_SetRectFrame    *syscall.Proc
	//	xLayout_GetRect *syscall.Proc
	//	xLayout_GetRectFrame    *syscall.Proc
	xLayout_SetAlignH  *syscall.Proc
	xLayout_SetAlignV  *syscall.Proc
	xLayout_SetPadding *syscall.Proc
	//	xLayout_SetLayoutWidth  *syscall.Proc
	//	xLayout_SetLayoutHeight *syscall.Proc
	//	xLayout_GetLayoutWidth  *syscall.Proc
	//	xLayout_GetLayoutHeight *syscall.Proc
	xLayout_SetSpace *syscall.Proc
	//	xLayout_GetWidth       *syscall.Proc
	//	xLayout_GetHeight      *syscall.Proc
	xLayout_GetWidthIn  *syscall.Proc
	xLayout_GetHeightIn *syscall.Proc
	//	xLayout_GetContentSize *syscall.Proc
	//	xLayout_ShowLayout    *syscall.Proc
	//	xLayout_GetChildCount *syscall.Proc
	//	xLayout_GetChildType *syscall.Proc
	//	xLayout_GetChild *syscall.Proc
//	xLayout_Draw *syscall.Proc
)

func init() {
	// Functions
	xLayout_Create = xcDLL.MustFindProc("XLayout_Create")
	//	xLayout_Destroy = xcDLL.MustFindProc("XLayout_Destroy")
	//	xLayout_AddEle = xcDLL.MustFindProc("XLayout_AddEle")
	//	xLayout_AddLayoutObject = xcDLL.MustFindProc("XLayout_AddLayoutObject")
	//	xLayout_AddShape = xcDLL.MustFindProc("XLayout_AddShape")
	//	xLayout_RemoveChild = xcDLL.MustFindProc("XLayout_RemoveChild")
	//	xLayout_AdjustLayout = xcDLL.MustFindProc("XLayout_AdjustLayout")
	//	xLayout_GetEle = xcDLL.MustFindProc("XLayout_GetEle")
	//	xLayout_GetParentLayout = xcDLL.MustFindProc("XLayout_GetParentLayout")
	//	xLayout_GetParent = xcDLL.MustFindProc("XLayout_GetParent")
	//	xLayout_GetWindow = xcDLL.MustFindProc("XLayout_GetWindow")
	//	xLayout_SetID = xcDLL.MustFindProc("XLayout_SetID")
	//	xLayout_GetID = xcDLL.MustFindProc("XLayout_GetID")
	xLayout_SetHorizon = xcDLL.MustFindProc("XLayout_SetHorizon")
	//	xLayout_SetRectFrame = xcDLL.MustFindProc("XLayout_SetRectFrame")
	//	xLayout_GetRect = xcDLL.MustFindProc("XLayout_GetRect")
	//	xLayout_GetRectFrame = xcDLL.MustFindProc("XLayout_GetRectFrame")
	xLayout_SetAlignH = xcDLL.MustFindProc("XLayout_SetAlignH")
	xLayout_SetAlignV = xcDLL.MustFindProc("XLayout_SetAlignV")
	xLayout_SetPadding = xcDLL.MustFindProc("XLayout_SetPadding")
	//	xLayout_SetLayoutWidth = xcDLL.MustFindProc("XLayout_SetLayoutWidth")
	//	xLayout_SetLayoutHeight = xcDLL.MustFindProc("XLayout_SetLayoutHeight")
	//	xLayout_GetLayoutWidth = xcDLL.MustFindProc("XLayout_GetLayoutWidth")
	//	xLayout_GetLayoutHeight = xcDLL.MustFindProc("XLayout_GetLayoutHeight")
	xLayout_SetSpace = xcDLL.MustFindProc("XLayout_SetSpace")
	//	xLayout_GetWidth = xcDLL.MustFindProc("XLayout_GetWidth")
	//	xLayout_GetHeight = xcDLL.MustFindProc("XLayout_GetHeight")
	xLayout_GetWidthIn = xcDLL.MustFindProc("XLayout_GetWidthIn")
	xLayout_GetHeightIn = xcDLL.MustFindProc("XLayout_GetHeightIn")
	//	xLayout_GetContentSize = xcDLL.MustFindProc("XLayout_GetContentSize")
	//	xLayout_ShowLayout = xcDLL.MustFindProc("XLayout_ShowLayout")
	//	xLayout_GetChildCount = xcDLL.MustFindProc("XLayout_GetChildCount")
	//	xLayout_GetChildType = xcDLL.MustFindProc("XLayout_GetChildType")
	//	xLayout_GetChild = xcDLL.MustFindProc("XLayout_GetChild")
	//	xLayout_Draw = xcDLL.MustFindProc("XLayout_Draw")

}

/*
创建布局对象.

返回:
	返回布局对象句柄.
*/
func XLayout_Create() HXCGUI {
	ret, _, _ := xLayout_Create.Call()

	return HXCGUI(ret)
}

/*
销毁布局对象.

参数:
	hLayout 布局对象句柄.
*/
//func XLayout_Destroy(hLayout HXCGUI) {
//	xLayout_Destroy.Call(uintptr(hLayout))
//}

/*
添加元素到布局对象,自动将元素添加到父UI中.

参数:
	hLayout 布局对象句柄.
	hEle 元素句柄.
*/
//func XLayout_AddEle(hLayout HXCGUI, hEle HELE) {
//	xLayout_AddEle.Call(
//		uintptr(hLayout),
//		uintptr(hEle))
//}

/*
添加布局对象到当前布局对象.

参数:
	hLayout 布局对象句柄.
	hLayoutObject 布局对象句柄.
*/
//func XLayout_AddLayoutObject(hLayout, hLayoutObject HXCGUI) {
//	xLayout_AddLayoutObject.Call(
//		uintptr(hLayout),
//		uintptr(hLayoutObject))
//}

/*
添加形状对象到当前布局对象,自动将形状对象添加到父UI中.

参数:
	hLayout 布局对象句柄.
	hShape 文本块对象.
*/
//func XLayout_AddShape(hLayout HXCGUI, hShape HXCGUI) {
//	xLayout_AddShape.Call(
//		uintptr(hLayout),
//		uintptr(hShape))
//}

/*
移除子对象.

参数:
	hLayout 布局对象句柄.
	hChild 子对象句柄.
*/
//func XLayout_RemoveChild(hLayout, hChild HXCGUI) {
//	xLayout_RemoveChild.Call(
//		uintptr(hLayout),
//		uintptr(hChild))
//}

/*
调整布局.

参数:
	hLayout 布局对象句柄.
*/
//func XLayout_AdjustLayout(hLayout HXCGUI) {
//	xLayout_AdjustLayout.Call(uintptr(hLayout))
//}

/*
获取布局对象所在元素.

参数:
	hLayout 布局对象句柄.
返回:
	返回元素句柄.
*/
//func XLayout_GetEle(hLayout HXCGUI) HELE {
//	ret, _, _ := xLayout_GetEle.Call(uintptr(hLayout))

//	return HELE(ret)
//}

/*
获取父布局对象.

参数:
	hLayout 布局对象句柄.
返回:
	返回父布局对象,如果没有返回空.
*/
//func XLayout_GetParentLayout(hLayout HXCGUI) HXCGUI {
//	ret, _, _ := xLayout_GetParentLayout.Call(uintptr(hLayout))

//	return HXCGUI(ret)
//}

/*
获取父对象,父可能是元素或窗口.

参数:
	hLayout 布局对象句柄.
返回:
	返回父对象.
*/
//func XLayout_GetParent(hLayout HXCGUI) HXCGUI {
//	ret, _, _ := xLayout_GetParent.Call(uintptr(hLayout))

//	return HXCGUI(ret)
//}

/*
获取布局对象所在窗口.

参数:
	hLayout 布局对象句柄.
返回:
	返回窗口句柄.
*/
//func XLayout_GetWindow(hLayout HXCGUI) HWINDOW {
//	ret, _, _ := xLayout_GetWindow.Call(uintptr(hLayout))

//	return HWINDOW(ret)
//}

/*
设置ID.

参数:
	hLayout 布局对象句柄.
	nID ID值.
*/
//func XLayout_SetID(hLayout HXCGUI, nID int) {
//	xLayout_SetID.Call(
//		uintptr(hLayout),
//		uintptr(nID))
//}

/*
获取ID.

参数:
	hLayout 布局对象句柄.
返回:
	返回ID值.
*/
//func XLayout_GetID(hLayout HXCGUI) int {
//	ret, _, _ := xLayout_GetID.Call(uintptr(hLayout))

//	return int(ret)
//}

/*
设置水平或垂直布局.

参数:
	hLayout 布局对象句柄.
	bHorizon 水平或垂直.
*/
func XLayout_SetHorizon(hLayout HXCGUI, bHorizon bool) {
	xLayout_SetHorizon.Call(
		uintptr(hLayout),
		uintptr(BoolToBOOL(bHorizon)))
}

/*
设置外围框架坐标.

参数:
	hLayout 布局对象句柄.
	pRect 坐标.
*/
//func XLayout_SetRectFrame(hLayout HXCGUI, pRect *RECT) {
//	xLayout_SetRectFrame.Call(
//		uintptr(hLayout),
//		uintptr(unsafe.Pointer(pRect)))
//}

/*
获取内围坐标.

参数:
	hLayout 布局对象句柄.
	pRect 坐标.
*/
//func XLayout_GetRect(hLayout HXCGUI, pRect *RECT) {
//	xLayout_GetRect.Call(
//		uintptr(hLayout),
//		uintptr(unsafe.Pointer(pRect)))
//}

/*
获取外围框架坐标.

参数:
	hLayout 布局对象句柄.
	pRect 接收返回坐标值.
*/
//func XLayout_GetRectFrame(hLayout HXCGUI, pRect *RECT) {
//	xLayout_GetRectFrame.Call(
//		uintptr(hLayout),
//		uintptr(unsafe.Pointer(pRect)))
//}

/*
设置水平对齐方式.

参数:
	hLayout 布局对象句柄.
	nAlign 对齐方式参见宏定义.
*/
func XLayout_SetAlignH(hLayout HXCGUI, nAlign Layout_align_) {
	xLayout_SetAlignH.Call(
		uintptr(hLayout),
		uintptr(nAlign))
}

/*
设置垂直对齐方式.

参数:
	hLayout 布局对象句柄.
	nAlign 对齐方式参见宏定义.
*/
func XLayout_SetAlignV(hLayout HXCGUI, nAlign Layout_align_) {
	xLayout_SetAlignV.Call(
		uintptr(hLayout),
		uintptr(nAlign))
}

/*
设置内填充大小.

参数:
	hLayout 布局对象句柄.
	left 左边大小.
	top 上边大小.
	right 右边大小.
	bottom 下边大小.
*/
func XLayout_SetPadding(hLayout HXCGUI, left, top, right, bottom int) {
	xLayout_SetPadding.Call(
		uintptr(hLayout),
		uintptr(left),
		uintptr(top),
		uintptr(right),
		uintptr(bottom))
}

/*
设置宽度.

参数:
	hLayout 布局对象句柄.
	nType 属性类型.
	nWidth 宽度.
*/
//func XLayout_SetLayoutWidth(hLayout HXCGUI, nType Layout_size_type_, nWidth int) {
//	xLayout_SetLayoutWidth.Call(
//		uintptr(hLayout),
//		uintptr(nType),
//		uintptr(nWidth))
//}

/*
设置高度.

参数:
	hLayout 布局对象句柄.
	nType 属性类型.
	nHeight 高度.
*/
//func XLayout_SetLayoutHeight(hLayout HXCGUI, nType Layout_size_type_, nHeight int) {
//	xLayout_SetLayoutHeight.Call(
//		uintptr(hLayout),
//		uintptr(nType),
//		uintptr(nHeight))
//}

/*
获取布局宽度.

参数:
	hLayout 布局对象句柄.
	pType 属性标识.
	pWidth 宽度.
*/
//func XLayout_GetLayoutWidth(hLayout HXCGUI, pType *Layout_size_type_, pWidth *int) {
//	xLayout_GetLayoutWidth.Call(
//		uintptr(hLayout),
//		uintptr(unsafe.Pointer(pType)),
//		uintptr(unsafe.Pointer(pWidth)))
//}

/*
获取布局高度.

参数:
	hLayout 布局对象句柄.
	pType 属性标识.
	pHeight 高度.
*/
//func XLayout_GetLayoutHeight(hLayout HXCGUI, pType *Layout_size_type_, pHeight *int) {
//	xLayout_GetLayoutHeight.Call(
//		uintptr(hLayout),
//		uintptr(unsafe.Pointer(pType)),
//		uintptr(unsafe.Pointer(pHeight)))
//}

/*
设置子对象之间的间距.

参数:
	hLayout 布局对象句柄.
	nSpace 间距大小.
*/
func XLayout_SetSpace(hLayout HXCGUI, nSpace int) {
	xLayout_SetSpace.Call(
		uintptr(hLayout),
		uintptr(nSpace))
}

/*
获取宽度.

参数:
	hLayout 布局对象句柄.
返回:
	返回宽度.
*/
//func XLayout_GetWidth(hLayout HXCGUI) int {
//	ret, _, _ := xLayout_GetWidth.Call(uintptr(hLayout))

//	return int(ret)
//}

/*
获取高度.

参数:
	hLayout 布局对象句柄.
返回:
	返回高度.
*/
//func XLayout_GetHeight(hLayout HXCGUI) int {
//	ret, _, _ := xLayout_GetHeight.Call(uintptr(hLayout))

//	return int(ret)
//}

/*
获取宽度,不包含内边距大小.

参数:
	hLayout 布局对象句柄.
返回:
	返回宽度.
*/
func XLayout_GetWidthIn(hLayout HXCGUI) int {
	ret, _, _ := xLayout_GetWidthIn.Call(uintptr(hLayout))

	return int(ret)
}

/*
获取高度,不包含内边距大小.

参数:
	hLayout 布局对象句柄.
返回:
	返回高度.
*/
func XLayout_GetHeightIn(hLayout HXCGUI) int {
	ret, _, _ := xLayout_GetHeightIn.Call(uintptr(hLayout))

	return int(ret)
}

/*
获取内容大小.

参数:
	hLayout 布局对象句柄.
	pSize 接收返回大小值.
*/
//func XLayout_GetContentSize(hLayout HXCGUI, pSize *SIZE) {
//	xLayout_GetContentSize.Call(
//		uintptr(hLayout),
//		uintptr(unsafe.Pointer(pSize)))
//}

/*
是否显示布局对象.

参数:
	hLayout 布局对象句柄.
	bShow 是否显示.
*/
//func XLayout_ShowLayout(hLayout HXCGUI, bShow bool) {
//	xLayout_ShowLayout.Call(
//		uintptr(hLayout),
//		uintptr(BoolToBOOL(bShow)))
//}

/*
获取子对象数量.

参数:
	hLayout 布局对象句柄.
返回:
	返回子对象数量.
*/
//func XLayout_GetChildCount(hLayout HXCGUI) int {
//	ret, _, _ := xLayout_GetChildCount.Call(uintptr(hLayout))

//	return int(ret)
//}

/*
获取子对象类型.

参数:
	hLayout 布局对象句柄.
	index 索引值.
返回:
	返回类型.
*/
//func XLayout_GetChildType(hLayout HXCGUI, index int) XC_OBJECT_TYPE {
//	ret, _, _ := xLayout_GetChildType.Call(
//		uintptr(hLayout),
//		uintptr(index))

//	return XC_OBJECT_TYPE(ret)
//}

/*
获取子对象.

参数:
	hLayout 布局对象句柄.
	index 索引值.
返回:
	返回对象句柄.
*/
//func XLayout_GetChild(hLayout HXCGUI, index int) HXCGUI {
//	ret, _, _ := xLayout_GetChild.Call(
//		uintptr(hLayout),
//		uintptr(index))

//	return HXCGUI(ret)
//}

/*
绘制布局对象.

参数:
	hLayout 布局对象句柄.
	hDraw 图形绘制句柄.
*/
//func XLayout_Draw(hLayout HXCGUI, hDraw HDRAW) {
//	xLayout_Draw.Call(
//		uintptr(hLayout),
//		uintptr(hDraw))
//}
