//go:build windows

package edge

import (
	"unsafe"
)

type iCoreWebView2_7Vtbl struct {
	iCoreWebView2_6Vtbl
	PrintToPdf ComProc
}

type ICoreWebView2_7 struct {
	vtbl *iCoreWebView2_7Vtbl
}

func (i *ICoreWebView2) GetICoreWebView2_7() *ICoreWebView2_7 {
	var result *ICoreWebView2_7

	iidICoreWebView2_7 := NewGUID("{79c24d83-09a3-45ae-9418-487f32a58740}")
	_, _, _ = i.vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2_7)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (e *Chromium) GetICoreWebView2_7() *ICoreWebView2_7 {
	return e.webview.GetICoreWebView2_7()
}
