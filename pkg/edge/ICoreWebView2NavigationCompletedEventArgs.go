//go:build windows

package edge

import (
	"unsafe"
)

type _ICoreWebView2NavigationCompletedEventArgsVtbl struct {
	_IUnknownVtbl
	GetIsSuccess      ComProc
	GetWebErrorStatus ComProc
	GetNavigationId   ComProc
}

type ICoreWebView2NavigationCompletedEventArgs struct {
	vtbl *_ICoreWebView2NavigationCompletedEventArgsVtbl
}

func (i *ICoreWebView2NavigationCompletedEventArgs) GetNavigationId() (uint64, error) {
	var navigationId uint64
	hr, _, err := i.vtbl.GetNavigationId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&navigationId)),
	)
	err = Error(hr, err)
	if err != nil {
		return 0, err
	}
	return navigationId, nil
}
func (i *ICoreWebView2NavigationCompletedEventArgs) GetIsSuccess() (bool, error) {
	// Create int32 to hold bool result
	var _isSuccess int32

	hr, _, err := i.vtbl.GetIsSuccess.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_isSuccess)),
	)
	err = Error(hr, err)
	if err != nil {
		return false, err
	}
	// Get result and cleanup
	isSuccess := _isSuccess != 0
	return isSuccess, err
}

func (i *ICoreWebView2NavigationCompletedEventArgs) GetWebErrorStatus() (COREWEBVIEW2_WEB_ERROR_STATUS, error) {

	var webErrorStatus COREWEBVIEW2_WEB_ERROR_STATUS

	hr, _, err := i.vtbl.GetWebErrorStatus.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&webErrorStatus)),
	)
	err = Error(hr, err)
	if err != nil {
		return 0, err
	}
	return webErrorStatus, err
}
