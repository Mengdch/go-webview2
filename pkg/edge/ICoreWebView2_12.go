//go:build windows

package edge

import (
	"unsafe"
)

type iCoreWebView2_12Vtbl struct {
	iCoreWebView2_11Vtbl
	AddStatusBarTextChanged    ComProc
	RemoveStatusBarTextChanged ComProc
	GetStatusBarText           ComProc
}

type ICoreWebView2_12 struct {
	vtbl *iCoreWebView2_12Vtbl
}

func (i *ICoreWebView2) GetICoreWebView2_12() *ICoreWebView2_12 {
	var result *ICoreWebView2_12

	iidICoreWebView2_12 := NewGUID("{35D69927-BCFA-4566-9349-6B3E0D154CAC}")
	_, _, _ = i.vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2_12)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (e *Chromium) GetICoreWebView2_12() *ICoreWebView2_12 {
	return e.webview.GetICoreWebView2_12()
}
