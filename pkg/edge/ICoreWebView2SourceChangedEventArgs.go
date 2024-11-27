//go:build windows

package edge

type _ICoreWebView2SourceChangedEventArgsVtbl struct {
	_IUnknownVtbl
	GetIsNewDocument ComProc
}

type ICoreWebView2SourceChangedEventArgs struct {
	vtbl *_ICoreWebView2SourceChangedEventArgsVtbl
}
