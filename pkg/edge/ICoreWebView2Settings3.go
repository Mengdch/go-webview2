//go:build windows

package edge

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2Settings3Vtbl struct {
	ICoreWebView2Settings2Vtbl
	GetAreBrowserAcceleratorKeysEnabled ComProc
	PutAreBrowserAcceleratorKeysEnabled ComProc
}

type ICoreWebView2Settings3 struct {
	Vtbl *ICoreWebView2Settings3Vtbl
}

func (i *ICoreWebView2Settings3) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2Settings) GetICoreWebView2Settings3() *ICoreWebView2Settings3 {
	var result *ICoreWebView2Settings3

	iidICoreWebView2Settings3 := NewGUID("{fdb5ab74-af33-4854-84f0-0a631deb5eba}")
	_, _, _ = i.vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2Settings3)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (i *ICoreWebView2Settings3) GetAreBrowserAcceleratorKeysEnabled() (bool, error) {
	// Create int32 to hold bool result
	var _areBrowserAcceleratorKeysEnabled int32

	hr, _, err := i.Vtbl.GetAreBrowserAcceleratorKeysEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_areBrowserAcceleratorKeysEnabled)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return false, syscall.Errno(hr)
	}
	// Get result and cleanup
	areBrowserAcceleratorKeysEnabled := _areBrowserAcceleratorKeysEnabled != 0
	return areBrowserAcceleratorKeysEnabled, err
}

func (i *ICoreWebView2Settings3) PutAreBrowserAcceleratorKeysEnabled(areBrowserAcceleratorKeysEnabled bool) error {

	hr, _, err := i.Vtbl.PutAreBrowserAcceleratorKeysEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(boolToInt(areBrowserAcceleratorKeysEnabled)),
	)
	return Error(hr, err)
}
