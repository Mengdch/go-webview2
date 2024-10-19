//go:build windows

package edge

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type iCoreWebView2_4Vtbl struct {
	iCoreWebView2_3Vtbl
	AddFrameCreated        ComProc
	RemoveFrameCreated     ComProc
	AddDownloadStarting    ComProc
	RemoveDownloadStarting ComProc
}

type ICoreWebView2_4 struct {
	vtbl *iCoreWebView2_4Vtbl
}

func (i *ICoreWebView2) GetICoreWebView2_4() *ICoreWebView2_4 {
	var result *ICoreWebView2_4

	iidICoreWebView2_4 := NewGUID("{20d02d59-6df2-42dc-bd06-f98a694b1302}")
	_, _, _ = i.vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2_4)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (e *Chromium) GetICoreWebView2_4() *ICoreWebView2_4 {
	return e.webview.GetICoreWebView2_4()
}

func (i *ICoreWebView2_4) AddDownloadStarting(eventHandler *ICoreWebView2DownloadStartingEventHandler) (EventRegistrationToken, error) {

	var token EventRegistrationToken

	hr, _, err := i.vtbl.AddDownloadStarting.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return EventRegistrationToken{}, syscall.Errno(hr)
	}
	return token, err
}

func (i *ICoreWebView2_4) RemoveDownloadStarting(token EventRegistrationToken) error {

	hr, _, err := i.vtbl.RemoveDownloadStarting.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}
