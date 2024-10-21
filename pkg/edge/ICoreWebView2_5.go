//go:build windows

package edge

import (
	"unsafe"
)

type iCoreWebView2_5Vtbl struct {
	iCoreWebView2_4Vtbl
	AddClientCertificateRequested    ComProc
	RemoveClientCertificateRequested ComProc
}

type ICoreWebView2_5 struct {
	vtbl *iCoreWebView2_5Vtbl
}

func (i *ICoreWebView2) GetICoreWebView2_5() *ICoreWebView2_5 {
	var result *ICoreWebView2_5

	iidICoreWebView2_5 := NewGUID("{bedb11b8-d63c-11eb-b8bc-0242ac130003}")
	_, _, _ = i.vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2_5)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (e *Chromium) GetICoreWebView2_5() *ICoreWebView2_5 {
	return e.webview.GetICoreWebView2_5()
}
