//go:build windows

package edge

type _ICoreWebView2NewWindowRequestedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2NewWindowRequestedEventHandler struct {
	vtbl *_ICoreWebView2NewWindowRequestedEventHandlerVtbl
	impl ICoreWebView2NewWindowRequestedEventHandlerImpl
}

func ICoreWebView2NewWindowRequestedEventHandlerIUnknownQueryInterface(this *ICoreWebView2NewWindowRequestedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func ICoreWebView2NewWindowRequestedEventHandlerIUnknownAddRef(this *ICoreWebView2NewWindowRequestedEventHandler) uintptr {
	return this.impl.AddRef()
}

func ICoreWebView2NewWindowRequestedEventHandlerIUnknownRelease(this *ICoreWebView2NewWindowRequestedEventHandler) uintptr {
	return this.impl.Release()
}

func ICoreWebView2NewWindowRequestedEventHandlerInvoke(this *ICoreWebView2NewWindowRequestedEventHandler, sender *ICoreWebView2, args *ICoreWebView2NewWindowRequestedEventArgs) uintptr {
	return this.impl.NewWindowRequested(sender, args)
}

type ICoreWebView2NewWindowRequestedEventHandlerImpl interface {
	IUnknownImpl
	NewWindowRequested(sender *ICoreWebView2, args *ICoreWebView2NewWindowRequestedEventArgs) uintptr
}

var ICoreWebView2NewWindowRequestedEventHandlerFn = _ICoreWebView2NewWindowRequestedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(ICoreWebView2NewWindowRequestedEventHandlerIUnknownQueryInterface),
		NewComProc(ICoreWebView2NewWindowRequestedEventHandlerIUnknownAddRef),
		NewComProc(ICoreWebView2NewWindowRequestedEventHandlerIUnknownRelease),
	},
	NewComProc(ICoreWebView2NewWindowRequestedEventHandlerInvoke),
}

func newICoreWebView2NewWindowRequestedEventHandler(impl ICoreWebView2NewWindowRequestedEventHandlerImpl) *ICoreWebView2NewWindowRequestedEventHandler {
	return &ICoreWebView2NewWindowRequestedEventHandler{
		vtbl: &ICoreWebView2NewWindowRequestedEventHandlerFn,
		impl: impl,
	}
}
