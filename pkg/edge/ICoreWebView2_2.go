//go:build windows

package edge

import (
	"unsafe"
)

type iCoreWebView2_2Vtbl struct {
	iCoreWebView2Vtbl
	AddWebResourceResponseReceived    ComProc
	RemoveWebResourceResponseReceived ComProc
	NavigateWithWebResourceRequest    ComProc
	AddDomContentLoaded               ComProc
	RemoveDomContentLoaded            ComProc
	GetCookieManager                  ComProc
	GetEnvironment                    ComProc
}

type ICoreWebView2_2 struct {
	vtbl *iCoreWebView2_2Vtbl
}

func (i *ICoreWebView2) GetICoreWebView2_2() *ICoreWebView2_2 {
	var result *ICoreWebView2_2

	iidICoreWebView2_2 := NewGUID("{9E8F0CF8-E670-4B5E-B2BC-73E061E3184C}")
	_, _, _ = i.vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2_2)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (i *ICoreWebView2_2) AddWebResourceResponseReceived(eventHandler *ICoreWebView2WebResourceResponseReceivedEventHandler) (EventRegistrationToken, error) {

	var token EventRegistrationToken

	hr, _, err := i.vtbl.AddWebResourceResponseReceived.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	return token, Error(hr, err)
}
