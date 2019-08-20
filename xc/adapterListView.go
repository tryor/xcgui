package xc

import (
	"syscall"
	"unsafe"
)

var (
	// Functions
	xAdListView_Create               *syscall.Proc
	xAdListView_Group_AddColumn      *syscall.Proc
	xAdListView_Group_AddItemText    *syscall.Proc
	xAdListView_Group_AddItemTextEx  *syscall.Proc
	xAdListView_Group_AddItemImage   *syscall.Proc
	xAdListView_Group_AddItemImageEx *syscall.Proc
	xAdListView_Group_SetText        *syscall.Proc
	xAdListView_Group_SetTextEx      *syscall.Proc
	xAdListView_Group_SetImage       *syscall.Proc
	xAdListView_Group_SetImageEx     *syscall.Proc
	xAdListView_Item_AddColumn       *syscall.Proc
	xAdListView_Group_GetCount       *syscall.Proc
	xAdListView_Item_GetCount        *syscall.Proc
	xAdListView_Item_AddItemText     *syscall.Proc
	xAdListView_Item_AddItemTextEx   *syscall.Proc
	xAdListView_Item_AddItemImage    *syscall.Proc
	xAdListView_Item_AddItemImageEx  *syscall.Proc
	xAdListView_Group_DeleteItem     *syscall.Proc
	xAdListView_Item_DeleteItem      *syscall.Proc
	xAdListView_DeleteAll            *syscall.Proc
	xAdListView_Item_GetTextEx       *syscall.Proc
	xAdListView_Item_GetImageEx      *syscall.Proc
	xAdListView_Item_SetText         *syscall.Proc
	xAdListView_Item_SetTextEx       *syscall.Proc
	xAdListView_Item_SetImage        *syscall.Proc
	xAdListView_Item_SetImageEx      *syscall.Proc
)

func init() {
	// Functions
	xAdListView_Create = xcDLL.MustFindProc("XAdListView_Create")
	xAdListView_Group_AddColumn = xcDLL.MustFindProc("XAdListView_Group_AddColumn")
	xAdListView_Group_AddItemText = xcDLL.MustFindProc("XAdListView_Group_AddItemText")
	xAdListView_Group_AddItemTextEx = xcDLL.MustFindProc("XAdListView_Group_AddItemTextEx")
	xAdListView_Group_AddItemImage = xcDLL.MustFindProc("XAdListView_Group_AddItemImage")
	xAdListView_Group_AddItemImageEx = xcDLL.MustFindProc("XAdListView_Group_AddItemImageEx")
	xAdListView_Group_SetText = xcDLL.MustFindProc("XAdListView_Group_SetText")
	xAdListView_Group_SetTextEx = xcDLL.MustFindProc("XAdListView_Group_SetTextEx")
	xAdListView_Group_SetImage = xcDLL.MustFindProc("XAdListView_Group_SetImage")
	xAdListView_Group_SetImageEx = xcDLL.MustFindProc("XAdListView_Group_SetImageEx")
	xAdListView_Item_AddColumn = xcDLL.MustFindProc("XAdListView_Item_AddColumn")
	xAdListView_Group_GetCount = xcDLL.MustFindProc("XAdListView_Group_GetCount")
	xAdListView_Item_GetCount = xcDLL.MustFindProc("XAdListView_Item_GetCount")
	xAdListView_Item_AddItemText = xcDLL.MustFindProc("XAdListView_Item_AddItemText")
	xAdListView_Item_AddItemTextEx = xcDLL.MustFindProc("XAdListView_Item_AddItemTextEx")
	xAdListView_Item_AddItemImage = xcDLL.MustFindProc("XAdListView_Item_AddItemImage")
	xAdListView_Item_AddItemImageEx = xcDLL.MustFindProc("XAdListView_Item_AddItemImageEx")
	xAdListView_Group_DeleteItem = xcDLL.MustFindProc("XAdListView_Group_DeleteItem")
	xAdListView_Item_DeleteItem = xcDLL.MustFindProc("XAdListView_Item_DeleteItem")
	xAdListView_DeleteAll = xcDLL.MustFindProc("XAdListView_DeleteAll")
	xAdListView_Item_GetTextEx = xcDLL.MustFindProc("XAdListView_Item_GetTextEx")
	xAdListView_Item_GetImageEx = xcDLL.MustFindProc("XAdListView_Item_GetImageEx")
	xAdListView_Item_SetText = xcDLL.MustFindProc("XAdListView_Item_SetText")
	xAdListView_Item_SetTextEx = xcDLL.MustFindProc("XAdListView_Item_SetTextEx")
	xAdListView_Item_SetImage = xcDLL.MustFindProc("XAdListView_Item_SetImage")
	xAdListView_Item_SetImageEx = xcDLL.MustFindProc("XAdListView_Item_SetImageEx")
}

/*
创建列表视元素数据适配器.

返回:
	返回数据适配器句柄.
*/
func XAdListView_Create() HXCGUI {
	ret, _, _ := xAdListView_Create.Call()

	return HXCGUI(ret)
}

/*
组操作,添加数据列.

参数:
	hAdapter 数据适配器句柄.
	pName 名称.*uint16
返回:
	返回列索引.
*/
func XAdListView_GroupAddColumn(hAdapter HXCGUI, pName string) int {
	ret, _, _ := xAdListView_Group_AddColumn.Call(
		uintptr(hAdapter),
		StringToUintPtr(pName))
	// uintptr(unsafe.Pointer(pName)))

	return int(ret)
}

/*
组操作,添加组,数据默认填充到第一列.

参数:
	hAdapter 数据适配器句柄.
	pValue 值.*uint16
返回:
	返回组索引.
*/
func XAdListView_GroupAddItemText(hAdapter HXCGUI, pValue string) int {
	ret, _, _ := xAdListView_Group_AddItemText.Call(
		uintptr(hAdapter),
		StringToUintPtr(pValue))
	// uintptr(unsafe.Pointer(pValue)))

	return int(ret)
}

/*
组操作,添加组,数据填充指定列.

参数:
	hAdapter 数据适配器句柄.
	pName 列名.
	pValue 值.*uint16
返回:
	返回组索引.
*/
func XAdListView_GroupAddItemTextEx(hAdapter HXCGUI, pName, pValue string) int {
	ret, _, _ := xAdListView_Group_AddItemTextEx.Call(
		uintptr(hAdapter),
		StringToUintPtr(pName),
		StringToUintPtr(pValue))
	// uintptr(unsafe.Pointer(pName)),
	// uintptr(unsafe.Pointer(pValue)))

	return int(ret)
}

/*
组操作,添加组,数据默认填充第一列.

参数:
	hAdapter 数据适配器句柄.
	hImage 图片句柄.
返回:
	返回组索引.
*/
func XAdListView_GroupAddItemImage(hAdapter HXCGUI, hImage HIMAGE) int {
	ret, _, _ := xAdListView_Group_AddItemImage.Call(
		uintptr(hAdapter),
		uintptr(hImage))

	return int(ret)
}

/*
组操作,添加组,数据填充指定列.

参数:
	hAdapter 数据适配器句柄.
	pName 列名. *uint16
	hImage 图片句柄.
返回:
	返回组索引.
*/
func XAdListView_GroupAddItemImageEx(hAdapter HXCGUI, pName string, hImage HIMAGE) int {
	ret, _, _ := xAdListView_Group_AddItemImageEx.Call(
		uintptr(hAdapter),
		StringToUintPtr(pName),
		// uintptr(unsafe.Pointer(pName)),
		uintptr(hImage))

	return int(ret)
}

/*
组操作,设置指定项数据.

参数:
	hAdapter 数据适配器句柄.
	iGroup 组索引.
	iColumn 列索引.
	pValue 值.*uint16
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdListView_GroupSetText(hAdapter HXCGUI, iGroup, iColumn int, pValue string) bool {
	ret, _, _ := xAdListView_Group_SetText.Call(
		uintptr(hAdapter),
		uintptr(iGroup),
		uintptr(iColumn),
		StringToUintPtr(pValue))
	// uintptr(unsafe.Pointer(pValue)))

	return ret == TRUE
}

/*
组操作,设置指定项数据.

参数:
	hAdapter 数据适配器句柄.
	iGroup 组索引.
	pName 字段名.
	pValue 值.*uint16
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdListView_GroupSetTextEx(hAdapter HXCGUI, iGroup int, pName, pValue string) bool {
	ret, _, _ := xAdListView_Group_SetTextEx.Call(
		uintptr(hAdapter),
		uintptr(iGroup),
		StringToUintPtr(pName),
		StringToUintPtr(pValue))
	// uintptr(unsafe.Pointer(pName)),
	// uintptr(unsafe.Pointer(pValue)))

	return ret == TRUE
}

/*
组操作,设置指定项数据.

参数:
	hAdapter 数据适配器句柄.
	iGroup 组索引.
	iColumn 列索引.
	hImage 图片句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdListView_GroupSetImage(hAdapter HXCGUI, iGroup, iColumn int, hImage HIMAGE) bool {
	ret, _, _ := xAdListView_Group_SetImage.Call(
		uintptr(hAdapter),
		uintptr(iGroup),
		uintptr(iColumn),
		uintptr(hImage))

	return ret == TRUE
}

/*
组操作,设置指定项数据.

参数:
	hAdapter 数据适配器句柄.
	iGroup 组索引.
	pName 字段名.*uint16
	hImage 图片句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdListView_GroupSetImageEx(hAdapter HXCGUI, iGroup int, pName string, hImage HIMAGE) bool {
	ret, _, _ := xAdListView_Group_SetImageEx.Call(
		uintptr(hAdapter),
		uintptr(iGroup),
		StringToUintPtr(pName),
		// uintptr(unsafe.Pointer(pName)),
		uintptr(hImage))

	return ret == TRUE
}

/*
项操作,添加列.

参数:
	hAdapter 数据适配器句柄.
	pName 名称.*uint16
返回:
	返回列索引.
*/
func XAdListView_ItemAddColumn(hAdapter HXCGUI, pName string) int {
	ret, _, _ := xAdListView_Item_AddColumn.Call(
		uintptr(hAdapter),
		StringToUintPtr(pName))
	// uintptr(unsafe.Pointer(pName)))

	return int(ret)
}

/*
组操作,获取组数量.

参数:
	hAdapter 数据适配器句柄.
返回:
	返回组数量.
*/
func XAdListView_GroupGetCount(hAdapter HXCGUI) int {
	ret, _, _ := xAdListView_Group_GetCount.Call(uintptr(hAdapter))

	return int(ret)
}

/*
项操作,获取指定组中项数量.

参数:
	hAdapter 数据适配器句柄.
	iGroup 组索引.
返回:
	成功返回项数量,否则返回 XC_ID_ERROR.
*/
func XAdListView_ItemGetCount(hAdapter HXCGUI, iGroup int) int {
	ret, _, _ := xAdListView_Item_GetCount.Call(
		uintptr(hAdapter),
		uintptr(iGroup))

	return int(ret)
}

/*
项操作,添加项.

参数:
	hAdapter 数据适配器句柄.
	iGroup 组索引.
	pValue 值.*uint16
返回:
	返回项索引.
*/
func XAdListView_ItemAddItemText(hAdapter HXCGUI, iGroup int, pValue string) int {
	ret, _, _ := xAdListView_Item_AddItemText.Call(
		uintptr(hAdapter),
		uintptr(iGroup),
		StringToUintPtr(pValue))
	// uintptr(unsafe.Pointer(pValue)))

	return int(ret)
}

/*
项操作,添加项,数据填充指定列.

参数:
	hAdapter 数据适配器句柄.
	iGroup 组索引.
	pName 列名称.
	pValue 值.*uint16
返回:
	返回项索引.
*/
func XAdListView_ItemAddItemTextEx(hAdapter HXCGUI, iGroup int, pName, pValue string) int {
	ret, _, _ := xAdListView_Item_AddItemTextEx.Call(
		uintptr(hAdapter),
		uintptr(iGroup),
		StringToUintPtr(pName),
		StringToUintPtr(pValue))
	// uintptr(unsafe.Pointer(pName)),
	// uintptr(unsafe.Pointer(pValue)))

	return int(ret)
}

/*
项操作,添加项.

参数:
	hAdapter 数据适配器句柄.
	iGroup 组索引.
	hImage 图片句柄.
返回:
	返回项索引.
*/
func XAdListView_ItemAddItemImage(hAdapter HXCGUI, iGroup int, hImage HIMAGE) int {
	ret, _, _ := xAdListView_Item_AddItemImage.Call(
		uintptr(hAdapter),
		uintptr(iGroup),
		uintptr(hImage))

	return int(ret)
}

/*
项操作,添加项,填充指定列数据.

参数:
	hAdapter 数据适配器句柄.
	iGroup 组索引.
	pName 列名称. *uint16
	hImage 图片句柄.
返回:
	返回项索引.
*/
func XAdListView_ItemAddItemImageEx(hAdapter HXCGUI, iGroup int, pName string, hImage HIMAGE) int {
	ret, _, _ := xAdListView_Item_AddItemImageEx.Call(
		uintptr(hAdapter),
		uintptr(iGroup),
		StringToUintPtr(pName),
		// uintptr(unsafe.Pointer(pName)),
		uintptr(hImage))

	return int(ret)
}

/*
删除组,自动删除子项.

参数:
	hAdapter 数据适配器句柄.
	iGroup 组索引.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdListView_GroupDeleteItem(hAdapter HXCGUI, iGroup int) bool {
	ret, _, _ := xAdListView_Group_DeleteItem.Call(
		uintptr(hAdapter),
		uintptr(iGroup))

	return ret == TRUE
}

/*
删除项.

参数:
	hAdapter 数据适配器句柄.
	iGroup 组索引.
	iItem 项索引.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdListView_ItemDeleteItem(hAdapter HXCGUI, iGroup, iItem int) bool {
	ret, _, _ := xAdListView_Item_DeleteItem.Call(
		uintptr(hAdapter),
		uintptr(iGroup),
		uintptr(iItem))

	return ret == TRUE
}

/*
删除所有.

参数:
	hAdapter 数据适配器句柄.
*/
func XAdListView_DeleteAll(hAdapter HXCGUI) {
	xAdListView_DeleteAll.Call(uintptr(hAdapter))
}

/*
项操作,获取项文本内容.

参数:
	hAdapter 数据适配器句柄.
	iGroup 组索引.
	iItem 项索引.
	pName 列名.
	pOut 接收内容缓冲区.
	nOutLen 缓冲区大小,字符为单位.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdListView_ItemGetTextEx(hAdapter HXCGUI, iGroup, iItem int, pName, pOut *uint16, nOutLen int) bool {
	ret, _, _ := xAdListView_Item_GetTextEx.Call(
		uintptr(hAdapter),
		uintptr(iGroup),
		uintptr(iItem),
		uintptr(unsafe.Pointer(pName)),
		uintptr(unsafe.Pointer(pOut)),
		uintptr(nOutLen))

	return ret == TRUE
}

/*
项操作,获取项图片句柄.

参数:
	hAdapter 数据适配器句柄.
	iGroup 组索引.
	iItem 项索引.
	pName 列名.
返回:
	返回图片句柄.
*/
func XAdListView_ItemGetImageEx(hAdapter HXCGUI, iGroup, iItem int, pName *uint16) HIMAGE {
	ret, _, _ := xAdListView_Item_GetImageEx.Call(
		uintptr(hAdapter),
		uintptr(iGroup),
		uintptr(iItem),
		uintptr(unsafe.Pointer(pName)))

	return HIMAGE(ret)
}

/*
项操作,数据填充指定列.

参数:
	hAdapter 数据适配器句柄.
	iGroup 组索引.
	iItem 项索引.
	iColumn 列索引.
	pValue 值.*uint16
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdListView_ItemSetText(hAdapter HXCGUI, iGroup, iItem, iColumn int, pValue string) bool {
	ret, _, _ := xAdListView_Item_SetText.Call(
		uintptr(hAdapter),
		uintptr(iGroup),
		uintptr(iItem),
		uintptr(iColumn),
		StringToUintPtr(pValue))
	// uintptr(unsafe.Pointer(pValue)))

	return ret == TRUE
}

/*
项操作,数据填充指定列.

参数:
	hAdapter 数据适配器句柄.
	iGroup 组索引.
	iItem 项索引.
	pName 列名称.
	pValue 值.*uint16
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdListView_ItemSetTextEx(hAdapter HXCGUI, iGroup, iItem int, pName, pValue string) bool {
	ret, _, _ := xAdListView_Item_SetTextEx.Call(
		uintptr(hAdapter),
		uintptr(iGroup),
		uintptr(iItem),
		StringToUintPtr(pName),
		StringToUintPtr(pValue))
	// uintptr(unsafe.Pointer(pName)),
	// uintptr(unsafe.Pointer(pValue)))

	return ret == TRUE
}

/*
项操作,数据填充指定列.

参数:
	hAdapter 数据适配器句柄.
	iGroup 组索引.
	iItem 项索引.
	iColumn 列索引.
	hImage 图片句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdListView_ItemSetImage(hAdapter HXCGUI, iGroup, iItem, iColumn int, hImage HIMAGE) bool {
	ret, _, _ := xAdListView_Item_SetImage.Call(
		uintptr(hAdapter),
		uintptr(iGroup),
		uintptr(iItem),
		uintptr(iColumn),
		uintptr(hImage))

	return ret == TRUE
}

/*
项操作,数据填充指定列.

参数:
	hAdapter 数据适配器句柄.
	iGroup 组索引.
	iItem 项索引.
	pName 列名称.*uint16
	hImage 图片句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdListView_ItemSetImageEx(hAdapter HXCGUI, iGroup, iItem int, pName string, hImage HIMAGE) bool {
	ret, _, _ := xAdListView_Item_SetImageEx.Call(
		uintptr(hAdapter),
		uintptr(iGroup),
		uintptr(iItem),
		StringToUintPtr(pName),
		// uintptr(unsafe.Pointer(pName)),
		uintptr(hImage))

	return ret == TRUE
}
