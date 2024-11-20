//go:build windows

package edge

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2CompositionController2Vtbl struct {
	ICoreWebView2CompositionControllerVtbl
	GetAutomationProvider ComProc
}

type ICoreWebView2CompositionController2 struct {
	vtbl *ICoreWebView2CompositionController2Vtbl
}

func (i *ICoreWebView2CompositionController2) AddRef() uintptr {
	refCounter, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2) GetICoreWebView2CompositionController2() *ICoreWebView2CompositionController2 {
	var result *ICoreWebView2CompositionController2

	iidICoreWebView2CompositionController2 := NewGUID("{0b6a3d24-49cb-4806-ba20-b5e0734a7b26}")
	_, _, _ = i.vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2CompositionController2)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (i *ICoreWebView2CompositionController2) GetAutomationProvider() (*IUnknown, error) {

	var provider *IUnknown

	hr, _, err := i.vtbl.GetAutomationProvider.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&provider)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	return provider, err
}
