package xc

import (
	"syscall"
	"unsafe"
)

var (
	// Functions
	xEdit_Create           *syscall.Proc
	xEdit_CreateEx         *syscall.Proc
	xEdit_SetText          *syscall.Proc
	xEdit_EnableAutoWrap   *syscall.Proc
	xEdit_EnableMultiLine  *syscall.Proc
	xEdit_GetLength        *syscall.Proc
	xEdit_SetCaretColor    *syscall.Proc
	xEdit_GetText          *syscall.Proc
	xEdit_SetTextAlign     *syscall.Proc
	xEdit_SelectAll        *syscall.Proc
	xEdit_GetSelectText    *syscall.Proc
	xEdit_EnableAutoSelAll *syscall.Proc

	xEdit_GetCurRow      *syscall.Proc
	xEdit_GetCurCol      *syscall.Proc
	xEdit_SetCurPos      *syscall.Proc
	xEdit_GetCurPos      *syscall.Proc
	xEdit_SetDefaultText *syscall.Proc
	xEdit_EnablePassword *syscall.Proc
)

func init() {
	xEdit_Create = xcDLL.MustFindProc("XEdit_Create")
	xEdit_CreateEx = xcDLL.MustFindProc("XEdit_CreateEx")
	xEdit_SetText = xcDLL.MustFindProc("XEdit_SetText")
	xEdit_EnableAutoWrap = xcDLL.MustFindProc("XEdit_EnableAutoWrap")
	xEdit_EnableMultiLine = xcDLL.MustFindProc("XEdit_EnableMultiLine")

	//int WINAPI  XEdit_GetLength (HELE hEle)
	xEdit_GetLength = xcDLL.MustFindProc("XEdit_GetLength")
	xEdit_SetCaretColor = xcDLL.MustFindProc("XEdit_SetCaretColor")
	xEdit_GetText = xcDLL.MustFindProc("XEdit_GetText")
	xEdit_SetTextAlign = xcDLL.MustFindProc("XEdit_SetTextAlign")
	xEdit_SelectAll = xcDLL.MustFindProc("XEdit_SelectAll")
	xEdit_GetSelectText = xcDLL.MustFindProc("XEdit_GetSelectText")
	xEdit_EnableAutoSelAll = xcDLL.MustFindProc("XEdit_EnableAutoSelAll")

	xEdit_GetCurRow = xcDLL.MustFindProc("XEdit_GetCurRow")
	xEdit_GetCurCol = xcDLL.MustFindProc("XEdit_GetCurCol")
	xEdit_SetCurPos = xcDLL.MustFindProc("XEdit_SetCurPos")
	xEdit_GetCurPos = xcDLL.MustFindProc("XEdit_GetCurPos")

	xEdit_SetDefaultText = xcDLL.MustFindProc("XEdit_SetDefaultText")
	xEdit_EnablePassword = xcDLL.MustFindProc("XEdit_EnablePassword")
}

func XEdit_GetCurPos(hEle HELE) int {
	r, _, _ := xEdit_GetCurPos.Call(
		uintptr(hEle))
	return int(r)
}

func XEdit_SetCurPos(hEle HELE, r, c int) {
	xEdit_SetCurPos.Call(
		uintptr(hEle), uintptr(r), uintptr(c))
}

func XEdit_GetCurRow(hEle HELE) int {
	r, _, _ := xEdit_GetCurRow.Call(
		uintptr(hEle))
	return int(r)
}

func XEdit_GetCurCol(hEle HELE) int {
	r, _, _ := xEdit_GetCurCol.Call(
		uintptr(hEle))
	return int(r)
}

//XEdit_Create (int x, int y, int cx, int cy, HXCGUI hParent)
func XEdit_Create(x, y, cx, cy int, hParent HXCGUI) HELE {
	ret, _, _ := xEdit_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))

	return HELE(ret)
}

///XEdit_CreateEx (int x, int y, int cx, int cy, edit_type_ type, HXCGUI hParent)
func XEdit_CreateEx(x, y, cx, cy int, typ Edit_type_, hParent HXCGUI) HELE {
	ret, _, _ := xEdit_CreateEx.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(typ),
		uintptr(hParent))

	return HELE(ret)
}

//XEdit_SetText (HELE hEle, const wchar_t *pString)
func XEdit_SetText(hEle HELE, text string) {
	xEdit_SetText.Call(
		uintptr(hEle),
		StringToUintPtr(text))
}

func XEdit_SetDefaultText(hEle HELE, text string) {
	xEdit_SetDefaultText.Call(
		uintptr(hEle),
		StringToUintPtr(text))
}

func XEdit_EnablePassword(hEle HELE, b bool) {
	xEdit_EnablePassword.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(b)))
}

func XEdit_EnableAutoWrap(hEle HELE, b bool) {
	xEdit_EnableAutoWrap.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(b)))
}

//  XEdit_EnableMultiLine (HELE hEle, BOOL bEnable)
func XEdit_EnableMultiLine(hEle HELE, b bool) {
	xEdit_EnableMultiLine.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(b)))
}

func XEdit_EnableAutoSelAll(hEle HELE, b bool) {
	xEdit_EnableAutoSelAll.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(b)))
}

func XEdit_GetLength(hEle HELE) int {
	r, _, _ := xEdit_GetLength.Call(
		uintptr(hEle))
	return int(r)
}

//XEdit_SetCaretColor (HELE hEle, COLORREF color)
func XEdit_SetCaretColor(hEle HELE, c COLORREF) {
	xEdit_SetCaretColor.Call(
		uintptr(hEle),
		uintptr(c))
}

//int WINAPI  XEdit_GetSelectText (HELE hEle, out_ wchar_t *pOut, int nOutLen)
func XEdit_GetSelectText(hEle HELE) string {
	size := XEdit_GetLength(hEle)
	if size <= 0 {
		return ""
	}
	pOut := make([]uint16, size+1)
	r, _, _ := xEdit_GetSelectText.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(&pOut[0])), uintptr(len(pOut)))
	return syscall.UTF16ToString(pOut[:r])
}

func XEdit_SelectLength(hEle HELE) int {
	size := XEdit_GetLength(hEle)
	if size <= 0 {
		return 0
	}
	pOut := make([]uint16, size+1)
	r, _, _ := xEdit_GetSelectText.Call(uintptr(hEle), uintptr(unsafe.Pointer(&pOut[0])), uintptr(len(pOut)))
	return int(r)
}

//int WINAPI  XEdit_GetText (HELE hEle, out_ wchar_t *pOut, int nOutlen)
func XEdit_GetText(hEle HELE) string {
	size := XEdit_GetLength(hEle)
	if size <= 0 {
		return ""
	}
	pOut := make([]uint16, size+1)
	r, _, _ := xEdit_GetText.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(&pOut[0])), uintptr(len(pOut)))
	return syscall.UTF16ToString(pOut[:r])
}

//void WINAPI XEdit_SetTextAlign  ( HELE  hEle,    int  align   )
func XEdit_SetTextAlign(hEle HELE, align Align_flag_) {
	xEdit_SetTextAlign.Call(
		uintptr(hEle),
		uintptr(align))
}

//XEdit_SelectAll (HELE hEle)
func XEdit_SelectAll(hEle HELE) {
	xEdit_SelectAll.Call(
		uintptr(hEle))
}
