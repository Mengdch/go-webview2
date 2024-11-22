//go:build windows

package edge

import (
	"unsafe"
)

type ICoreWebView2ContextMenuTargetVtbl struct {
	IUnknownVtbl
	GetKind                    ComProc //获取当前上下文菜单目标的类型（例如是链接、图片、视频等）。返回一个枚举值，表示具体的目标类型
	GetIsEditable              ComProc //判断当前的上下文菜单目标是否可编辑，通常适用于输入框或其他用户可编辑的元素
	GetIsRequestedForMainFrame ComProc //确定当前的上下文菜单是否是为主框架 (Main Frame) 请求的。如果是嵌入的 iframe，这个方法可以帮助区分是否来自主框架
	GetPageUri                 ComProc //获取当前页面的 URI，表示用户点击的位置所在的页面
	GetFrameUri                ComProc //获取当前框架的 URI，通常用于判断是否是在某个嵌入的 iframe 中点击的。
	GetHasLinkUri              ComProc //判断当前点击的目标是否是一个链接。如果返回 true，则表示目标是一个可点击的超链接。
	GetLinkUri                 ComProc //获取当前点击的链接的 URI。如果目标是一个超链接，这个方法会返回该链接的地址
	GetHasLinkText             ComProc //判断当前点击的目标是否有链接文本。如果返回 true，表示目标是一个包含文本的超链接。
	GetLinkText                ComProc //获取当前链接的文本内容，通常用于获取超链接显示的文本。
	GetHasSourceUri            ComProc //判断当前目标是否有资源的 URI。例如，点击图片时，这个方法会判断是否有图片的源地址。
	GetSourceUri               ComProc //获取当前目标的资源 URI。如果目标是图片、视频等资源文件，这个方法会返回资源的地址。
	GetHasSelection            ComProc //判断当前目标是否包含选中的文本。如果用户在页面上选择了部分文本，这个方法会返回 true。
	GetSelectionText           ComProc //获取当前选中的文本内容。如果用户在页面上选择了部分文字，这个方法会返回所选的文本。
}

type ICoreWebView2ContextMenuTarget struct {
	vtbl *ICoreWebView2ContextMenuTargetVtbl
}

func (i *ICoreWebView2ContextMenuTarget) AddRef() uintptr {
	refCounter, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2ContextMenuTarget) GetKind() (COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND, error) {

	var value COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND

	hr, _, err := i.vtbl.GetKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	return value, Error(hr, err)
}

func (i *ICoreWebView2ContextMenuTarget) GetIsEditable() (bool, error) {
	// Create int32 to hold bool result
	var _value int32

	hr, _, err := i.vtbl.GetIsEditable.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	err = Error(hr, err)
	if err != nil {
		return false, err
	}
	// Get result and cleanup
	value := _value != 0
	return value, nil
}

func (i *ICoreWebView2ContextMenuTarget) GetIsRequestedForMainFrame() (bool, error) {
	// Create int32 to hold bool result
	var _value int32

	hr, _, err := i.vtbl.GetIsRequestedForMainFrame.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	err = Error(hr, err)
	if err != nil {
		return false, err
	}

	// Get result and cleanup
	value := _value != 0
	return value, nil
}

func (i *ICoreWebView2ContextMenuTarget) GetPageUri() (string, error) {
	// Create *uint16 to hold result
	var _value *uint16

	hr, _, err := i.vtbl.GetPageUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	err = Error(hr, err)
	if err != nil {
		return "", err
	}

	// Get result and cleanup
	value := UTF16PtrToString(_value)
	CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

func (i *ICoreWebView2ContextMenuTarget) GetFrameUri() (string, error) {
	// Create *uint16 to hold result
	var _value *uint16

	hr, _, err := i.vtbl.GetFrameUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	err = Error(hr, err)
	if err != nil {
		return "", err
	}

	// Get result and cleanup
	value := UTF16PtrToString(_value)
	CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

func (i *ICoreWebView2ContextMenuTarget) GetHasLinkUri() (bool, error) {
	// Create int32 to hold bool result
	var _value int32

	hr, _, err := i.vtbl.GetHasLinkUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	err = Error(hr, err)
	if err != nil {
		return false, err
	}
	// Get result and cleanup
	value := _value != 0
	return value, nil
}

func (i *ICoreWebView2ContextMenuTarget) GetLinkUri() (string, error) {
	// Create *uint16 to hold result
	var _value *uint16

	hr, _, err := i.vtbl.GetLinkUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	err = Error(hr, err)
	if err != nil {
		return "", err
	}

	// Get result and cleanup
	value := UTF16PtrToString(_value)
	CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

func (i *ICoreWebView2ContextMenuTarget) GetHasLinkText() (bool, error) {
	// Create int32 to hold bool result
	var _value int32

	hr, _, err := i.vtbl.GetHasLinkText.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	err = Error(hr, err)
	if err != nil {
		return false, err
	}
	// Get result and cleanup
	value := _value != 0
	return value, nil
}

func (i *ICoreWebView2ContextMenuTarget) GetLinkText() (string, error) {
	// Create *uint16 to hold result
	var _value *uint16

	hr, _, err := i.vtbl.GetLinkText.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	err = Error(hr, err)
	if err != nil {
		return "", err
	}
	// Get result and cleanup
	value := UTF16PtrToString(_value)
	CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

func (i *ICoreWebView2ContextMenuTarget) GetHasSourceUri() (bool, error) {
	// Create int32 to hold bool result
	var _value int32

	hr, _, err := i.vtbl.GetHasSourceUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	err = Error(hr, err)
	if err != nil {
		return false, err
	}
	// Get result and cleanup
	value := _value != 0
	return value, nil
}

func (i *ICoreWebView2ContextMenuTarget) GetSourceUri() (string, error) {
	// Create *uint16 to hold result
	var _value *uint16

	hr, _, err := i.vtbl.GetSourceUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	err = Error(hr, err)
	if err != nil {
		return "", err
	}
	// Get result and cleanup
	value := UTF16PtrToString(_value)
	CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

func (i *ICoreWebView2ContextMenuTarget) GetHasSelection() (bool, error) {
	// Create int32 to hold bool result
	var _value int32

	hr, _, err := i.vtbl.GetHasSelection.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	err = Error(hr, err)
	if err != nil {
		return false, err
	}
	// Get result and cleanup
	value := _value != 0
	return value, nil
}

func (i *ICoreWebView2ContextMenuTarget) GetSelectionText() (string, error) {
	// Create *uint16 to hold result
	var _value *uint16

	hr, _, err := i.vtbl.GetSelectionText.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	err = Error(hr, err)
	if err != nil {
		return "", err
	}
	// Get result and cleanup
	value := UTF16PtrToString(_value)
	CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}
