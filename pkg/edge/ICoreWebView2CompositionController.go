//go:build windows

package edge

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2CompositionControllerVtbl struct {
	_IUnknownVtbl
	GetRootVisualTarget ComProc
	PutRootVisualTarget ComProc
	SendMouseInput      ComProc
	SendPointerInput    ComProc
	GetCursor           ComProc
	GetSystemCursorId   ComProc
	AddCursorChanged    ComProc
	RemoveCursorChanged ComProc
}

type ICoreWebView2CompositionController struct {
	vtbl *ICoreWebView2CompositionControllerVtbl
}

func (i *ICoreWebView2CompositionController) AddRef() uintptr {
	refCounter, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2CompositionController) GetRootVisualTarget() (*IUnknown, error) {

	var target *IUnknown

	hr, _, err := i.vtbl.GetRootVisualTarget.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&target)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	return target, err
}

func (i *ICoreWebView2CompositionController) PutRootVisualTarget(target *IUnknown) error {

	hr, _, err := i.vtbl.PutRootVisualTarget.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(target)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}

func (i *ICoreWebView2CompositionController) SendMouseInput(eventKind COREWEBVIEW2_MOUSE_EVENT_KIND, virtualKeys COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS, mouseData uint32, point POINT) error {

	hr, _, err := i.vtbl.SendMouseInput.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(eventKind),
		uintptr(virtualKeys),
		uintptr(unsafe.Pointer(&mouseData)),
		uintptr(unsafe.Pointer(&point)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}

func (i *ICoreWebView2CompositionController) GetCursor() (HCURSOR, error) {

	var cursor HCURSOR

	hr, _, err := i.vtbl.GetCursor.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&cursor)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return 0, syscall.Errno(hr)
	}
	return cursor, err
}

func (i *ICoreWebView2CompositionController) GetSystemCursorId() (uint32, error) {

	var systemCursorId uint32

	hr, _, err := i.vtbl.GetSystemCursorId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&systemCursorId)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return 0, syscall.Errno(hr)
	}
	return systemCursorId, err
}

func (i *ICoreWebView2CompositionController) RemoveCursorChanged(token EventRegistrationToken) error {

	hr, _, err := i.vtbl.RemoveCursorChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}
