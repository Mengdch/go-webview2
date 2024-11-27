//go:build windows

package edge

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2Settings7Vtbl struct {
	ICoreWebView2Settings6Vtbl
	GetHiddenPdfToolbarItems ComProc
	PutHiddenPdfToolbarItems ComProc
}

type ICoreWebView2Settings7 struct {
	Vtbl *ICoreWebView2Settings7Vtbl
}

func (i *ICoreWebView2Settings7) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2Settings) GetICoreWebView2Settings7() *ICoreWebView2Settings7 {
	var result *ICoreWebView2Settings7

	iidICoreWebView2Settings7 := NewGUID("{488dc902-35ef-42d2-bc7d-94b65c4bc49c}")
	_, _, _ = i.vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2Settings7)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (i *ICoreWebView2Settings7) GetHiddenPdfToolbarItems() (COREWEBVIEW2_PDF_TOOLBAR_ITEMS, error) {

	var hidden_pdf_toolbar_items COREWEBVIEW2_PDF_TOOLBAR_ITEMS

	hr, _, err := i.Vtbl.GetHiddenPdfToolbarItems.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&hidden_pdf_toolbar_items)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return 0, syscall.Errno(hr)
	}
	return hidden_pdf_toolbar_items, err
}

func (i *ICoreWebView2Settings7) PutHiddenPdfToolbarItems(hidden_pdf_toolbar_items COREWEBVIEW2_PDF_TOOLBAR_ITEMS) error {

	hr, _, err := i.Vtbl.PutHiddenPdfToolbarItems.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(hidden_pdf_toolbar_items),
	)
	return Error(hr, err)
}
