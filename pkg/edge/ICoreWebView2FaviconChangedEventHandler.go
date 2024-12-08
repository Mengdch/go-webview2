//go:build windows

package edge

type _ICoreWebView2FaviconChangedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2FaviconChangedEventHandler struct {
	vtbl *_ICoreWebView2FaviconChangedEventHandlerVtbl
	impl _ICoreWebView2FaviconChangedEventHandlerImpl
}

func _ICoreWebView2FaviconChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2FaviconChangedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2FaviconChangedEventHandlerIUnknownAddRef(this *ICoreWebView2FaviconChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2FaviconChangedEventHandlerIUnknownRelease(this *ICoreWebView2FaviconChangedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2FaviconChangedEventHandlerInvoke(this *ICoreWebView2FaviconChangedEventHandler, sender *ICoreWebView2, args *IUnknown) uintptr {
	return this.impl.FaviconChanged(sender, args)
}

type _ICoreWebView2FaviconChangedEventHandlerImpl interface {
	_IUnknownImpl
	FaviconChanged(sender *ICoreWebView2, args *IUnknown) uintptr
}

var _ICoreWebView2FaviconChangedEventHandlerFn = _ICoreWebView2FaviconChangedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2FaviconChangedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2FaviconChangedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2FaviconChangedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2FaviconChangedEventHandlerInvoke),
}

func newICoreWebView2FaviconChangedEventHandler(impl _ICoreWebView2FaviconChangedEventHandlerImpl) *ICoreWebView2FaviconChangedEventHandler {
	return &ICoreWebView2FaviconChangedEventHandler{
		vtbl: &_ICoreWebView2FaviconChangedEventHandlerFn,
		impl: impl,
	}
}
