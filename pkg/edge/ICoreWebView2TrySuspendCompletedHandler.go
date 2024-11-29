//go:build windows

package edge

import (
	"unsafe"
)

type ICoreWebView2TrySuspendCompletedHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2TrySuspendCompletedHandler struct {
	vtbl *ICoreWebView2TrySuspendCompletedHandlerVtbl
	impl ICoreWebView2TrySuspendCompletedHandlerImpl
}

func (i *ICoreWebView2TrySuspendCompletedHandler) AddRef() uintptr {
	refCounter, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func ICoreWebView2TrySuspendCompletedHandlerIUnknownQueryInterface(this *ICoreWebView2TrySuspendCompletedHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func ICoreWebView2TrySuspendCompletedHandlerIUnknownAddRef(this *ICoreWebView2TrySuspendCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func ICoreWebView2TrySuspendCompletedHandlerIUnknownRelease(this *ICoreWebView2TrySuspendCompletedHandler) uintptr {
	return this.impl.Release()
}

func ICoreWebView2TrySuspendCompletedHandlerInvoke(this *ICoreWebView2TrySuspendCompletedHandler, errorCode uintptr, isSuccessful bool) uintptr {
	return this.impl.TrySuspendCompleted(errorCode, isSuccessful)
}

type ICoreWebView2TrySuspendCompletedHandlerImpl interface {
	IUnknownImpl
	TrySuspendCompleted(errorCode uintptr, isSuccessful bool) uintptr
}

var ICoreWebView2TrySuspendCompletedHandlerFn = ICoreWebView2TrySuspendCompletedHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(ICoreWebView2TrySuspendCompletedHandlerIUnknownQueryInterface),
		NewComProc(ICoreWebView2TrySuspendCompletedHandlerIUnknownAddRef),
		NewComProc(ICoreWebView2TrySuspendCompletedHandlerIUnknownRelease),
	},
	NewComProc(ICoreWebView2TrySuspendCompletedHandlerInvoke),
}

func newICoreWebView2TrySuspendCompletedHandler(impl ICoreWebView2TrySuspendCompletedHandlerImpl) *ICoreWebView2TrySuspendCompletedHandler {
	return &ICoreWebView2TrySuspendCompletedHandler{
		vtbl: &ICoreWebView2TrySuspendCompletedHandlerFn,
		impl: impl,
	}
}
