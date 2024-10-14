//go:build windows

package edge

import (
	"github.com/Mengdch/win"
	"unsafe"
)

type ICoreWebView2Settings4Vtbl struct {
	ICoreWebView2Settings3Vtbl
	GetIsPasswordAutosaveEnabled ComProc
	PutIsPasswordAutosaveEnabled ComProc
	GetIsGeneralAutofillEnabled  ComProc
	PutIsGeneralAutofillEnabled  ComProc
}

type ICoreWebView2Settings4 struct {
	Vtbl *ICoreWebView2Settings4Vtbl
}

func (i *ICoreWebView2Settings4) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2Settings) GetICoreWebView2Settings4() *ICoreWebView2Settings4 {
	var result *ICoreWebView2Settings4

	iidICoreWebView2Settings4 := NewGUID("{cb56846c-4168-4d53-b04f-03b6d6796ff2}")
	_, _, _ = i.vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2Settings4)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (i *ICoreWebView2Settings4) GetIsPasswordAutosaveEnabled() (bool, error) {
	// Create int32 to hold bool result
	var _value int32
	hr, _, err := i.Vtbl.GetIsPasswordAutosaveEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	err = Error(hr, err)
	if err != nil {
		return false, err
	}
	// Get result and cleanup
	return _value != 0, nil
}

func (i *ICoreWebView2Settings4) PutIsPasswordAutosaveEnabled(value bool) error {
	hr, _, err := i.Vtbl.PutIsPasswordAutosaveEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(win.BoolToBOOL(value)),
	)
	return Error(hr, err)
}

func (i *ICoreWebView2Settings4) GetIsGeneralAutofillEnabled() (bool, error) {
	// Create int32 to hold bool result
	var _value int32
	hr, _, err := i.Vtbl.GetIsGeneralAutofillEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	err = Error(hr, err)
	if err != nil {
		return false, err
	}
	// Get result and cleanup
	return _value != 0, nil
}

func (i *ICoreWebView2Settings4) PutIsGeneralAutofillEnabled(value bool) error {
	hr, _, err := i.Vtbl.PutIsGeneralAutofillEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(win.BoolToBOOL(value)),
	)
	return Error(hr, err)
}
