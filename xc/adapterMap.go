package xc

import (
	"syscall"
	"unsafe"
)

var (
	// Functions
	xAdMap_Create       *syscall.Proc
	xAdMap_AddItemText  *syscall.Proc
	xAdMap_AddItemImage *syscall.Proc
	xAdMap_DeleteItem   *syscall.Proc
	xAdMap_GetCount     *syscall.Proc
	xAdMap_GetItemText  *syscall.Proc
	xAdMap_GetItemImage *syscall.Proc
)

func init() {
	// Funtions
	xAdMap_Create = xcDLL.MustFindProc("XAdMap_Create")
	xAdMap_AddItemText = xcDLL.MustFindProc("XAdMap_AddItemText")
	xAdMap_AddItemImage = xcDLL.MustFindProc("XAdMap_AddItemImage")
	xAdMap_DeleteItem = xcDLL.MustFindProc("XAdMap_DeleteItem")
	xAdMap_GetCount = xcDLL.MustFindProc("XAdMap_GetCount")
	xAdMap_GetItemText = xcDLL.MustFindProc("XAdMap_GetItemText")
	xAdMap_GetItemImage = xcDLL.MustFindProc("XAdMap_GetItemImage")
}

/*
创建数据适配器,单列数据.

返回:
	返回数据适配器句柄.
*/
func XAdMap_Create() HXCGUI {
	ret, _, _ := xAdMap_Create.Call()

	return HXCGUI(ret)
}

/*
增加数据项.

参数:
	hAdapter 数据适配器句柄.
	pName 项名称.
	pValue 值.*uint16
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdMap_AddItemText(hAdapter HXCGUI, pName, pValue string) bool {
	ret, _, _ := xAdMap_AddItemText.Call(
		uintptr(hAdapter),
		StringToUintPtr(pName),
		StringToUintPtr(pValue))
	// uintptr(unsafe.Pointer(pName)),
	// uintptr(unsafe.Pointer(pValue)))

	return ret == TRUE
}

/*
增加数据项.

参数:
	hAdapter 数据适配器句柄.
	pName 项名称.*uint16
	hImage 图片句柄.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdMap_AddItemImage(hAdapter HXCGUI, pName string, hImage HIMAGE) bool {
	ret, _, _ := xAdMap_AddItemImage.Call(
		uintptr(hAdapter),
		StringToUintPtr(pName),
		// uintptr(unsafe.Pointer(pName)),
		uintptr(hImage))

	return ret == TRUE
}

/*
删除数据项.

参数:
	hAdapter 数据适配器句柄.
	pName 项名称.*uint16
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdMap_DeleteItem(hAdapter HXCGUI, pName string) bool {
	ret, _, _ := xAdMap_DeleteItem.Call(
		uintptr(hAdapter),
		StringToUintPtr(pName))
	// uintptr(unsafe.Pointer(pName)))

	return ret == TRUE
}

/*
获取项数量.

参数:
	hAdapter 数据适配器句柄.
返回:
	返回项数量.
*/
func XAdMap_GetCount(hAdapter HXCGUI) int {
	ret, _, _ := xAdMap_GetCount.Call(uintptr(hAdapter))

	return int(ret)
}

/*
获取项内容,如果内容为文本.

参数:
	hAdapter 数据适配器句柄.
	pName 项名称.
	pOut 接收返回文本缓冲区.
	nOutLen 缓冲区长度,字符为单位.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XAdMap_GetItemText(hAdapter HXCGUI, pName, pOut *uint16, nOutLen int) bool {
	ret, _, _ := xAdMap_GetItemText.Call(
		uintptr(hAdapter),
		uintptr(unsafe.Pointer(pName)),
		uintptr(unsafe.Pointer(pOut)),
		uintptr(nOutLen))

	return ret == TRUE
}

/*
获取项内容,如果内容为图片句柄.

参数:
	hAdapter 数据适配器句柄.
	pName 项名称.*uint16
返回:
	返回图片句柄.
*/
func XAdMap_GetItemImage(hAdapter HXCGUI, pName string) HIMAGE {
	ret, _, _ := xAdMap_GetItemImage.Call(
		uintptr(hAdapter),
		StringToUintPtr(pName))
	// uintptr(unsafe.Pointer(pName)))

	return HIMAGE(ret)
}
