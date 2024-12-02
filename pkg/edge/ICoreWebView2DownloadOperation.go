//go:build windows

package edge

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2DownloadOperationVtbl struct {
	_IUnknownVtbl
	AddBytesReceivedChanged       ComProc
	RemoveBytesReceivedChanged    ComProc
	AddEstimatedEndTimeChanged    ComProc
	RemoveEstimatedEndTimeChanged ComProc
	AddStateChanged               ComProc
	RemoveStateChanged            ComProc
	GetUri                        ComProc
	GetContentDisposition         ComProc
	GetMimeType                   ComProc
	GetTotalBytesToReceive        ComProc
	GetBytesReceived              ComProc
	GetEstimatedEndTime           ComProc
	GetResultFilePath             ComProc
	GetState                      ComProc
	GetInterruptReason            ComProc
	Cancel                        ComProc
	Pause                         ComProc
	Resume                        ComProc
	GetCanResume                  ComProc
}

type ICoreWebView2DownloadOperation struct {
	vtbl *ICoreWebView2DownloadOperationVtbl
}

func (i *ICoreWebView2DownloadOperation) AddBytesReceivedChanged(eventHandler *ICoreWebView2BytesReceivedChangedEventHandler) (EventRegistrationToken, error) {
	var token EventRegistrationToken
	hr, _, err := i.vtbl.AddBytesReceivedChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	err = Error(hr, err)
	if err != nil {
		return EventRegistrationToken{}, err
	}
	return token, nil
}

func (i *ICoreWebView2DownloadOperation) RemoveBytesReceivedChanged(token EventRegistrationToken) error {

	hr, _, err := i.vtbl.RemoveBytesReceivedChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}

func (i *ICoreWebView2DownloadOperation) RemoveEstimatedEndTimeChanged(token EventRegistrationToken) error {

	hr, _, err := i.vtbl.RemoveEstimatedEndTimeChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}

func (i *ICoreWebView2DownloadOperation) AddStateChanged(eventHandler *ICoreWebView2StateChangedEventHandler) (EventRegistrationToken, error) {
	var token EventRegistrationToken
	hr, _, err := i.vtbl.AddStateChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	err = Error(hr, err)
	if err != nil {
		return EventRegistrationToken{}, err
	}
	return token, nil
}

func (i *ICoreWebView2DownloadOperation) RemoveStateChanged(token EventRegistrationToken) error {

	hr, _, err := i.vtbl.RemoveStateChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}

func (i *ICoreWebView2DownloadOperation) GetUri() (string, error) {
	// Create *uint16 to hold result
	var _uri *uint16

	hr, _, err := i.vtbl.GetUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_uri)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return "", syscall.Errno(hr)
	}
	// Get result and cleanup
	uri := UTF16PtrToString(_uri)
	CoTaskMemFree(unsafe.Pointer(_uri))
	return uri, err
}

func (i *ICoreWebView2DownloadOperation) GetContentDisposition() (string, error) {
	// Create *uint16 to hold result
	var _contentDisposition *uint16

	hr, _, err := i.vtbl.GetContentDisposition.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_contentDisposition)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return "", syscall.Errno(hr)
	}
	// Get result and cleanup
	contentDisposition := UTF16PtrToString(_contentDisposition)
	CoTaskMemFree(unsafe.Pointer(_contentDisposition))
	return contentDisposition, err
}

func (i *ICoreWebView2DownloadOperation) GetMimeType() (string, error) {
	// Create *uint16 to hold result
	var _mimeType *uint16

	hr, _, err := i.vtbl.GetMimeType.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_mimeType)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return "", syscall.Errno(hr)
	}
	// Get result and cleanup
	mimeType := UTF16PtrToString(_mimeType)
	CoTaskMemFree(unsafe.Pointer(_mimeType))
	return mimeType, err
}

func (i *ICoreWebView2DownloadOperation) GetTotalBytesToReceive() (int64, error) {

	var totalBytesToReceive int64

	hr, _, err := i.vtbl.GetTotalBytesToReceive.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&totalBytesToReceive)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return 0, syscall.Errno(hr)
	}
	return totalBytesToReceive, err
}

func (i *ICoreWebView2DownloadOperation) GetBytesReceived() (int64, error) {

	var bytesReceived int64

	hr, _, err := i.vtbl.GetBytesReceived.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&bytesReceived)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return 0, syscall.Errno(hr)
	}
	return bytesReceived, err
}

func (i *ICoreWebView2DownloadOperation) GetEstimatedEndTime() (string, error) {
	// Create *uint16 to hold result
	var _estimatedEndTime *uint16

	hr, _, err := i.vtbl.GetEstimatedEndTime.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_estimatedEndTime)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return "", syscall.Errno(hr)
	}
	// Get result and cleanup
	estimatedEndTime := UTF16PtrToString(_estimatedEndTime)
	CoTaskMemFree(unsafe.Pointer(_estimatedEndTime))
	return estimatedEndTime, err
}

func (i *ICoreWebView2DownloadOperation) GetResultFilePath() (string, error) {
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

func (i *ICoreWebView2DownloadOperation) GetState() (COREWEBVIEW2_DOWNLOAD_STATE, error) {

	var downloadState COREWEBVIEW2_DOWNLOAD_STATE

	hr, _, err := i.vtbl.GetState.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&downloadState)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return 0, syscall.Errno(hr)
	}
	return downloadState, err
}

func (i *ICoreWebView2DownloadOperation) GetInterruptReason() (COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON, error) {

	var interruptReason COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON

	hr, _, err := i.vtbl.GetInterruptReason.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&interruptReason)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return 0, syscall.Errno(hr)
	}
	return interruptReason, err
}

func (i *ICoreWebView2DownloadOperation) Cancel() error {

	hr, _, err := i.vtbl.Cancel.Call(
		uintptr(unsafe.Pointer(i)),
	)
	return Error(hr, err)
}

func (i *ICoreWebView2DownloadOperation) Pause() error {

	hr, _, err := i.vtbl.Pause.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}

func (i *ICoreWebView2DownloadOperation) Resume() error {

	hr, _, err := i.vtbl.Resume.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}

func (i *ICoreWebView2DownloadOperation) GetCanResume() (bool, error) {
	// Create int32 to hold bool result
	var _canResume int32

	hr, _, err := i.vtbl.GetCanResume.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_canResume)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return false, syscall.Errno(hr)
	}
	// Get result and cleanup
	canResume := _canResume != 0
	return canResume, err
}
