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

	e.controller.vtbl.PutBounds.Call(
		uintptr(unsafe.Pointer(e.controller)),
		uintptr(bounds.Left),
		uintptr(bounds.Top),
		uintptr(bounds.Right),
		uintptr(bounds.Bottom),
	)
}
func (i *ICoreWebView2Controller) PutZoomFactor(zoomFactor float64) error {
	var err error
	bits := math.Float64bits(zoomFactor)
	first := uint32(bits >> 32)
	second := uint32(bits & 0xFFFFFFFF)
	_, _, err = i.vtbl.PutZoomFactor.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(second),
		uintptr(first),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
