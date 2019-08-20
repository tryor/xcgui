package xc

import (
	"syscall"
	"unsafe"
)

var (
	// Functions
	xAdTable_Create            *syscall.Proc
	xAdTable_AddColumn         *syscall.Proc
	xAdTable_SetColumn         *syscall.Proc
	xAdTable_AddItemText       *syscall.Proc
	xAdTable_AddItemTextEx     *syscall.Proc
	xAdTable_AddItemImage      *syscall.Proc
	xAdTable_AddItemImageEx    *syscall.Proc
	xAdTable_InsertItemText    *syscall.Proc
	xAdTable_InsertItemTextEx  *syscall.Proc
	xAdTable_InsertItemImage   *syscall.Proc
	xAdTable_InsertItemImageEx *syscall.Proc
	xAdTable_SetItemText       *syscall.Proc
	xAdTable_SetItemTextEx     *syscall.Proc
	xAdTable_SetItemImage      *syscall.Proc
	xAdTable_SetItemImageEx    *syscall.Proc
	xAdTable_DeleteItem        *syscall.Proc
	xAdTable_DeleteItemEx      *syscall.Proc
	xAdTable_DeleteItemAll     *syscall.Proc
	xAdTable_DeleteColumnAll   *syscall.Proc
	xAdTable_GetCount          *syscall.Proc
	xAdTable_GetCountColumn    *syscall.Proc
	xAdTable_GetItemText       *syscall.Proc
	xAdTable_GetItemImage      *syscall.Proc
	xAdTable_GetItemTextEx     *syscall.Proc
	xAdTable_GetItemImageEx    *syscall.Proc
)

func init() {
	// Functions
	xAdTable_Create = xcDLL.MustFindProc("XAdTable_Create")
	xAdTable_AddColumn = xcDLL.MustFindProc("XAdTable_AddColumn")
	xAdTable_SetColumn = xcDLL.MustFindProc("XAdTable_SetColumn")
	xAdTable_AddItemText = xcDLL.MustFindProc("XAdTable_AddItemText")
	xAdTable_AddItemTextEx = xcDLL.MustFindProc("XAdTable_AddItemTextEx")
	xAdTable_AddItemImage = xcDLL.MustFindProc("XAdTable_AddItemImage")
	xAdTable_AddItemImageEx = xcDLL.MustFindProc("XAdTable_AddItemImageEx")
	xAdTable_InsertItemText = xcDLL.MustFindProc("XAdTable_InsertItemText")
	xAdTable_InsertItemTextEx = xcDLL.MustFindProc("XAdTable_InsertItemTextEx")
	xAdTable_InsertItemImage = xcDLL.MustFindProc("XAdTable_InsertItemImage")
	xAdTable_InsertItemImageEx = xcDLL.MustFindProc("XAdTable_InsertItemImageEx")
	xAdTable_SetItemText = xcDLL.MustFindProc("XAdTable_SetItemText")
	xAdTable_SetItemTextEx = xcDLL.MustFindProc("XAdTable_SetItemTextEx")
	xAdTable_SetItemImage = xcDLL.MustFindProc("XAdTable_SetItemImage")
	xAdTable_SetItemImageEx = xcDLL.MustFindProc("XAdTable_SetItemImageEx")
	xAdTable_DeleteItem = xcDLL.MustFindProc("XAdTable_DeleteItem")
	xAdTable_DeleteItemEx = xcDLL.MustFindProc("XAdTable_DeleteItemEx")
	xAdTable_DeleteItemAll = xcDLL.MustFindProc("XAdTable_DeleteItemAll")
	xAdTable_DeleteColumnAll = xcDLL.MustFindProc("XAdTable_DeleteColumnAll")
	xAdTable_GetCount = xcDLL.MustFindProc("XAdTable_GetCount")
	xAdTable_GetCountColumn = xcDLL.MustFindProc("XAdTable_GetCountColumn")
	xAdTable_GetItemText = xcDLL.MustFindProc("XAdTable_GetItemText")
	xAdTable_GetItemImage = xcDLL.MustFindProc("XAdTable_GetItemImage")
	xAdTable_GetItemTextEx = xcDLL.MustFindProc("XAdTable_GetItemTextEx")
	xAdTable_GetItemImageEx = xcDLL.MustFindProc("XAdTable_GetItemImageEx")
}

/*
创建列表框元素数据适配器.

返回:
	返回数据适配器句柄.
*/
func XAdTable_Create() HXCGUI {
	ret, _, _ := xAdTable_Create.Call()

	return HXCGUI(ret)
}

/*
添加数据列.

参数:
	hAdapter 数据适配器句柄.
	pName 名称.*uint16
返回:
	返回列索引.
*/
func XAdTable_AddColumn(hAdapter HXCGUI, pName string) int {
	ret, _, _ := xAdTable_AddColumn.Call(
		uintptr(hAdapter),
		StringToUintPtr(pName))
	// uintptr(unsafe.Pointer(pName)))

	return int(ret)
}

/*
设置列.

参数:
	hAdapter 数据适配器句柄.
	pColName 列名,多个列名用逗号分开.*uint16
返回:
	返回列数量. 注解:例如: XAdTable_SetColumn(hAdapter, L"name1,name2,mame3");
*/
func XAdTable_SetColumn(hAdapter HXCGUI, pColName string) int {
	ret, _, _ := xAdTable_SetColumn.Call(
		uintptr(hAdapter),
		StringToUintPtr(pColName))
	// uintptr(unsafe.Pointer(pColName)))

	return int(ret)
}

/*
添加数据项,默认值放到第一列.

参数:
	hAdapter 数据适配器句柄.
	pValue 值.*uint16
返回:
	返回项索引值.
*/
func XAdTable_AddItemText(hAdapter HXCGUI, pValue string) int {
	ret, _, _ := xAdTable_AddItemText.Call(
		uintptr(hAdapter),
		StringToUintPtr(pValue))
	// uintptr(unsafe.Pointer(pValue)))

	return int(ret)
}

/*
添加数据项,填充指定列数据.

参数:
	hAdapter 数据适配器句柄.
	pName 列名.*uint16
	pValue 值.*uint16
返回:
	返回项索引.
*/
func XAdTable_AddItemTextEx(hAdapter HXCGUI, pName, pValue string) int {
	ret, _, _ := xAdTable_AddItemTextEx.Call(
		uintptr(hAdapter),
		StringToUintPtr(pName),
		StringToUintPtr(pValue))
	// uintptr(unsafe.Pointer(pName)),
	// uintptr(unsafe.Pointer(pValue)))

	return int(ret)
}

/*
添加数据项,默认值放到第一列.

参数:
	hAdapter 数据适配器句柄.
	hImage 图片句柄.
返回:
	返回项索引值.
*/
func XAdTable_AddItemImage(hAdapter HXCGUI, hImage HIMAGE) int {
	ret, _, _ := xAdTable_AddItemImage.Call(
		uintptr(hAdapter),
		uintptr(hImage))

	return int(ret)
}

/*
添加数据项,并填充指定列数据.

参数:
	hAdapter 数据适配器句柄.
	pName 列名.*uint16
	hImage 图片句柄.
返回:

*/
func XAdTable_AddItemImageEx(hAdapter HXCGUI, pName string, hImage HIMAGE) int {
	ret, _, _ := xAdTable_AddItemImageEx.Call(
		uintptr(hAdapter),
		StringToUintPtr(pName),
		// uintptr(unsafe.Pointer(pName)),
		uintptr(hImage))

	return int(ret)
}

/*
插入数据项,填充第一列数据.

参数:
	hAdapter 数据适配器句柄.
	iItem 插入位置索引.
	pValue 值.*uint16
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdTable_InsertItemText(hAdapter HXCGUI, iItem int, pValue string) bool {
	ret, _, _ := xAdTable_InsertItemText.Call(
		uintptr(hAdapter),
		uintptr(iItem),
		StringToUintPtr(pValue))
	// uintptr(unsafe.Pointer(pValue)))

	return ret == TRUE
}

/*
插入数据项,并填充指定列数据.

参数:
	hAdapter 数据适配器句柄.
	iItem 插入位置索引.
	pName 列名.
	pValue 值.*uint16
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdTable_InsertItemTextEx(hAdapter HXCGUI, iItem int, pName, pValue string) bool {
	ret, _, _ := xAdTable_InsertItemTextEx.Call(
		uintptr(hAdapter),
		uintptr(iItem),
		StringToUintPtr(pName),
		StringToUintPtr(pValue))
	// uintptr(unsafe.Pointer(pName)),
	// uintptr(unsafe.Pointer(pValue)))

	return ret == TRUE
}

/*
插入数据项,填充第一列数据.

参数:
	hAdapter 数据适配器句柄.
	iItem 插入位置索引.
	hImage 图片句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdTable_InsertItemImage(hAdapter HXCGUI, iItem int, hImage HIMAGE) bool {
	ret, _, _ := xAdTable_InsertItemImage.Call(
		uintptr(hAdapter),
		uintptr(iItem),
		uintptr(hImage))

	return ret == TRUE
}

/*
插入数据项,并填充指定列数据.

参数:
	hAdapter 数据适配器句柄.
	iItem 插入位置索引.
	pName 列名.*uint16
	hImage 图片句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdTable_InsertItemImageEx(hAdapter HXCGUI, iItem int, pName string, hImage HIMAGE) bool {
	ret, _, _ := xAdTable_InsertItemImageEx.Call(
		uintptr(hAdapter),
		uintptr(iItem),
		StringToUintPtr(pName),
		// uintptr(unsafe.Pointer(pName)),
		uintptr(hImage))

	return ret == TRUE
}

/*
设置项数据.

参数:
	hAdapter 数据适配器句柄.
	iItem 项索引.
	iColumn 列索引.
	pValue 值.*uint16
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdTable_SetItemText(hAdapter HXCGUI, iItem, iColumn int, pValue string) bool {
	ret, _, _ := xAdTable_SetItemText.Call(
		uintptr(hAdapter),
		uintptr(iItem),
		uintptr(iColumn),
		StringToUintPtr(pValue))
	// uintptr(unsafe.Pointer(pValue)))

	return ret == TRUE
}

/*
设置项数据.

参数:
	hAdapter 数据适配器句柄.
	iItem 项索引.
	pName 列名.*uint16
	pValue 值.*uint16
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdTable_SetItemTextEx(hAdapter HXCGUI, iItem int, pName, pValue string) bool {
	ret, _, _ := xAdTable_SetItemTextEx.Call(
		uintptr(hAdapter),
		uintptr(iItem),
		StringToUintPtr(pName),
		StringToUintPtr(pValue))
	// uintptr(unsafe.Pointer(pName)),
	// uintptr(unsafe.Pointer(pValue)))

	return ret == TRUE
}

/*
设置项数据.

参数:
	hAdapter 数据适配器句柄.
	iItem 项索引.
	iColumn 列索引.
	hImage 图片句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdTable_SetItemImage(hAdapter HXCGUI, iItem, iColumn int, hImage HIMAGE) bool {
	ret, _, _ := xAdTable_SetItemImage.Call(
		uintptr(hAdapter),
		uintptr(iItem),
		uintptr(iColumn),
		uintptr(hImage))

	return ret == TRUE
}

/*
设置项数据.

参数:
	hAdapter 数据适配器句柄.
	iItem 项索引.
	pName 列名.*uint16
	hImage 图片句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdTable_SetItemImageEx(hAdapter HXCGUI, iItem int, pName string, hImage HIMAGE) bool {
	ret, _, _ := xAdTable_SetItemImageEx.Call(
		uintptr(hAdapter),
		uintptr(iItem),
		StringToUintPtr(pName),
		// uintptr(unsafe.Pointer(pName)),
		uintptr(hImage))

	return ret == TRUE
}

/*
删除项.

参数:
	hAdapter 数据适配器句柄.
	iItem 项索引.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdTable_DeleteItem(hAdapter HXCGUI, iItem int) bool {
	ret, _, _ := xAdTable_DeleteItem.Call(
		uintptr(hAdapter),
		uintptr(iItem))

	return ret == TRUE
}

/*
删除多个项.

参数:
	hAdapter 数据适配器句柄.
	iItem 项起始索引.
	nCount 删除项数量.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdTable_DeleteItemEx(hAdapter HXCGUI, iItem, nCount int) bool {
	ret, _, _ := xAdTable_DeleteItemEx.Call(
		uintptr(hAdapter),
		uintptr(iItem),
		uintptr(nCount))

	return ret == TRUE
}

/*
删除所有项.

参数:
	hAdapter 数据适配器句柄.
*/
func XAdTable_DeleteItemAll(hAdapter HXCGUI) {
	xAdTable_DeleteItemAll.Call(uintptr(hAdapter))
}

/*
删除所有列,并且清空所有数据.

参数:
	hAdapter 数据适配器句柄.
*/
func XAdTable_DeleteColumnAll(hAdapter HXCGUI) {
	xAdTable_DeleteColumnAll.Call(uintptr(hAdapter))
}

/*
获取项数量.

参数:
	hAdapter 数据适配器句柄.
返回:
	返回数量.
*/
func XAdTable_GetCount(hAdapter HXCGUI) int {
	ret, _, _ := xAdTable_GetCount.Call(uintptr(hAdapter))

	return int(ret)
}

/*
获取列数量.

参数:
	hAdapter 数据适配器句柄.
返回:
	返回列数量.
*/
func XAdTable_GetCountColumn(hAdapter HXCGUI) int {
	ret, _, _ := xAdTable_GetCountColumn.Call(uintptr(hAdapter))

	return int(ret)
}

/*
获取项文本内容.

参数:
	hAdapter 数据适配器句柄.
	iItem 项索引.
	iColumn 列索引.
	pOut 接收内容缓冲区.
	nOutLen 缓冲区大小,字符为单位.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdTable_GetItemText(hAdapter HXCGUI, iItem, iColumn int, pOut *uint16, nOutLen int) bool {
	ret, _, _ := xAdTable_GetItemText.Call(
		uintptr(hAdapter),
		uintptr(iItem),
		uintptr(iColumn),
		uintptr(unsafe.Pointer(pOut)),
		uintptr(nOutLen))

	return ret == TRUE
}

/*
获取项图片句柄.

参数:
	hAdapter 数据适配器句柄.
	iItem 项索引.
	iColumn 列索引.
返回:
	返回图片句柄.
*/
func XAdTable_GetItemImage(hAdapter HXCGUI, iItem, iColumn int) HIMAGE {
	ret, _, _ := xAdTable_GetItemImage.Call(
		uintptr(hAdapter),
		uintptr(iItem),
		uintptr(iColumn))

	return HIMAGE(ret)
}

/*
获取项文本内容.

参数:
	hAdapter 数据适配器句柄.
	iItem 项索引.
	pName 列名.
	pOut 接收内容缓冲区.
	nOutLen 缓冲区大小,字符为单位.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdTable_GetItemTextEx(hAdapter HXCGUI, iItem int, pName, pOut *uint16, nOutLen int) bool {
	ret, _, _ := xAdTable_GetItemTextEx.Call(
		uintptr(hAdapter),
		uintptr(iItem),
		uintptr(unsafe.Pointer(pName)),
		uintptr(unsafe.Pointer(pOut)),
		uintptr(nOutLen))

	return ret == TRUE
}

/*
获取项图片句柄.

参数:
	hAdapter 数据适配器句柄.
	iItem 项索引.
	pName 列名.*uint16
返回:
	返回图片句柄.
*/
func XAdTable_GetItemImageEx(hAdapter HXCGUI, iItem int, pName string) HIMAGE {
	ret, _, _ := xAdTable_GetItemImageEx.Call(
		uintptr(hAdapter),
		uintptr(iItem),
		StringToUintPtr(pName))
	// uintptr(unsafe.Pointer(pName)))

	return HIMAGE(ret)
}
