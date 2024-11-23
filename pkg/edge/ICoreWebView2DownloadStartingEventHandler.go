//go:build windows

package edge

type ICoreWebView2DownloadStartingEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2DownloadStartingEventHandler struct {
	vtbl *ICoreWebView2DownloadStartingEventHandlerVtbl
	impl ICoreWebView2DownloadStartingEventHandlerImpl
}

func ICoreWebView2DownloadStartingEventHandlerIUnknownQueryInterface(this *ICoreWebView2DownloadStartingEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func ICoreWebView2DownloadStartingEventHandlerIUnknownAddRef(this *ICoreWebView2DownloadStartingEventHandler) uintptr {
	return this.impl.AddRef()
}

func ICoreWebView2DownloadStartingEventHandlerIUnknownRelease(this *ICoreWebView2DownloadStartingEventHandler) uintptr {
	return this.impl.Release()
}

func ICoreWebView2DownloadStartingEventHandlerInvoke(this *ICoreWebView2DownloadStartingEventHandler, sender *ICoreWebView2, args *ICoreWebView2DownloadStartingEventArgs) uintptr {
	return this.impl.DownloadStarting(sender, args)
}

type ICoreWebView2DownloadStartingEventHandlerImpl interface {
	IUnknownImpl
	DownloadStarting(sender *ICoreWebView2, args *ICoreWebView2DownloadStartingEventArgs) uintptr
}

var ICoreWebView2DownloadStartingEventHandlerFn = ICoreWebView2DownloadStartingEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(ICoreWebView2DownloadStartingEventHandlerIUnknownQueryInterface),
		NewComProc(ICoreWebView2DownloadStartingEventHandlerIUnknownAddRef),
		NewComProc(ICoreWebView2DownloadStartingEventHandlerIUnknownRelease),
	},
	NewComProc(ICoreWebView2DownloadStartingEventHandlerInvoke),
}

func newICoreWebView2DownloadStartingEventHandler(impl ICoreWebView2DownloadStartingEventHandlerImpl) *ICoreWebView2DownloadStartingEventHandler {
	return &ICoreWebView2DownloadStartingEventHandler{
		vtbl: &ICoreWebView2DownloadStartingEventHandlerFn,
		impl: impl,
	}
}
