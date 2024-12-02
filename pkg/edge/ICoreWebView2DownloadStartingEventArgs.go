//go:build windows

package edge

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2DownloadStartingEventArgsVtbl struct {
	_IUnknownVtbl
	GetDownloadOperation ComProc
	GetCancel            ComProc
	PutCancel            ComProc
	GetResultFilePath    ComProc
	PutResultFilePath    ComProc
	GetHandled           ComProc
	PutHandled           ComProc
	GetDeferral          ComProc
}

type ICoreWebView2DownloadStartingEventArgs struct {
	vtbl *ICoreWebView2DownloadStartingEventArgsVtbl
}

func (i *ICoreWebView2DownloadStartingEventArgs) GetDownloadOperation() (*ICoreWebView2DownloadOperation, error) {
	var downloadOperation *ICoreWebView2DownloadOperation
	hr, _, err := i.vtbl.GetDownloadOperation.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&downloadOperation)),
	)
	err = Error(hr, err)
	if err != nil {
		return nil, err
	}
	return downloadOperation, nil
}

func (i *ICoreWebView2DownloadStartingEventArgs) GetCancel() (bool, error) {
	// Create int32 to hold bool result
	var _cancel int32

	hr, _, err := i.vtbl.GetCancel.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_cancel)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return false, syscall.Errno(hr)
	}
	// Get result and cleanup
	cancel := _cancel != 0
	return cancel, err
}

func (i *ICoreWebView2DownloadStartingEventArgs) PutCancel(cancel bool) error {

	hr, _, err := i.vtbl.PutCancel.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&cancel)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}

func (i *ICoreWebView2DownloadStartingEventArgs) GetResultFilePath() (string, error) {
	// Create *uint16 to hold result
	var _resultFilePath *uint16

	hr, _, err := i.vtbl.GetResultFilePath.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_resultFilePath)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return "", syscall.Errno(hr)
	}
	// Get result and cleanup
	resultFilePath := UTF16PtrToString(_resultFilePath)
	CoTaskMemFree(unsafe.Pointer(_resultFilePath))
	return resultFilePath, err
}

func (i *ICoreWebView2DownloadStartingEventArgs) PutResultFilePath(resultFilePath string) error {

	// Convert string 'resultFilePath' to *uint16
	_resultFilePath, err := UTF16PtrFromString(resultFilePath)
	if err != nil {
		return err
	}

	hr, _, err := i.vtbl.PutResultFilePath.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_resultFilePath)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}

func (i *ICoreWebView2DownloadStartingEventArgs) GetHandled() (bool, error) {
	// Create int32 to hold bool result
	var _handled int32
	hr, _, err := i.vtbl.GetHandled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_handled)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return false, syscall.Errno(hr)
	}
	// Get result and cleanup
	handled := _handled != 0
	return handled, err
}

func (i *ICoreWebView2DownloadStartingEventArgs) PutHandled(handled bool) error {
	hr, _, err := i.vtbl.PutHandled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&handled)),
	)
	return Error(hr, err)
}

func (i *ICoreWebView2DownloadStartingEventArgs) GetDeferral() (*ICoreWebView2Deferral, error) {
	var deferral *ICoreWebView2Deferral
	hr, _, err := i.vtbl.GetDeferral.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&deferral)),
	)
	err = Error(hr, err)
	if err != nil {
		return nil, err
	}
	return deferral, nil
}
