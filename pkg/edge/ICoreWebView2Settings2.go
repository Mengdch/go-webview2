//go:build windows

package edge

import (
	"unsafe"
)

type ICoreWebView2Settings2Vtbl struct {
	_ICoreWebView2SettingsVtbl
	GetUserAgent ComProc
	PutUserAgent ComProc
}

type ICoreWebView2Settings2 struct {
	Vtbl *ICoreWebView2Settings2Vtbl
}

func (i *ICoreWebView2Settings2) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2Settings) GetICoreWebView2Settings2() *ICoreWebView2Settings2 {
	var result *ICoreWebView2Settings2

	iidICoreWebView2Settings2 := NewGUID("{ee9a0f68-f46c-4e32-ac23-ef8cac224d2a}")
	_, _, _ = i.vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2Settings2)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (i *ICoreWebView2Settings2) GetUserAgent() (string, error) {
	// Create *uint16 to hold result
	var _userAgent *uint16

	hr, _, err := i.Vtbl.GetUserAgent.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_userAgent)),
	)
	err = Error(hr, err)
	if err != nil {
		return "", err
	}
	// Get result and cleanup
	userAgent := UTF16PtrToString(_userAgent)
	CoTaskMemFree(unsafe.Pointer(_userAgent))
	return userAgent, nil
}

func (i *ICoreWebView2Settings2) PutUserAgent(userAgent string) error {

	// Convert string 'userAgent' to *uint16
	_userAgent, err := UTF16PtrFromString(userAgent)
	if err != nil {
		return err
	}

	hr, _, err := i.Vtbl.PutUserAgent.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_userAgent)),
	)
	return Error(hr, err)
}
