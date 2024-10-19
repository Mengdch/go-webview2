//go:build windows

package edge

import (
	"errors"
	"unsafe"

	"golang.org/x/sys/windows"
)

type iCoreWebView2_3Vtbl struct {
	iCoreWebView2_2Vtbl
	TrySuspend                          ComProc
	Resume                              ComProc
	GetIsSuspended                      ComProc
	SetVirtualHostNameToFolderMapping   ComProc
	ClearVirtualHostNameToFolderMapping ComProc
}

type ICoreWebView2_3 struct {
	vtbl *iCoreWebView2_3Vtbl
}

func (i *ICoreWebView2_3) SetVirtualHostNameToFolderMapping(hostName, folderPath string, accessKind COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND) error {
	_hostName, err := windows.UTF16PtrFromString(hostName)
	if err != nil {
		return err
	}

	_folderPath, err := windows.UTF16PtrFromString(folderPath)
	if err != nil {
		return err
	}

	_, _, err = i.vtbl.SetVirtualHostNameToFolderMapping.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_hostName)),
		uintptr(unsafe.Pointer(_folderPath)),
		uintptr(accessKind),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}

	return nil
}

func (i *ICoreWebView2) GetICoreWebView2_3() *ICoreWebView2_3 {
	var result *ICoreWebView2_3

	iidICoreWebView2_3 := NewGUID("{A0D6DF20-3B92-416D-AA0C-437A9C727857}")
	_, _, _ = i.vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2_3)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (e *Chromium) GetICoreWebView2_3() *ICoreWebView2_3 {
	return e.webview.GetICoreWebView2_3()
}
func (e *Chromium) TrySuspend() error {
	if e.webview == nil {
		return errors.New("webviewNil")
	}
	v3 := e.GetICoreWebView2_3()
	if v3 == nil {
		return errors.New("GetICoreWebView2_3Nil")
	}
	return v3.TrySuspend(e.trySuspendCompleted)
}
func (i *ICoreWebView2_3) TrySuspend(handler *ICoreWebView2TrySuspendCompletedHandler) error {

	hr, _, err := i.vtbl.TrySuspend.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(handler)),
	)
	return Error(hr, err)
}

func (i *ICoreWebView2_3) Resume() error {

	hr, _, err := i.vtbl.Resume.Call(
		uintptr(unsafe.Pointer(i)),
	)
	return Error(hr, err)
}

func (i *ICoreWebView2_3) GetIsSuspended() (bool, error) {
	// Create int32 to hold bool result
	var _isSuspended int32

	hr, _, err := i.vtbl.GetIsSuspended.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_isSuspended)),
	)
	err = Error(hr, err)
	if err != nil {
		return false, err
	}
	// Get result and cleanup
	isSuspended := _isSuspended != 0
	return isSuspended, nil
}
