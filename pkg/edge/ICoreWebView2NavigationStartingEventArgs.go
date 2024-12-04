//go:build windows

package edge

import (
	"unsafe"
)

type _ICoreWebView2NavigationStartingEventArgsVtbl struct {
	_IUnknownVtbl
	GetUri                             ComProc
	GetIsUserInitiated                 ComProc
	GetIsRedirected                    ComProc
	GetRequestHeaders                  ComProc
	GetCancel                          ComProc
	PutCancel                          ComProc
	GetNavigationId                    ComProc
	GetAdditionalAllowedFrameAncestors ComProc
	PutAdditionalAllowedFrameAncestors ComProc
	GetNavigationKind                  ComProc
}

type ICoreWebView2NavigationStartingEventArgs struct {
	vtbl *_ICoreWebView2NavigationStartingEventArgsVtbl
}

func (i *ICoreWebView2NavigationStartingEventArgs) GetUri() (string, error) {
	// Create *uint16 to hold result
	var _uri *uint16
	hr, _, err := i.vtbl.GetUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_uri)),
	)
	err = Error(hr, err)
	if err != nil {
		return "", err
	}
	// Get result and cleanup
	uri := UTF16PtrToString(_uri)
	CoTaskMemFree(unsafe.Pointer(_uri))
	return uri, nil
}

func (i *ICoreWebView2NavigationStartingEventArgs) GetIsUserInitiated() (bool, error) {
	// Create int32 to hold bool result
	var _isUserInitiated int32
	hr, _, err := i.vtbl.GetIsUserInitiated.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_isUserInitiated)),
	)
	err = Error(hr, err)
	if err != nil {
		return false, err
	}
	// Get result and cleanup
	isUserInitiated := _isUserInitiated != 0
	return isUserInitiated, nil
}

func (i *ICoreWebView2NavigationStartingEventArgs) GetIsRedirected() (bool, error) {
	// Create int32 to hold bool result
	var _isRedirected int32
	hr, _, err := i.vtbl.GetIsRedirected.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_isRedirected)),
	)
	err = Error(hr, err)
	if err != nil {
		return false, err
	}
	// Get result and cleanup
	isRedirected := _isRedirected != 0
	return isRedirected, nil
}

func (i *ICoreWebView2NavigationStartingEventArgs) GetRequestHeaders() (*ICoreWebView2HttpRequestHeaders, error) {
	var requestHeaders *ICoreWebView2HttpRequestHeaders
	hr, _, err := i.vtbl.GetRequestHeaders.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&requestHeaders)),
	)
	err = Error(hr, err)
	if err != nil {
		return nil, err
	}
	return requestHeaders, nil
}

func (i *ICoreWebView2NavigationStartingEventArgs) GetCancel() (bool, error) {
	// Create int32 to hold bool result
	var _cancel int32
	hr, _, err := i.vtbl.GetCancel.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_cancel)),
	)
	err = Error(hr, err)
	if err != nil {
		return false, err
	}
	// Get result and cleanup
	cancel := _cancel != 0
	return cancel, nil
}

func (i *ICoreWebView2NavigationStartingEventArgs) PutCancel(cancel bool) error {
	hr, _, err := i.vtbl.PutCancel.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&cancel)),
	)
	return Error(hr, err)
}

func (i *ICoreWebView2NavigationStartingEventArgs) GetNavigationId() (uint64, error) {
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
func (i *ICoreWebView2NavigationStartingEventArgs) GetNavigationKind() (COREWEBVIEW2_NAVIGATION_KIND, error) {

	var navigation_kind COREWEBVIEW2_NAVIGATION_KIND

	hr, _, err := i.vtbl.GetNavigationKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&navigation_kind)),
	)
	err = Error(hr, err)
	if err != nil {
		return COREWEBVIEW2_NAVIGATION_KIND_RELOAD, err
	}
	return navigation_kind, err
}
