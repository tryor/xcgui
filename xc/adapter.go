package xc

import (
	"syscall"
)

var (
	// Functions
	xAd_EnableAutoDestroy *syscall.Proc
)

func init() {
	// Functions
	xAd_EnableAutoDestroy = xcDLL.MustFindProc("XAd_EnableAutoDestroy")

}

func XAd_EnableAutoDestroy(hAdapter HXCGUI, b bool) {
	xAd_EnableAutoDestroy.Call(
		uintptr(hAdapter),
		uintptr(BoolToBOOL(b)))
}
