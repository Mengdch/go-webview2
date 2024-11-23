//go:build windows

package edge

import (
	"unsafe"
)

type ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler struct {
	vtbl *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerVtbl
	impl ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerImpl
}

func (i *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler) AddRef() uintptr {
	refCounter, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerIUnknownQueryInterface(this *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerIUnknownAddRef(this *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerIUnknownRelease(this *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler) uintptr {
	return this.impl.Release()
}

func ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerInvoke(this *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler, errorCode uintptr, webView *ICoreWebView2CompositionController) uintptr {
	return this.impl.CreateCoreWebView2CompositionControllerCompleted(errorCode, webView)
}

type ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerImpl interface {
	IUnknownImpl
	CreateCoreWebView2CompositionControllerCompleted(errorCode uintptr, webView *ICoreWebView2CompositionController) uintptr
}

var ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerFn = ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerIUnknownQueryInterface),
		NewComProc(ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerIUnknownAddRef),
		NewComProc(ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerIUnknownRelease),
	},
	NewComProc(ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerInvoke),
}

func newICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler(impl ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerImpl) *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler {
	return &ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler{
		vtbl: &ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerFn,
		impl: impl,
	}
}
