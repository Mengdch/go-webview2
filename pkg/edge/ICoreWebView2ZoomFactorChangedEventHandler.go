//go:build windows

package edge

import (
	"unsafe"
)

type _ICoreWebView2ZoomFactorChangedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type iCoreWebView2ZoomFactorChangedEventHandler struct {
	vtbl *_ICoreWebView2ZoomFactorChangedEventHandlerVtbl
	impl ICoreWebView2ZoomFactorChangedEventHandlerImpl
}

func (i *iCoreWebView2ZoomFactorChangedEventHandler) AddRef() uintptr {
	refCounter, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func ICoreWebView2ZoomFactorChangedEventHandlerIUnknownQueryInterface(this *iCoreWebView2ZoomFactorChangedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func ICoreWebView2ZoomFactorChangedEventHandlerIUnknownAddRef(this *iCoreWebView2ZoomFactorChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func ICoreWebView2ZoomFactorChangedEventHandlerIUnknownRelease(this *iCoreWebView2ZoomFactorChangedEventHandler) uintptr {
	return this.impl.Release()
}

func ICoreWebView2ZoomFactorChangedEventHandlerInvoke(this *iCoreWebView2ZoomFactorChangedEventHandler, sender *ICoreWebView2Controller, args *IUnknown) uintptr {
	return this.impl.ZoomFactorChanged(sender, args)
}

type ICoreWebView2ZoomFactorChangedEventHandlerImpl interface {
	IUnknownImpl
	ZoomFactorChanged(sender *ICoreWebView2Controller, args *IUnknown) uintptr
}

var ICoreWebView2ZoomFactorChangedEventHandlerFn = _ICoreWebView2ZoomFactorChangedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(ICoreWebView2ZoomFactorChangedEventHandlerIUnknownQueryInterface),
		NewComProc(ICoreWebView2ZoomFactorChangedEventHandlerIUnknownAddRef),
		NewComProc(ICoreWebView2ZoomFactorChangedEventHandlerIUnknownRelease),
	},
	NewComProc(ICoreWebView2ZoomFactorChangedEventHandlerInvoke),
}

func newICoreWebView2ZoomFactorChangedEventHandler(impl ICoreWebView2ZoomFactorChangedEventHandlerImpl) *iCoreWebView2ZoomFactorChangedEventHandler {
	return &iCoreWebView2ZoomFactorChangedEventHandler{
		vtbl: &ICoreWebView2ZoomFactorChangedEventHandlerFn,
		impl: impl,
	}
}
