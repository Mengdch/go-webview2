//go:build windows

package edge

import (
	"unsafe"
)

type iCoreWebView2_6Vtbl struct {
	iCoreWebView2_5Vtbl
	OpenTaskManagerWindow ComProc
}

type ICoreWebView2_6 struct {
	vtbl *iCoreWebView2_6Vtbl
}

func (i *ICoreWebView2) GetICoreWebView2_6() *ICoreWebView2_6 {
	var result *ICoreWebView2_6

	iidICoreWebView2_6 := NewGUID("{499aadac-d92c-4589-8a75-111bfc167795}")
	_, _, _ = i.vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2_6)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (e *Chromium) GetICoreWebView2_6() *ICoreWebView2_6 {
	return e.webview.GetICoreWebView2_6()
}
