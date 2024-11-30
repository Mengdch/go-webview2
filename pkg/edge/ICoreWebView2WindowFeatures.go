//go:build windows

package edge

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2WindowFeaturesVtbl struct {
	_IUnknownVtbl
	GetHasPosition             ComProc
	GetHasSize                 ComProc
	GetLeft                    ComProc
	GetTop                     ComProc
	GetHeight                  ComProc
	GetWidth                   ComProc
	GetShouldDisplayMenuBar    ComProc
	GetShouldDisplayStatus     ComProc
	GetShouldDisplayToolbar    ComProc
	GetShouldDisplayScrollBars ComProc
}

type ICoreWebView2WindowFeatures struct {
	vtbl *ICoreWebView2WindowFeaturesVtbl
}

func (i *ICoreWebView2WindowFeatures) GetHasPosition() (bool, error) {
	// Create int32 to hold bool result
	var _value int32

	hr, _, err := i.vtbl.GetHasPosition.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	// Get result and cleanup
	value := _value != 0
	return value, Error(hr, err)
}

func (i *ICoreWebView2WindowFeatures) GetHasSize() (bool, error) {
	// Create int32 to hold bool result
	var _value int32

	hr, _, err := i.vtbl.GetHasSize.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return false, syscall.Errno(hr)
	}
	// Get result and cleanup
	value := _value != 0
	return value, err
}

func (i *ICoreWebView2WindowFeatures) GetLeft() (uint32, error) {

	var value uint32

	hr, _, err := i.vtbl.GetLeft.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return 0, syscall.Errno(hr)
	}
	return value, err
}

func (i *ICoreWebView2WindowFeatures) GetTop() (uint32, error) {

	var value uint32

	hr, _, err := i.vtbl.GetTop.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return 0, syscall.Errno(hr)
	}
	return value, err
}

func (i *ICoreWebView2WindowFeatures) GetHeight() (uint32, error) {

	var value uint32

	hr, _, err := i.vtbl.GetHeight.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return 0, syscall.Errno(hr)
	}
	return value, err
}

func (i *ICoreWebView2WindowFeatures) GetWidth() (uint32, error) {

	var value uint32

	hr, _, err := i.vtbl.GetWidth.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return 0, syscall.Errno(hr)
	}
	return value, err
}

func (i *ICoreWebView2WindowFeatures) GetShouldDisplayMenuBar() (bool, error) {
	// Create int32 to hold bool result
	var _value int32

	hr, _, err := i.vtbl.GetShouldDisplayMenuBar.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return false, syscall.Errno(hr)
	}
	// Get result and cleanup
	value := _value != 0
	return value, err
}

func (i *ICoreWebView2WindowFeatures) GetShouldDisplayStatus() (bool, error) {
	// Create int32 to hold bool result
	var _value int32

	hr, _, err := i.vtbl.GetShouldDisplayStatus.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return false, syscall.Errno(hr)
	}
	// Get result and cleanup
	value := _value != 0
	return value, err
}

func (i *ICoreWebView2WindowFeatures) GetShouldDisplayToolbar() (bool, error) {
	// Create int32 to hold bool result
	var _value int32

	hr, _, err := i.vtbl.GetShouldDisplayToolbar.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return false, syscall.Errno(hr)
	}
	// Get result and cleanup
	value := _value != 0
	return value, err
}

func (i *ICoreWebView2WindowFeatures) GetShouldDisplayScrollBars() (bool, error) {
	// Create int32 to hold bool result
	var _value int32

	hr, _, err := i.vtbl.GetShouldDisplayScrollBars.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return false, syscall.Errno(hr)
	}
	// Get result and cleanup
	value := _value != 0
	return value, err
}
