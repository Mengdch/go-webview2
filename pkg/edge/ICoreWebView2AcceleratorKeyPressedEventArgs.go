//go:build windows

package edge

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

type _ICoreWebView2AcceleratorKeyPressedEventArgsVtbl struct {
	_IUnknownVtbl
	GetKeyEventKind      ComProc
	GetVirtualKey        ComProc
	GetKeyEventLParam    ComProc
	GetPhysicalKeyStatus ComProc
	GetHandled           ComProc
	PutHandled           ComProc
}

type ICoreWebView2AcceleratorKeyPressedEventArgs struct {
	vtbl *_ICoreWebView2AcceleratorKeyPressedEventArgsVtbl
}

func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) GetKeyEventKind() (COREWEBVIEW2_KEY_EVENT_KIND, error) {
	var err error
	var keyEventKind COREWEBVIEW2_KEY_EVENT_KIND
	_, _, err = i.vtbl.GetKeyEventKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&keyEventKind)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0, err
	}
	return keyEventKind, nil
}

func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) GetVirtualKey() (uint, error) {
	var err error
	var virtualKey uint
	_, _, err = i.vtbl.GetVirtualKey.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&virtualKey)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0, err
	}
	return virtualKey, nil
}

func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) GetKeyEventLParam() (int, error) {

	var lParam int

	hr, _, err := i.vtbl.GetKeyEventLParam.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&lParam)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return 0, syscall.Errno(hr)
	}
	return lParam, err
}

func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) GetPhysicalKeyStatus() (COREWEBVIEW2_PHYSICAL_KEY_STATUS, error) {
	var err error
	var physicalKeyStatus internal_COREWEBVIEW2_PHYSICAL_KEY_STATUS
	_, _, err = i.vtbl.GetPhysicalKeyStatus.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&physicalKeyStatus)),
	)
	if err != windows.ERROR_SUCCESS {
		return COREWEBVIEW2_PHYSICAL_KEY_STATUS{}, err
	}
	return COREWEBVIEW2_PHYSICAL_KEY_STATUS{
		RepeatCount:   physicalKeyStatus.RepeatCount,
		ScanCode:      physicalKeyStatus.ScanCode,
		IsExtendedKey: physicalKeyStatus.IsExtendedKey != 0,
		IsMenuKeyDown: physicalKeyStatus.IsMenuKeyDown != 0,
		WasKeyDown:    physicalKeyStatus.WasKeyDown != 0,
		IsKeyReleased: physicalKeyStatus.IsKeyReleased != 0,
	}, nil
}

func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) PutHandled(handled bool) error {
	var err error

	_, _, err = i.vtbl.PutHandled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(boolToInt(handled)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
