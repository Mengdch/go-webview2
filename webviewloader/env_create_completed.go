//go:build windows && !native_webview2loader

package webviewloader

import (
	"github.com/Mengdch/win"
	"github.com/wailsapp/go-webview2/pkg/combridge"
	"golang.org/x/sys/windows"
	"unsafe"
)

// HRESULT
//
// See https://docs.microsoft.com/en-us/windows/win32/seccrypto/common-hresult-values
type HRESULT int32

// ICoreWebView2Environment Represents the WebView2 Environment
//
// See https://docs.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment
type ICoreWebView2Environment = combridge.IUnknownImpl

// ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandler receives the WebView2Environment created using CreateCoreWebView2Environment.
type ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandler interface {
	// EnvironmentCompleted is invoked to receive the created WebView2Environment
	//
	// See https://docs.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2createcorewebview2environmentcompletedhandler?#invoke
	EnvironmentCompleted(errorCode HRESULT, createdEnvironment *ICoreWebView2Environment) HRESULT
}

type iCoreWebView2CreateCoreWebView2EnvironmentCompletedHandler interface {
	combridge.IUnknown
	ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandler
}
type iCoreWebView2CustomSchemeRegistration interface {
	combridge.IUnknown
	ICoreWebView2CustomSchemeRegistrationIn
}
type ICoreWebView2CustomSchemeRegistrationIn interface {
	GetSchemeName(name *uintptr) uintptr
	GetTreatAsSecure(name *uintptr) uintptr
	SetTreatAsSecure(name uintptr) uintptr
	GetAllowedOrigins(count, value *uintptr) uintptr
	SetAllowedOrigins(count, value uintptr) uintptr
	GetHasAuthorityComponent(name *uintptr) uintptr
	SetHasAuthorityComponent(name uintptr) uintptr
}

func init() {
	combridge.RegisterVTable[combridge.IUnknown, iCoreWebView2CreateCoreWebView2EnvironmentCompletedHandler](
		"{4e8a3389-c9d8-4bd2-b6b5-124fee6cc14d}",
		_iCoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerInvoke,
	)
	combridge.RegisterVTable[combridge.IUnknown, iCoreWebView2CustomSchemeRegistration](
		"{d60ac92c-37a6-4b26-a39e-95cfe59047bb}",
		_iCoreWebView2CreateCoreWebView2GetSchemeName,
		_iCoreWebView2CreateCoreWebView2GetTreatAsSecure,
		_iCoreWebView2CreateCoreWebView2SetTreatAsSecure,
		_iCoreWebView2CreateCoreWebView2GetAllowedOrigins,
		_iCoreWebView2CreateCoreWebView2SetAllowedOrigins,
		_iCoreWebView2CreateCoreWebView2GetHasAuthorityComponent,
		_iCoreWebView2CreateCoreWebView2SetHasAuthorityComponent,
	)
}

func _iCoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerInvoke(this uintptr, errorCode HRESULT, env *combridge.IUnknownImpl) uintptr {
	res := combridge.Resolve[iCoreWebView2CreateCoreWebView2EnvironmentCompletedHandler](this).EnvironmentCompleted(errorCode, env)
	return uintptr(res)
}
func _iCoreWebView2CreateCoreWebView2GetSchemeName(this uintptr, name *uintptr) uintptr {
	//fmt.Println("_iCoreWebView2CreateCoreWebView2GetSchemeName")
	res := combridge.Resolve[iCoreWebView2CustomSchemeRegistration](this).GetSchemeName(name)
	return res
}
func _iCoreWebView2CreateCoreWebView2GetTreatAsSecure(this uintptr, name *uintptr) uintptr { //BOOL* treatAsSecure
	//fmt.Println("_iCoreWebView2CreateCoreWebView2GetTreatAsSecure")
	toBOOL := win.BoolToBOOL(true)
	*name = uintptr(unsafe.Pointer(&toBOOL))
	return uintptr(windows.S_OK)
}
func _iCoreWebView2CreateCoreWebView2SetTreatAsSecure(this, name uintptr) uintptr {
	//fmt.Println("_iCoreWebView2CreateCoreWebView2SetTreatAsSecure")

	return uintptr(windows.S_OK)
}
func _iCoreWebView2CreateCoreWebView2GetAllowedOrigins(this uintptr, count, value *uintptr) uintptr {
	//fmt.Println("_iCoreWebView2CreateCoreWebView2GetAllowedOrigins")
	val := []string{} //"*.*"
	oc := len(val)
	if oc > 0 {
		ret, list := combridge.AllocUintptrObject(oc)
		for i, v := range val {
			sp := stringToOleString(v)
			list[i] = uintptr(unsafe.Pointer(sp))
		}
		*count = uintptr(unsafe.Pointer(&oc))
		*value = ret
	} else {
		*count, *value = 0, 0
	}
	return uintptr(windows.S_OK)
}
func _iCoreWebView2CreateCoreWebView2SetAllowedOrigins(this, count, value uintptr) uintptr {
	//fmt.Println("_iCoreWebView2CreateCoreWebView2SetAllowedOrigins")

	return uintptr(windows.S_OK)
}
func _iCoreWebView2CreateCoreWebView2GetHasAuthorityComponent(this uintptr, name *uintptr) uintptr { //BOOL* hasAuthorityComponent
	//fmt.Println("_iCoreWebView2CreateCoreWebView2GetHasAuthorityComponent")
	toBOOL := win.BoolToBOOL(true)
	*name = uintptr(unsafe.Pointer(&toBOOL))

	return uintptr(windows.S_OK)
}
func _iCoreWebView2CreateCoreWebView2SetHasAuthorityComponent(this, name uintptr) uintptr {
	//fmt.Println("_iCoreWebView2CreateCoreWebView2SetHasAuthorityComponent")

	return uintptr(windows.S_OK)
}

type CustomSchemeRegistration struct {
	scheme string
}
