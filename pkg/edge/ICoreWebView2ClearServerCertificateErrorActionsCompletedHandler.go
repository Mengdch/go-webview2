//go:build windows

package edge

import (
	"unsafe"
)

type ICoreWebView2ClearServerCertificateErrorActionsCompletedHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2ClearServerCertificateErrorActionsCompletedHandler struct {
	vtbl *ICoreWebView2ClearServerCertificateErrorActionsCompletedHandlerVtbl
	impl ICoreWebView2ClearServerCertificateErrorActionsCompletedHandlerImpl
}

func (i *ICoreWebView2ClearServerCertificateErrorActionsCompletedHandler) AddRef() uintptr {
	refCounter, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func ICoreWebView2ClearServerCertificateErrorActionsCompletedHandlerIUnknownQueryInterface(this *ICoreWebView2ClearServerCertificateErrorActionsCompletedHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func ICoreWebView2ClearServerCertificateErrorActionsCompletedHandlerIUnknownAddRef(this *ICoreWebView2ClearServerCertificateErrorActionsCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func ICoreWebView2ClearServerCertificateErrorActionsCompletedHandlerIUnknownRelease(this *ICoreWebView2ClearServerCertificateErrorActionsCompletedHandler) uintptr {
	return this.impl.Release()
}

func ICoreWebView2ClearServerCertificateErrorActionsCompletedHandlerInvoke(this *ICoreWebView2ClearServerCertificateErrorActionsCompletedHandler, errorCode uintptr) uintptr {
	return this.impl.ClearServerCertificateErrorActionsCompleted(errorCode)
}

type ICoreWebView2ClearServerCertificateErrorActionsCompletedHandlerImpl interface {
	IUnknownImpl
	ClearServerCertificateErrorActionsCompleted(errorCode uintptr) uintptr
}

var ICoreWebView2ClearServerCertificateErrorActionsCompletedHandlerFn = ICoreWebView2ClearServerCertificateErrorActionsCompletedHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(ICoreWebView2ClearServerCertificateErrorActionsCompletedHandlerIUnknownQueryInterface),
		NewComProc(ICoreWebView2ClearServerCertificateErrorActionsCompletedHandlerIUnknownAddRef),
		NewComProc(ICoreWebView2ClearServerCertificateErrorActionsCompletedHandlerIUnknownRelease),
	},
	NewComProc(ICoreWebView2ClearServerCertificateErrorActionsCompletedHandlerInvoke),
}

func newICoreWebView2ClearServerCertificateErrorActionsCompletedHandler(impl ICoreWebView2ClearServerCertificateErrorActionsCompletedHandlerImpl) *ICoreWebView2ClearServerCertificateErrorActionsCompletedHandler {
	return &ICoreWebView2ClearServerCertificateErrorActionsCompletedHandler{
		vtbl: &ICoreWebView2ClearServerCertificateErrorActionsCompletedHandlerFn,
		impl: impl,
	}
}
