//go:build windows

package edge

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2ContextMenuItemCollectionVtbl struct {
	IUnknownVtbl
	GetCount           ComProc
	GetValueAtIndex    ComProc
	RemoveValueAtIndex ComProc
	InsertValueAtIndex ComProc
}

type ICoreWebView2ContextMenuItemCollection struct {
	Vtbl *ICoreWebView2ContextMenuItemCollectionVtbl
}

func (i *ICoreWebView2ContextMenuItemCollection) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2ContextMenuItemCollection) GetCount() (uint32, error) {

	var value uint32

	hr, _, err := i.Vtbl.GetCount.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return 0, syscall.Errno(hr)
	}
	return value, err
}

func (i *ICoreWebView2ContextMenuItemCollection) GetValueAtIndex(index uint32) (*ICoreWebView2ContextMenuItem, error) {

	var value *ICoreWebView2ContextMenuItem

	hr, _, err := i.Vtbl.GetValueAtIndex.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(index),
		uintptr(unsafe.Pointer(&value)),
	)
	//if windows.Handle(hr) != windows.S_OK {
	//	return nil, syscall.Errno(hr)
	//}
	return value, Error(hr, err)
}

func (i *ICoreWebView2ContextMenuItemCollection) RemoveValueAtIndex(index uint32) error {
	hr, _, err := i.Vtbl.RemoveValueAtIndex.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(index),
	)
	return Error(hr, err)
}

func (i *ICoreWebView2ContextMenuItemCollection) InsertValueAtIndex(index uint32, value *ICoreWebView2ContextMenuItem) error {
	hr, _, err := i.Vtbl.InsertValueAtIndex.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&index)),
		uintptr(unsafe.Pointer(value)),
	)
	return Error(hr, err)
}
