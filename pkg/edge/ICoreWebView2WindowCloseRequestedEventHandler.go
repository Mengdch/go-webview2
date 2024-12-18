//go:build windows

package edge

import (
	"unsafe"
)

type ICoreWebView2WindowCloseRequestedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2WindowCloseRequestedEventHandler struct {
	vtbl *ICoreWebView2WindowCloseRequestedEventHandlerVtbl
	impl ICoreWebView2WindowCloseRequestedEventHandlerImpl
}

func (i *ICoreWebView2WindowCloseRequestedEventHandler) AddRef() uintptr {
	refCounter, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func ICoreWebView2WindowCloseRequestedEventHandlerIUnknownQueryInterface(this *ICoreWebView2WindowCloseRequestedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func ICoreWebView2WindowCloseRequestedEventHandlerIUnknownAddRef(this *ICoreWebView2WindowCloseRequestedEventHandler) uintptr {
	return this.impl.AddRef()
}

func ICoreWebView2WindowCloseRequestedEventHandlerIUnknownRelease(this *ICoreWebView2WindowCloseRequestedEventHandler) uintptr {
	return this.impl.Release()
}

func ICoreWebView2WindowCloseRequestedEventHandlerInvoke(this *ICoreWebView2WindowCloseRequestedEventHandler, sender *ICoreWebView2, args *IUnknown) uintptr {
	return this.impl.WindowCloseRequested(sender, args)
}

type ICoreWebView2WindowCloseRequestedEventHandlerImpl interface {
	IUnknownImpl
	WindowCloseRequested(sender *ICoreWebView2, args *IUnknown) uintptr
}

var ICoreWebView2WindowCloseRequestedEventHandlerFn = ICoreWebView2WindowCloseRequestedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(ICoreWebView2WindowCloseRequestedEventHandlerIUnknownQueryInterface),
		NewComProc(ICoreWebView2WindowCloseRequestedEventHandlerIUnknownAddRef),
		NewComProc(ICoreWebView2WindowCloseRequestedEventHandlerIUnknownRelease),
	},
	NewComProc(ICoreWebView2WindowCloseRequestedEventHandlerInvoke),
}

func newICoreWebView2WindowCloseRequestedEventHandler(impl ICoreWebView2WindowCloseRequestedEventHandlerImpl) *ICoreWebView2WindowCloseRequestedEventHandler {
	return &ICoreWebView2WindowCloseRequestedEventHandler{
		vtbl: &ICoreWebView2WindowCloseRequestedEventHandlerFn,
		impl: impl,
	}
}
