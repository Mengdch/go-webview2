//go:build windows

package edge

import (
	"unsafe"
)

type _ICoreWebView2GetFaviconCompletedHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2GetFaviconCompletedHandler struct {
	vtbl *_ICoreWebView2GetFaviconCompletedHandlerVtbl
	impl ICoreWebView2GetFaviconCompletedHandlerImpl
}

func (i *ICoreWebView2GetFaviconCompletedHandler) AddRef() uintptr {
	refCounter, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func ICoreWebView2GetFaviconCompletedHandlerIUnknownQueryInterface(this *ICoreWebView2GetFaviconCompletedHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func ICoreWebView2GetFaviconCompletedHandlerIUnknownAddRef(this *ICoreWebView2GetFaviconCompletedHandler) uintptr {
	return this.impl.AddRef()
}

func ICoreWebView2GetFaviconCompletedHandlerIUnknownRelease(this *ICoreWebView2GetFaviconCompletedHandler) uintptr {
	return this.impl.Release()
}

func ICoreWebView2GetFaviconCompletedHandlerInvoke(this *ICoreWebView2GetFaviconCompletedHandler, errorCode uintptr, faviconStream *IStream) uintptr {
	return this.impl.GetFaviconCompleted(errorCode, faviconStream)
}

type ICoreWebView2GetFaviconCompletedHandlerImpl interface {
	IUnknownImpl
	GetFaviconCompleted(errorCode uintptr, faviconStream *IStream) uintptr
}

var ICoreWebView2GetFaviconCompletedHandlerFn = _ICoreWebView2GetFaviconCompletedHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(ICoreWebView2GetFaviconCompletedHandlerIUnknownQueryInterface),
		NewComProc(ICoreWebView2GetFaviconCompletedHandlerIUnknownAddRef),
		NewComProc(ICoreWebView2GetFaviconCompletedHandlerIUnknownRelease),
	},
	NewComProc(ICoreWebView2GetFaviconCompletedHandlerInvoke),
}

func newICoreWebView2GetFaviconCompletedHandler(impl ICoreWebView2GetFaviconCompletedHandlerImpl) *ICoreWebView2GetFaviconCompletedHandler {
	return &ICoreWebView2GetFaviconCompletedHandler{
		vtbl: &ICoreWebView2GetFaviconCompletedHandlerFn,
		impl: impl,
	}
}
