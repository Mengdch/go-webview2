//go:build windows

package edge

import (
	"unsafe"
)

type ICoreWebView2StateChangedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2StateChangedEventHandler struct {
	vtbl *ICoreWebView2StateChangedEventHandlerVtbl
	impl ICoreWebView2StateChangedEventHandlerImpl
}

func (i *ICoreWebView2StateChangedEventHandler) AddRef() uintptr {
	refCounter, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func ICoreWebView2StateChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2StateChangedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func ICoreWebView2StateChangedEventHandlerIUnknownAddRef(this *ICoreWebView2StateChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func ICoreWebView2StateChangedEventHandlerIUnknownRelease(this *ICoreWebView2StateChangedEventHandler) uintptr {
	return this.impl.Release()
}

func ICoreWebView2StateChangedEventHandlerInvoke(this *ICoreWebView2StateChangedEventHandler, sender *ICoreWebView2DownloadOperation, args *IUnknown) uintptr {
	return this.impl.StateChanged(sender, args)
}

type ICoreWebView2StateChangedEventHandlerImpl interface {
	IUnknownImpl
	StateChanged(sender *ICoreWebView2DownloadOperation, args *IUnknown) uintptr
}

var ICoreWebView2StateChangedEventHandlerFn = ICoreWebView2StateChangedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(ICoreWebView2StateChangedEventHandlerIUnknownQueryInterface),
		NewComProc(ICoreWebView2StateChangedEventHandlerIUnknownAddRef),
		NewComProc(ICoreWebView2StateChangedEventHandlerIUnknownRelease),
	},
	NewComProc(ICoreWebView2StateChangedEventHandlerInvoke),
}

func newICoreWebView2StateChangedEventHandler(impl ICoreWebView2StateChangedEventHandlerImpl) *ICoreWebView2StateChangedEventHandler {
	return &ICoreWebView2StateChangedEventHandler{
		vtbl: &ICoreWebView2StateChangedEventHandlerFn,
		impl: impl,
	}
}
