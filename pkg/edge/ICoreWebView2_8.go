//go:build windows

package edge

import (
	"github.com/Mengdch/win"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type iCoreWebView2_8Vtbl struct {
	iCoreWebView2_7Vtbl
	AddIsMutedChanged                   ComProc
	RemoveIsMutedChanged                ComProc
	GetIsMuted                          ComProc
	SetIsMuted                          ComProc
	AddIsDocumentPlayingAudioChanged    ComProc
	RemoveIsDocumentPlayingAudioChanged ComProc
	GetIsDocumentPlayingAudio           ComProc
}

type ICoreWebView2_8 struct {
	vtbl *iCoreWebView2_8Vtbl
}

func (i *ICoreWebView2) GetICoreWebView2_8() *ICoreWebView2_8 {
	var result *ICoreWebView2_8

	iidICoreWebView2_8 := NewGUID("{E9632730-6E1E-43AB-B7B8-7B2C9E62E094}")
	_, _, _ = i.vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2_8)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (e *Chromium) GetICoreWebView2_8() *ICoreWebView2_8 {
	return e.webview.GetICoreWebView2_8()
}
func (i *ICoreWebView2_8) SetIsMuted(value bool) error {
	val := win.BoolToBOOL(value)
	hr, _, err := i.vtbl.SetIsMuted.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(val),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}
