//go:build windows

package edge

import (
	"unsafe"
)

type _ICoreWebView2FocusChangedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2FocusChangedEventHandler struct {
	vtbl *_ICoreWebView2FocusChangedEventHandlerVtbl
	impl ICoreWebView2FocusChangedEventHandlerImpl
}

func (i *ICoreWebView2FocusChangedEventHandler) AddRef() uintptr {
	refCounter, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func ICoreWebView2FocusChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2FocusChangedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func ICoreWebView2FocusChangedEventHandlerIUnknownAddRef(this *ICoreWebView2FocusChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func ICoreWebView2FocusChangedEventHandlerIUnknownRelease(this *ICoreWebView2FocusChangedEventHandler) uintptr {
	return this.impl.Release()
}

func ICoreWebView2FocusChangedEventHandlerInvoke(this *ICoreWebView2FocusChangedEventHandler, sender *ICoreWebView2Controller, args *IUnknown) uintptr {
	return this.impl.FocusChanged(sender, args)
}

type ICoreWebView2FocusChangedEventHandlerImpl interface {
	IUnknownImpl
	FocusChanged(sender *ICoreWebView2Controller, args *IUnknown) uintptr
}

var ICoreWebView2FocusChangedEventHandlerFn = _ICoreWebView2FocusChangedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(ICoreWebView2FocusChangedEventHandlerIUnknownQueryInterface),
		NewComProc(ICoreWebView2FocusChangedEventHandlerIUnknownAddRef),
		NewComProc(ICoreWebView2FocusChangedEventHandlerIUnknownRelease),
	},
	NewComProc(ICoreWebView2FocusChangedEventHandlerInvoke),
}

func newICoreWebView2FocusChangedEventHandler(impl ICoreWebView2FocusChangedEventHandlerImpl) *ICoreWebView2FocusChangedEventHandler {
	return &ICoreWebView2FocusChangedEventHandler{
		vtbl: &ICoreWebView2FocusChangedEventHandlerFn,
		impl: impl,
	}
}
