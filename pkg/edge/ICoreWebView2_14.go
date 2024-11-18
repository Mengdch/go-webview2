//go:build windows

package edge

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type iCoreWebView2_14Vtbl struct {
	iCoreWebView2_13Vtbl
	AddServerCertificateErrorDetected    ComProc
	RemoveServerCertificateErrorDetected ComProc
	ClearServerCertificateErrorActions   ComProc
}

type ICoreWebView2_14 struct {
	vtbl *iCoreWebView2_14Vtbl
}

func (i *ICoreWebView2) GetICoreWebView2_14() *ICoreWebView2_14 {
	var result *ICoreWebView2_14

	iidICoreWebView2_14 := NewGUID("{6DAA4F10-4A90-4753-8898-77C5DF534165}")
	_, _, _ = i.vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2_14)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (i *ICoreWebView2_14) AddServerCertificateErrorDetected(eventHandler *ICoreWebView2ServerCertificateErrorDetectedEventHandler) (EventRegistrationToken, error) {

	var token EventRegistrationToken

	hr, _, err := i.vtbl.AddServerCertificateErrorDetected.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return EventRegistrationToken{}, syscall.Errno(hr)
	}
	return token, err
}

func (i *ICoreWebView2_14) RemoveServerCertificateErrorDetected(token EventRegistrationToken) error {

	hr, _, err := i.vtbl.RemoveServerCertificateErrorDetected.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}

func (i *ICoreWebView2_14) ClearServerCertificateErrorActions(handler *ICoreWebView2ClearServerCertificateErrorActionsCompletedHandler) error {

	hr, _, err := i.vtbl.ClearServerCertificateErrorActions.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(handler)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}
