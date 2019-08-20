package xc

import (
	"syscall"
	"unsafe"
)

var (
	// Functions
	xList_Create                       *syscall.Proc
	xList_AddColumn                    *syscall.Proc
	xList_InsertColumn                 *syscall.Proc
	xList_EnableMultiSel               *syscall.Proc
	xList_EnableDragChangeColumnWidth  *syscall.Proc
	xList_SetDrawItemBkFlags           *syscall.Proc
	xList_SetColumnWidth               *syscall.Proc
	xList_SetColumnMinWidth            *syscall.Proc
	xList_SetColumnWidthFixed          *syscall.Proc
	xList_GetColumnWidth               *syscall.Proc
	xList_GetColumnCount               *syscall.Proc
	xList_SetItemData                  *syscall.Proc
	xList_GetItemData                  *syscall.Proc
	xList_SetSelectItem                *syscall.Proc
	xList_GetSelectItem                *syscall.Proc
	xList_GetSelectItemCount           *syscall.Proc
	xList_SetSelectAll                 *syscall.Proc
	xList_GetSelectAll                 *syscall.Proc
	xList_CancelSelectItem             *syscall.Proc
	xList_CancelSelectAll              *syscall.Proc
	xList_GetHeaderHELE                *syscall.Proc
	xList_DeleteColumn                 *syscall.Proc
	xList_DeleteColumnAll              *syscall.Proc
	xList_BindAdapter                  *syscall.Proc
	xList_BindAdapterHeader            *syscall.Proc
	xList_GetAdapter                   *syscall.Proc
	xList_SetItemTemplateXML           *syscall.Proc
	xList_SetItemTemplateXMLFromString *syscall.Proc
	xList_GetTemplateObject            *syscall.Proc
	xList_GetItemIndexFromHXCGUI       *syscall.Proc
	xList_GetHeaderTemplateObject      *syscall.Proc
	xList_GetHeaderItemIndexFromHXCGUI *syscall.Proc
	xList_SetHeaderHeight              *syscall.Proc
	xList_GetHeaderHeight              *syscall.Proc
	xList_AddItemBkBorder              *syscall.Proc
	xList_AddItemBkFill                *syscall.Proc
	xList_AddItemBkImage               *syscall.Proc
	xList_GetItemBkInfoCount           *syscall.Proc
	xList_ClearItemBkInfo              *syscall.Proc
	// xList_GetItemBkInfoManager         *syscall.Proc
	xList_SetItemHeightDefault *syscall.Proc
	xList_GetItemHeightDefault *syscall.Proc
	xList_HitTest              *syscall.Proc
	xList_HitTestOffset        *syscall.Proc
	xList_RefreshData          *syscall.Proc

	xList_AddColumnText       *syscall.Proc
	xList_AddItemText         *syscall.Proc
	xList_SetItemText         *syscall.Proc
	xList_DeleteItemAll       *syscall.Proc
	xList_EnableItemBkFullRow *syscall.Proc
	xList_SetRowSpace         *syscall.Proc
	xList_DeleteColumnAll_AD  *syscall.Proc
)

func init() {
	// Functions
	xList_Create = xcDLL.MustFindProc("XList_Create")
	xList_AddColumn = xcDLL.MustFindProc("XList_AddColumn")
	xList_InsertColumn = xcDLL.MustFindProc("XList_InsertColumn")
	xList_EnableMultiSel = xcDLL.MustFindProc("XList_EnableMultiSel")
	xList_EnableDragChangeColumnWidth = xcDLL.MustFindProc("XList_EnableDragChangeColumnWidth")
	xList_SetDrawItemBkFlags = xcDLL.MustFindProc("XList_SetDrawItemBkFlags")
	xList_SetColumnWidth = xcDLL.MustFindProc("XList_SetColumnWidth")
	xList_SetColumnMinWidth = xcDLL.MustFindProc("XList_SetColumnMinWidth")
	xList_SetColumnWidthFixed = xcDLL.MustFindProc("XList_SetColumnWidthFixed")
	xList_GetColumnWidth = xcDLL.MustFindProc("XList_GetColumnWidth")
	xList_GetColumnCount = xcDLL.MustFindProc("XList_GetColumnCount")
	xList_SetItemData = xcDLL.MustFindProc("XList_SetItemData")
	xList_GetItemData = xcDLL.MustFindProc("XList_GetItemData")
	xList_SetSelectItem = xcDLL.MustFindProc("XList_SetSelectItem")
	xList_GetSelectItem = xcDLL.MustFindProc("XList_GetSelectItem")
	xList_GetSelectItemCount = xcDLL.MustFindProc("XList_GetSelectItemCount")
	xList_SetSelectAll = xcDLL.MustFindProc("XList_SetSelectAll")
	xList_GetSelectAll = xcDLL.MustFindProc("XList_GetSelectAll")
	xList_CancelSelectItem = xcDLL.MustFindProc("XList_CancelSelectItem")
	xList_CancelSelectAll = xcDLL.MustFindProc("XList_CancelSelectAll")
	xList_GetHeaderHELE = xcDLL.MustFindProc("XList_GetHeaderHELE")
	xList_DeleteColumn = xcDLL.MustFindProc("XList_DeleteColumn")
	xList_DeleteColumnAll = xcDLL.MustFindProc("XList_DeleteColumnAll")
	xList_BindAdapter = xcDLL.MustFindProc("XList_BindAdapter")
	xList_BindAdapterHeader = xcDLL.MustFindProc("XList_BindAdapterHeader")
	xList_GetAdapter = xcDLL.MustFindProc("XList_GetAdapter")
	xList_SetItemTemplateXML = xcDLL.MustFindProc("XList_SetItemTemplateXML")
	xList_SetItemTemplateXMLFromString = xcDLL.MustFindProc("XList_SetItemTemplateXMLFromString")
	xList_GetTemplateObject = xcDLL.MustFindProc("XList_GetTemplateObject")
	xList_GetItemIndexFromHXCGUI = xcDLL.MustFindProc("XList_GetItemIndexFromHXCGUI")
	xList_GetHeaderTemplateObject = xcDLL.MustFindProc("XList_GetHeaderTemplateObject")
	xList_GetHeaderItemIndexFromHXCGUI = xcDLL.MustFindProc("XList_GetHeaderItemIndexFromHXCGUI")
	xList_SetHeaderHeight = xcDLL.MustFindProc("XList_SetHeaderHeight")
	xList_GetHeaderHeight = xcDLL.MustFindProc("XList_GetHeaderHeight")
	xList_AddItemBkBorder = xcDLL.MustFindProc("XList_AddItemBkBorder")
	xList_AddItemBkFill = xcDLL.MustFindProc("XList_AddItemBkFill")
	xList_AddItemBkImage = xcDLL.MustFindProc("XList_AddItemBkImage")
	xList_GetItemBkInfoCount = xcDLL.MustFindProc("XList_GetItemBkInfoCount")
	xList_ClearItemBkInfo = xcDLL.MustFindProc("XList_ClearItemBkInfo")
	// xList_GetItemBkInfoManager = xcDLL.MustFindProc("XList_GetItemBkInfoManager")
	xList_SetItemHeightDefault = xcDLL.MustFindProc("XList_SetItemHeightDefault")
	xList_GetItemHeightDefault = xcDLL.MustFindProc("XList_GetItemHeightDefault")
	xList_HitTest = xcDLL.MustFindProc("XList_HitTest")
	xList_HitTestOffset = xcDLL.MustFindProc("XList_HitTestOffset")
	xList_RefreshData = xcDLL.MustFindProc("XList_RefreshData")

	xList_AddColumnText = xcDLL.MustFindProc("XList_AddColumnText")
	xList_AddItemText = xcDLL.MustFindProc("XList_AddItemText")
	xList_SetItemText = xcDLL.MustFindProc("XList_SetItemText")
	xList_DeleteItemAll = xcDLL.MustFindProc("XList_DeleteItemAll")
	xList_EnableItemBkFullRow = xcDLL.MustFindProc("XList_EnableItemBkFullRow")
	xList_SetRowSpace = xcDLL.MustFindProc("XList_SetRowSpace")
	xList_DeleteColumnAll_AD = xcDLL.MustFindProc("XList_DeleteColumnAll_AD")

}

func XList_SetRowSpace(hEle HELE, w int) {
	xList_SetRowSpace.Call(
		uintptr(hEle),
		uintptr(w))
}

func XList_EnableItemBkFullRow(hEle HELE, bEnable bool) {
	xList_EnableItemBkFullRow.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

func XList_DeleteItemAll(hEle HELE) {
	xList_SetItemText.Call(uintptr(hEle))
}

func XList_DeleteColumnAll_AD(hEle HELE) {
	xList_DeleteColumnAll_AD.Call(uintptr(hEle))
}

//HELE hEle, int iItem, int iColumn, const wchar_t *pText
func XList_SetItemText(hEle HELE, iItem, iColumn int, text string) int {
	ret, _, _ := xList_SetItemText.Call(
		uintptr(hEle),
		uintptr(iItem),
		uintptr(iColumn),
		StringToUintPtr(text))

	return int(ret)
}

func XList_AddItemText(hEle HELE, text string) int {
	ret, _, _ := xList_AddItemText.Call(
		uintptr(hEle),
		StringToUintPtr(text))

	return int(ret)
}

func XList_AddColumnText(hEle HELE, w int, text string) int {
	ret, _, _ := xList_AddColumnText.Call(
		uintptr(hEle),
		uintptr(w),
		StringToUintPtr(text))

	return int(ret)
}

/*
创建列表元素.

参数:
	x 元素x坐标.
	y 元素y坐标.
	cx 宽度.
	cy 高度.
	hParent 父是窗口资源句柄或UI元素资源句柄.如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.
返回:
	元素句柄.
*/
func XList_Create(x, y, cx, cy int, hParent HXCGUI) HELE {
	ret, _, _ := xList_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))

	return HELE(ret)
}

/*
增加列.

参数:
	hEle 元素句柄.
	width 列宽度.
返回:
	返回位置索引.
*/
func XList_AddColumn(hEle HELE, width int) int {
	ret, _, _ := xList_AddColumn.Call(
		uintptr(hEle),
		uintptr(width))

	return int(ret)
}

/*
插入列.

参数:
	hEle 元素句柄.
	width 列宽度.
	iItem 插入位置索引.
返回:
	返回插入位置索引.
*/
func XList_InsertColumn(hEle HELE, width, iItem int) int {
	ret, _, _ := xList_InsertColumn.Call(
		uintptr(hEle),
		uintptr(width),
		uintptr(iItem))

	return int(ret)
}

/*
启用或关闭多选功能.

参数:
	hEle 元素句柄.
	bEnable 是否启用.
*/
func XList_EnableMultiSel(hEle HELE, bEnable bool) {
	xList_EnableMultiSel.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

/*
启用拖动改变列宽度.

参数:
	hEle 元素句柄.
	bEnable 是否启用.
*/
func XList_EnableDragChangeColumnWidth(hEle HELE, bEnable bool) {
	xList_EnableDragChangeColumnWidth.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

/*
设置是否绘制指定状态下项的背景.

参数:
	hEle 元素句柄.
	nFlags 标志位 list_drawItemBk_flags_.
*/
func XList_SetDrawItemBkFlags(hEle HELE, nFlags List_drawItemBk_flags_) {
	xList_SetDrawItemBkFlags.Call(
		uintptr(hEle),
		uintptr(nFlags))
}

/*
设置列宽度.

参数:
	hEle 元素句柄.
	iItem 列索引.
	width 宽度.
*/
func XList_SetColumnWidth(hEle HELE, iItem, width int) {
	xList_SetColumnWidth.Call(
		uintptr(hEle),
		uintptr(iItem),
		uintptr(width))
}

/*
设置列最小宽度.

参数:
	hEle 元素句柄.
	iItem 列索引.
	width 宽度.
*/
func XList_SetColumnMinWidth(hEle HELE, iItem, width int) {
	xList_SetColumnMinWidth.Call(
		uintptr(hEle),
		uintptr(iItem),
		uintptr(width))
}

/*
设置指定列宽度固定.

参数:
	hEle 元素句柄.
	iColumn 列索引.
	bFixed 是否固定宽度.
*/
func XList_SetColumnWidthFixed(hEle HELE, iColumn int, bFixed bool) {
	xList_SetColumnWidthFixed.Call(
		uintptr(hEle),
		uintptr(iColumn),
		uintptr(BoolToBOOL(bFixed)))
}

/*
获取指定列宽度.

参数:
	hEle 元素句柄.
	iColumn 列索引.
返回:
	返回指定列宽度.
*/
func XList_GetColumnWidth(hEle HELE, iColumn int) int {
	ret, _, _ := xList_GetColumnWidth.Call(
		uintptr(hEle),
		uintptr(iColumn))

	return int(ret)
}

/*
获取列数量.

参数:
	hEle 元素句柄.
返回:
	返回列数量.
*/
func XList_GetColumnCount(hEle HELE) int {
	ret, _, _ := xList_GetColumnCount.Call(uintptr(hEle))

	return int(ret)
}

/*
设置项用户数据.

参数:
	hEle 元素句柄.
	iItem 项索引.
	iSubItem 子项索引.
	data 用户数据.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XList_SetItemData(hEle HELE, iItem, iSubItem, data int) bool {
	ret, _, _ := xList_SetItemData.Call(
		uintptr(hEle),
		uintptr(iItem),
		uintptr(iSubItem),
		uintptr(data))

	return ret == TRUE
}

/*
获取项用户数据.

参数:
	hEle 元素句柄.
	iItem 项索引.
	iSubItem 子项索引.
返回:
	返回用户数据.
*/
func XList_GetItemData(hEle HELE, iItem, iSubItem int) int {
	ret, _, _ := xList_GetItemData.Call(
		uintptr(hEle),
		uintptr(iItem),
		uintptr(iSubItem))

	return int(ret)
}

/*
设置选择项.

参数:
	hEle 元素句柄.
	iItem 项索引.
	bSelect 是否选择.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XList_SetSelectItem(hEle HELE, iItem int, bSelect bool) bool {
	ret, _, _ := xList_SetSelectItem.Call(
		uintptr(hEle),
		uintptr(iItem),
		uintptr(BoolToBOOL(bSelect)))

	return ret == TRUE
}

/*
获取选择项.

参数:
	hEle 元素句柄.
返回:
	项索引.
*/
func XList_GetSelectItem(hEle HELE) int {
	ret, _, _ := xList_GetSelectItem.Call(uintptr(hEle))

	return int(ret)
}

/*
获取选择项数量.

参数:
	hEle 元素句柄.
返回:
	返回选择项数量.
*/
func XList_GetSelectItemCount(hEle HELE) int {
	ret, _, _ := xList_GetSelectItemCount.Call(uintptr(hEle))

	return int(ret)
}

/*
选择全部行.

参数:
	hEle 元素句柄.
*/
func XList_SetSelectAll(hEle HELE) {
	xList_SetSelectAll.Call(uintptr(hEle))
}

/*
获取全部选择的行.

参数:
	hEle 元素句柄.
	pArray 接收行索引数组.
	nArraySize 数组大小.
返回:
	返回行数量.
*/
func XList_GetSelectAll(hEle HELE, pArray *int, nArraySize int) int {
	ret, _, _ := xList_GetSelectAll.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pArray)),
		uintptr(nArraySize))

	return int(ret)
}

/*
取消选择指定项(这里的项可以理解为行).

参数:
	hEle 元素句柄.
	iItem 项索引.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XList_CancelSelectItem(hEle HELE, iItem int) bool {
	ret, _, _ := xList_CancelSelectItem.Call(
		uintptr(hEle),
		uintptr(iItem))

	return ret == TRUE
}

/*
取消选择所有项(这里的项可以理解为行).

参数:
	hEle 元素句柄.
*/
func XList_CancelSelectAll(hEle HELE) {
	xList_CancelSelectAll.Call(uintptr(hEle))
}

/*
获取列表头元素.

参数:
	hEle 元素句柄.
返回:
	返回列表头元素句柄.
*/
func XList_GetHeaderHELE(hEle HELE) HELE {
	ret, _, _ := xList_GetHeaderHELE.Call(uintptr(hEle))

	return HELE(ret)
}

/*
删除列.

参数:
	hEle 元素句柄.
	iItem 项索引.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XList_DeleteColumn(hEle HELE, iItem int) bool {
	ret, _, _ := xList_DeleteColumn.Call(
		uintptr(hEle),
		uintptr(iItem))

	return ret == TRUE
}

/*
删除所有的列.

参数:
	hEle 元素句柄.
*/
func XList_DeleteColumnAll(hEle HELE) {
	xList_DeleteColumnAll.Call(uintptr(hEle))
}

/*
绑定数据适配器.

参数:
	hEle 元素句柄.
	hAdapter 数据适配器句柄 XAdapterTable.
*/
func XList_BindAdapter(hEle HELE, hAdapter HXCGUI) {
	xList_BindAdapter.Call(
		uintptr(hEle),
		uintptr(hAdapter))
}

/*
列表头绑定数据适配器.

参数:
	hEle 元素句柄.
	hAdapter 数据适配器句柄 XAdapterMap.
*/
func XList_BindAdapterHeader(hEle HELE, hAdapter HXCGUI) {
	xList_BindAdapterHeader.Call(
		uintptr(hEle),
		uintptr(hAdapter))
}

/*
获取数据适配器.

参数:
	hEle 元素句柄.
返回:
	数据适配器句柄.
*/
func XList_GetAdapter(hEle HELE) HELE {
	ret, _, _ := xList_GetAdapter.Call(uintptr(hEle))

	return HELE(ret)
}

/*
设置项布局模板文件.

参数:
	hEle 元素句柄.
	pXmlFile 文件名.*uint16
返回:
	成功返回TRUE否则返回FALSE.
*/
func XList_SetItemTemplateXML(hEle HELE, pXmlFile string) bool {
	ret, _, _ := xList_SetItemTemplateXML.Call(
		uintptr(hEle),
		StringToUintPtr(pXmlFile))
	// uintptr(unsafe.Pointer(pXmlFile)))

	return ret == TRUE
}

/*
设置项布局模板文件.
参数:
hEle 元素句柄.
pStringXML 字符串指针.*uint16
返回:成功返回TRUE否则返回FALSE.
*/
func XList_SetItemTemplateXMLFromString(hEle HELE, pStringXML string) bool {
	ret, _, _ := xList_SetItemTemplateXMLFromString.Call(
		uintptr(hEle),
		//StringToUintPtr(pStringXML))
		uintptr(unsafe.Pointer(&[]byte(pStringXML)[0])))

	return ret == TRUE
}

/*
通过模板项ID,获取实例化模板项ID对应的对象句柄.

参数:
	hEle 元素句柄.
	iItem 项索引.
	nTempItemID 模板项ID.
返回:
	成功返回对象句柄,否则返回NULL.
*/
func XList_GetTemplateObject(hEle HELE, iItem, nTempItemID int) HXCGUI {
	ret, _, _ := xList_GetTemplateObject.Call(
		uintptr(hEle),
		uintptr(iItem),
		uintptr(nTempItemID))

	return HXCGUI(ret)
}

/*
获取当前对象所在模板实例,属于列表中哪一个项.

参数:
	hEle 元素句柄.
	hXCGUI 对象句柄, UI元素句柄或形状对象句柄.
返回:
	成功返回项索引, 否则返回XC_ID_ERROR.
*/
func XList_GetItemIndexFromHXCGUI(hEle HELE, hXCGUI HXCGUI) int {
	ret, _, _ := xList_GetItemIndexFromHXCGUI.Call(
		uintptr(hEle),
		uintptr(hXCGUI))

	return int(ret)
}

/*
列表头,通过模板项ID,获取实例化模板项ID对应的对象句柄.

参数:
	hEle 元素句柄.
	iItem 列表头项ID.
	nTempItemID 模板项ID.
返回:
	成功返回对象句柄,否则返回NULL.
*/
func XList_GetHeaderTemplateObject(hEle HELE, iItem, nTempItemID int) HXCGUI {
	ret, _, _ := xList_GetHeaderTemplateObject.Call(
		uintptr(hEle),
		uintptr(iItem),
		uintptr(nTempItemID))

	return HXCGUI(ret)
}

/*
列表头,获取当前对象所在模板实例,属于列表头中哪一个项.

参数:
	hEle 元素句柄.
	hXCGUI 对象句柄.
返回:
	成功返回项索引, 否则返回XC_ID_ERROR.
*/
func XList_GetHeaderItemIndexFromHXCGUI(hEle HELE, hXCGUI HXCGUI) int {
	ret, _, _ := xList_GetHeaderItemIndexFromHXCGUI.Call(
		uintptr(hEle),
		uintptr(hXCGUI))

	return int(ret)
}

/*
设置列表头高度.

参数:
	hEle 元素句柄.
	height 高度.
*/
func XList_SetHeaderHeight(hEle HELE, height int) {
	xList_SetHeaderHeight.Call(
		uintptr(hEle),
		uintptr(height))
}

/*
获取列表头高度.

参数:
	hEle 元素句柄.
返回:
	返回列表头高度.
*/
func XList_GetHeaderHeight(hEle HELE) int {
	ret, _, _ := xList_GetHeaderHeight.Call(uintptr(hEle))

	return int(ret)
}

/*
添加项背景内容边框.

参数:
	hEle 元素句柄.
	nState 项状态.
	color RGB颜色.
	alpha 透明度.
	width 线宽.
*/
func XList_AddItemBkBorder(hEle HELE, nState List_item_state_, color COLORREF, alpha byte, width int) {
	xList_AddItemBkBorder.Call(
		uintptr(hEle),
		uintptr(nState),
		uintptr(color),
		uintptr(alpha),
		uintptr(width))
}

/*
添加项背景内容填充.

参数:
	hEle 元素句柄.
	nState 项状态.
	color RGB颜色.
	alpha 透明度.
*/
func XList_AddItemBkFill(hEle HELE, nState List_item_state_, color COLORREF, alpha byte) {
	xList_AddItemBkFill.Call(
		uintptr(hEle),
		uintptr(nState),
		uintptr(color),
		uintptr(alpha))
}

/*
添加项背景内容图片.

参数:
	hEle 元素句柄.
	nState 项状态.
	hImage 图片句柄.
*/
func XList_AddItemBkImage(hEle HELE, nState List_item_state_, hImage HIMAGE) {
	xList_AddItemBkImage.Call(
		uintptr(hEle),
		uintptr(nState),
		uintptr(hImage))
}

/*
获取背景内容数量.

参数:
	hEle 元素句柄.
	nState 项状态.
返回:
	成功返回背景内容数量,否则返回XC_ID_ERROR.
*/
func XList_GetItemBkInfoCount(hEle HELE, nState List_item_state_) int {
	ret, _, _ := xList_GetItemBkInfoCount.Call(
		uintptr(hEle),
		uintptr(nState))

	return int(ret)

}

/*
清空项背景内容; 如果背景没有内容,将使用系统默认内容,以便保证背景正确.

参数:
	hEle 元素句柄.
	nState 项状态.
*/
func XList_ClearItemBkInfo(hEle HELE, nState List_item_state_) {
	xList_ClearItemBkInfo.Call(
		uintptr(hEle),
		uintptr(nState))

}

/*
获取项背景内容管理器. 1.8.9.6

参数:
	hEle 元素句柄.
	nState 项状态.
返回:
	项背景内容管理器.
*/
// func XList_GetItemBkInfoManager(hEle HELE, nState List_item_state_) HBKINFOM {
// 	ret, _, _ := xList_GetItemBkInfoManager.Call(
// 		uintptr(hEle),
// 		uintptr(nState))

// 	return HBKINFOM(ret)
// }

/*
设置项默认高度.

参数:
	hEle 元素句柄.
	nHeight 高度.
	nSelHeight 选中时高度.
*/
func XList_SetItemHeightDefault(hEle HELE, nHeight, nSelHeight int) {
	xList_SetItemHeightDefault.Call(
		uintptr(hEle),
		uintptr(nHeight),
		uintptr(nSelHeight))
}

/*
获取项默认高度.

参数:
	hEle 元素句柄.
	pHeight 高度.
	pSelHeight 选中时高度.
*/
func XList_GetItemHeightDefault(hEle HELE, pHeight, pSelHeight *int32) {
	xList_GetItemHeightDefault.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pHeight)),
		uintptr(unsafe.Pointer(pSelHeight)))
}

/*
检测坐标点所在项.

参数:
	hEle 元素句柄.
	pPt 坐标点.
	piItem 项索引.
	piSubItem 子项索引.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XList_HitTest(hEle HELE, pPt *POINT, piItem, piSubItem *int32) bool {
	ret, _, _ := xList_HitTest.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pPt)),
		uintptr(unsafe.Pointer(piItem)),
		uintptr(unsafe.Pointer(piSubItem)))

	return ret == TRUE
}

/*
检查坐标点所在项,自动添加滚动视图偏移量.

参数:
	hEle 元素句柄.
	pPt 坐标点.
	piItem 项索引.
	piSubItem 子项索引.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XList_HitTestOffset(hEle HELE, pPt *POINT, piItem, piSubItem *int32) bool {
	ret, _, _ := xList_HitTestOffset.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pPt)),
		uintptr(unsafe.Pointer(piItem)),
		uintptr(unsafe.Pointer(piSubItem)))

	return ret == TRUE
}

/*
刷新数据.

参数:
	hEle 元素句柄.
*/
func XList_RefreshData(hEle HELE) {
	xList_RefreshData.Call(uintptr(hEle))
}
