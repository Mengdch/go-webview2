//go:build windows

package edge

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2CompositionController3Vtbl struct {
	ICoreWebView2CompositionController2Vtbl
	DragEnter ComProc
	DragLeave ComProc
	DragOver  ComProc
	Drop      ComProc
}

type ICoreWebView2CompositionController3 struct {
	vtbl *ICoreWebView2CompositionController3Vtbl
}

func (i *ICoreWebView2CompositionController3) AddRef() uintptr {
	refCounter, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2) GetICoreWebView2CompositionController3() *ICoreWebView2CompositionController3 {
	var result *ICoreWebView2CompositionController3

	iidICoreWebView2CompositionController3 := NewGUID("{9570570e-4d76-4361-9ee1-f04d0dbdfb1e}")
	_, _, _ = i.vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2CompositionController3)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (i *ICoreWebView2CompositionController3) DragEnter(dataObject *IDataObject, keyState uint32, point POINT) (uint32, error) {

	var effect uint32

	hr, _, err := i.vtbl.DragEnter.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(dataObject)),
		uintptr(unsafe.Pointer(&keyState)),
		uintptr(unsafe.Pointer(&point)),
		uintptr(unsafe.Pointer(&effect)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return 0, syscall.Errno(hr)
	}
	return effect, err
}

func (i *ICoreWebView2CompositionController3) DragLeave() error {

	hr, _, err := i.vtbl.DragLeave.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}

func (i *ICoreWebView2CompositionController3) DragOver(keyState uint32, point POINT) (uint32, error) {

	var effect uint32

	hr, _, err := i.vtbl.DragOver.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&keyState)),
		uintptr(unsafe.Pointer(&point)),
		uintptr(unsafe.Pointer(&effect)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return 0, syscall.Errno(hr)
	}
	return effect, err
}

func (i *ICoreWebView2CompositionController3) Drop(dataObject *IDataObject, keyState uint32, point POINT) (uint32, error) {

	var effect uint32

	hr, _, err := i.vtbl.Drop.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(dataObject)),
		uintptr(unsafe.Pointer(&keyState)),
		uintptr(unsafe.Pointer(&point)),
		uintptr(unsafe.Pointer(&effect)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return 0, syscall.Errno(hr)
	}
	return effect, err
}
