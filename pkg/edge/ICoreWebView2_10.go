//go:build windows

package edge

import (
	"unsafe"
)

type iCoreWebView2_10Vtbl struct {
	iCoreWebView2_9Vtbl
	AddBasicAuthenticationRequested    ComProc
	RemoveBasicAuthenticationRequested ComProc
}

type ICoreWebView2_10 struct {
	vtbl *iCoreWebView2_10Vtbl
}

func (i *ICoreWebView2) GetICoreWebView2_10() *ICoreWebView2_10 {
	var result *ICoreWebView2_10

	iidICoreWebView2_10 := NewGUID("{b1690564-6f5a-4983-8e48-31d1143fecdb}")
	_, _, _ = i.vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2_10)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (e *Chromium) GetICoreWebView2_10() *ICoreWebView2_10 {
	return e.webview.GetICoreWebView2_10()
}
