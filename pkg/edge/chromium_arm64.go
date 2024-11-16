//go:build windows
// +build windows

package edge

import (
	"github.com/Mengdch/go-webview2/internal/w32"
	"golang.org/x/sys/windows"
	"math"
	"unsafe"
)

func (e *Chromium) SetSize(bounds w32.Rect) {
	if e.controller == nil {
		return
	}

	words := (*[2]uintptr)(unsafe.Pointer(&bounds))
	e.controller.vtbl.PutBounds.Call(
		uintptr(unsafe.Pointer(e.controller)),
		words[0],
		words[1],
	)
}
func (i *ICoreWebView2Controller) PutZoomFactor(zoomFactor float64) error {
	var err error
	bits := math.Float64bits(zoomFactor)
	_, _, err = i.vtbl.PutZoomFactor.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(bits),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
