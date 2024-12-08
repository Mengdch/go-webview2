//go:build windows

package edge

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type iCoreWebView2_15vtbl struct {
	iCoreWebView2_14Vtbl
	AddFaviconChanged    ComProc
	RemoveFaviconChanged ComProc
	GetFaviconUri        ComProc
	GetFavicon           ComProc
}

type ICoreWebView2_15 struct {
	vtbl *iCoreWebView2_15vtbl
}

func (i *ICoreWebView2) GetICoreWebView2_15() *ICoreWebView2_15 {
	var result *ICoreWebView2_15

	iidICoreWebView2_15 := NewGUID("{517B2D1D-7DAE-4A66-A4F4-10352FFB9518}")
	_, _, _ = i.vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2_15)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (e *Chromium) GetICoreWebView2_15() *ICoreWebView2_15 {
	return e.webview.GetICoreWebView2_15()
}
func (i *ICoreWebView2_15) AddFaviconChanged(eventHandler *ICoreWebView2FaviconChangedEventHandler) (EventRegistrationToken, error) {

	var token EventRegistrationToken

	hr, _, err := i.vtbl.AddFaviconChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return EventRegistrationToken{}, syscall.Errno(hr)
	}
	return token, err
}

func (i *ICoreWebView2_15) RemoveFaviconChanged(token EventRegistrationToken) error {

	hr, _, err := i.vtbl.RemoveFaviconChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}

func (i *ICoreWebView2_15) GetFaviconUri() (string, error) {
	// Create *uint16 to hold result
	var _value *uint16

	hr, _, err := i.vtbl.GetFaviconUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return "", syscall.Errno(hr)
	}
	// Get result and cleanup
	value := UTF16PtrToString(_value)
	CoTaskMemFree(unsafe.Pointer(_value))
	return value, err
}

type COREWEBVIEW2_FAVICON_IMAGE_FORMAT uint32

const (
	COREWEBVIEW2_FAVICON_IMAGE_FORMAT_PNG  = 0
	COREWEBVIEW2_FAVICON_IMAGE_FORMAT_JPEG = 1
)

func (i *ICoreWebView2_15) GetFavicon(format COREWEBVIEW2_FAVICON_IMAGE_FORMAT, completedHandler *ICoreWebView2GetFaviconCompletedHandler) error {

	hr, _, err := i.vtbl.GetFavicon.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(format),
		uintptr(unsafe.Pointer(completedHandler)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}
