//go:build windows
// +build windows

package edge

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync/atomic"
	"syscall"
	"time"
	"unsafe"

	"github.com/Mengdch/go-webview2/internal/w32"
	"github.com/Mengdch/go-webview2/webviewloader"
	"golang.org/x/sys/windows"
)

type Rect = w32.Rect

type Chromium struct {
	hwnd                             uintptr
	controller                       *ICoreWebView2Controller
	webview                          *ICoreWebView2
	inited                           uintptr
	envCompleted                     *iCoreWebView2CreateCoreWebView2EnvironmentCompletedHandler
	controllerCompleted              *iCoreWebView2CreateCoreWebView2ControllerCompletedHandler
	compositionCompleted             *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler
	webMessageReceived               *iCoreWebView2WebMessageReceivedEventHandler
	containsFullScreenElementChanged *ICoreWebView2ContainsFullScreenElementChangedEventHandler
	permissionRequested              *iCoreWebView2PermissionRequestedEventHandler
	webResourceRequested             *iCoreWebView2WebResourceRequestedEventHandler
	acceleratorKeyPressed            *ICoreWebView2AcceleratorKeyPressedEventHandler
	navigationCompleted              *ICoreWebView2NavigationCompletedEventHandler
	documentTitleChanged             *ICoreWebView2DocumentTitleChangedEventHandler
	downloadStart                    *ICoreWebView2DownloadStartingEventHandler
	downloadStateChanged             *ICoreWebView2StateChangedEventHandler
	downloadReceivedChanged          *ICoreWebView2BytesReceivedChangedEventHandler
	sourceChanged                    *ICoreWebView2SourceChangedEventHandler
	newWindow                        *ICoreWebView2NewWindowRequestedEventHandler
	processFailed                    *ICoreWebView2ProcessFailedEventHandler
	certificateError                 *ICoreWebView2ServerCertificateErrorDetectedEventHandler
	clearCertificate                 *ICoreWebView2ClearServerCertificateErrorActionsCompletedHandler
	callDelMethodCompleted           *ICoreWebView2CallDevToolsProtocolMethodCompletedHandler
	zoomFactorChanged                *iCoreWebView2ZoomFactorChangedEventHandler
	contextMenuRequested             *ICoreWebView2ContextMenuRequestedEventHandler
	windowCloseRequested             *ICoreWebView2WindowCloseRequestedEventHandler
	responseReceived                 *ICoreWebView2WebResourceResponseReceivedEventHandler
	trySuspendCompleted              *ICoreWebView2TrySuspendCompletedHandler

	environment            *ICoreWebView2Environment
	padding                Rect
	webview2RuntimeVersion string

	// Settings
	Debug                 bool
	DataPath              string
	BrowserPath           string
	AdditionalBrowserArgs []string
	Scheme                []string

	// permissions
	permissions      map[CoreWebView2PermissionKind]CoreWebView2PermissionState
	globalPermission *CoreWebView2PermissionState

	// Callbacks
	MessageCallback                             func(string)
	MessageWithAdditionalObjectsCallback        func(message string, sender *ICoreWebView2, args *ICoreWebView2WebMessageReceivedEventArgs)
	WebResourceRequestedCallback                func(request *ICoreWebView2WebResourceRequest, args *ICoreWebView2WebResourceRequestedEventArgs)
	NavigationCompletedCallback                 func(sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs)
	DocumentTitleChangedCallback                func(sender *ICoreWebView2, args *IUnknown)
	FaviconChangedCallback                      func(sender *ICoreWebView2, args *IUnknown)
	GetFaviconCompletedCallback                 func(errorCode uintptr, faviconStream *IStream)
	SourceChangedCallback                       func(sender *ICoreWebView2, args *ICoreWebView2SourceChangedEventArgs)
	DownloadStartCallback                       func(sender *ICoreWebView2, args *ICoreWebView2DownloadStartingEventArgs)
	DownloadStateChangedCallback                func(sender *ICoreWebView2DownloadOperation, args *IUnknown)
	DownloadReceivedChangedCallback             func(sender *ICoreWebView2DownloadOperation, args *IUnknown)
	NewWindowCallback                           func(sender *ICoreWebView2, args *ICoreWebView2NewWindowRequestedEventArgs)
	HistoryChangedCallback                      func(sender *ICoreWebView2, args *IUnknown)
	ProcessFailedCallback                       func(sender *ICoreWebView2, args *ICoreWebView2ProcessFailedEventArgs)
	ContainsFullScreenElementChangedCallback    func(sender *ICoreWebView2, args *ICoreWebView2ContainsFullScreenElementChangedEventArgs)
	AcceleratorKeyCallback                      func(uint, int) bool
	CertificateError                            func(sender *ICoreWebView2, args *ICoreWebView2ServerCertificateErrorDetectedEventArgs)
	ClearCertificate                            func(errorCode uintptr)
	PrintPdfCompleted                           func(errorCode uintptr, isSuccessful bool)
	FocusChangedCallback                        func(sender *ICoreWebView2Controller, args *IUnknown)
	ZoomFactorChangedCallback                   func(sender *ICoreWebView2Controller, args *IUnknown)
	CallDevToolsProtocolMethodCompletedCallback func(errorCode uintptr, returnObjectAsJson uintptr)
	ContextMenuRequestedCallback                func(sender *ICoreWebView2, args *ICoreWebView2ContextMenuRequestedEventArgs)
	WebResourceResponseReceivedCallback         func(sender *ICoreWebView2, args *ICoreWebView2WebResourceResponseReceivedEventArgs)
	InitCompletedCallback                       func()

	focus bool
	start time.Time
}

func NewChromium() *Chromium {
	e := &Chromium{}
	/*
	 All these handlers are passed to native code through syscalls with 'uintptr(unsafe.Pointer(handler))' and we know
	 that a pointer to those will be kept in the native code. Furthermore these handlers als contain pointer to other Go
	 structs like the vtable.
	 This violates the unsafe.Pointer rule '(4) Conversion of a Pointer to a uintptr when calling syscall.Syscall.' because
	 theres no guarantee that Go doesn't move these objects.
	 AFAIK currently the Go runtime doesn't move HEAP objects, so we should be safe with these handlers. But they don't
	 guarantee it, because in the future Go might use a compacting GC.
	 There's a proposal to add a runtime.Pin function, to prevent moving pinned objects, which would allow to easily fix
	 this issue by just pinning the handlers. The https://go-review.googlesource.com/c/go/+/367296/ should land in Go 1.19.
	*/
	e.envCompleted = newICoreWebView2CreateCoreWebView2EnvironmentCompletedHandler(e)
	runtime.KeepAlive(e.envCompleted)
	e.controllerCompleted = newICoreWebView2CreateCoreWebView2ControllerCompletedHandler(e)
	runtime.KeepAlive(e.controllerCompleted)
	e.webMessageReceived = newICoreWebView2WebMessageReceivedEventHandler(e)
	e.permissionRequested = newICoreWebView2PermissionRequestedEventHandler(e)
	e.webResourceRequested = newICoreWebView2WebResourceRequestedEventHandler(e)
	e.acceleratorKeyPressed = newICoreWebView2AcceleratorKeyPressedEventHandler(e)
	e.navigationCompleted = newICoreWebView2NavigationCompletedEventHandler(e)
	e.documentTitleChanged = newICoreWebView2DocumentTitleChangedEventHandler(e)
	e.downloadStart = newICoreWebView2DownloadStartingEventHandler(e)
	e.downloadStateChanged = newICoreWebView2StateChangedEventHandler(e)
	e.downloadReceivedChanged = newICoreWebView2BytesReceivedChangedEventHandler(e)
	e.sourceChanged = newICoreWebView2SourceChangedEventHandler(e)
	e.newWindow = newICoreWebView2NewWindowRequestedEventHandler(e)
	e.processFailed = newICoreWebView2ProcessFailedEventHandler(e)
	e.containsFullScreenElementChanged = newICoreWebView2ContainsFullScreenElementChangedEventHandler(e)
	e.certificateError = newICoreWebView2ServerCertificateErrorDetectedEventHandler(e)
	e.clearCertificate = newICoreWebView2ClearServerCertificateErrorActionsCompletedHandler(e)
	e.zoomFactorChanged = newICoreWebView2ZoomFactorChangedEventHandler(e)
	e.callDelMethodCompleted = newICoreWebView2CallDevToolsProtocolMethodCompletedHandler(e)
	e.contextMenuRequested = newICoreWebView2ContextMenuRequestedEventHandler(e)
	e.windowCloseRequested = newICoreWebView2WindowCloseRequestedEventHandler(e)
	e.responseReceived = newICoreWebView2WebResourceResponseReceivedEventHandler(e)
	e.trySuspendCompleted = newICoreWebView2TrySuspendCompletedHandler(e)
	/*
		// Pinner seems to panic in some cases as reported on Discord, maybe during shutdown when GC detects pinned objects
		// to be released that have not been unpinned.
		// It would also be better to use our ComBridge for this event handlers implementation instead of pinning them.
		// So all COM Implementations on the go-side use the same code.
		var pinner runtime.Pinner
		pinner.Pin(e.envCompleted)
		pinner.Pin(e.controllerCompleted)
		pinner.Pin(e.webMessageReceived)
		pinner.Pin(e.permissionRequested)
		pinner.Pin(e.webResourceRequested)
		pinner.Pin(e.acceleratorKeyPressed)
		pinner.Pin(e.navigationCompleted)
		pinner.Pin(e.processFailed)
		pinner.Pin(e.containsFullScreenElementChanged)
	*/
	e.permissions = make(map[CoreWebView2PermissionKind]CoreWebView2PermissionState)

	return e
}

func (e *Chromium) Valid() bool {
	return atomic.LoadUintptr(&e.inited) != 0
}
func (e *Chromium) Embed(hwnd uintptr) bool {
	var err error
	e.hwnd = hwnd
	dataPath := e.DataPath
	if dataPath == "" {
		currentExePath := make([]uint16, windows.MAX_PATH)
		_, err := windows.GetModuleFileName(windows.Handle(0), &currentExePath[0], windows.MAX_PATH)
		if err != nil {
			// What to do here?
			return false
		}
		currentExeName := filepath.Base(windows.UTF16ToString(currentExePath))
		dataPath = filepath.Join(os.Getenv("AppData"), currentExeName)
	}

	if e.BrowserPath != "" {
		if _, err := os.Stat(e.BrowserPath); errors.Is(err, os.ErrNotExist) {
			fmt.Println("Browser path %s does not exist", e.BrowserPath)
			return false
		}
	}

	browserArgs := strings.Join(e.AdditionalBrowserArgs, " ")
	e.start = time.Now()
	if err := createCoreWebView2EnvironmentWithOptions(e.BrowserPath, dataPath, e.envCompleted, browserArgs, e.Scheme); err != nil {
		fmt.Println("Error calling Webview2Loader: %v", err)
		return false
	}

	e.webview2RuntimeVersion, err = webviewloader.GetAvailableCoreWebView2BrowserVersionString(e.BrowserPath)
	if err != nil {
		fmt.Println("Error getting Webview2 runtime version: %v", err)
		return false
	}

	return true
}

func (e *Chromium) SetPadding(padding Rect) {
	if e.padding.Top == padding.Top && e.padding.Bottom == padding.Bottom &&
		e.padding.Left == padding.Left && e.padding.Right == padding.Right {

		return
	}

	e.padding = padding
	e.Resize()
}

func (e *Chromium) Resize() {
	if e.hwnd == 0 {
		return
	}

	var bounds w32.Rect
	w32.User32GetClientRect.Call(e.hwnd, uintptr(unsafe.Pointer(&bounds)))

	bounds.Top += e.padding.Top
	bounds.Bottom -= e.padding.Bottom
	bounds.Left += e.padding.Left
	bounds.Right -= e.padding.Right

	e.SetSize(bounds)
}

func (e *Chromium) Navigate(url string) {
	e.webview.vtbl.Navigate.Call(
		uintptr(unsafe.Pointer(e.webview)),
		uintptr(unsafe.Pointer(windows.StringToUTF16Ptr(url))),
	)
}

func (e *Chromium) NavigateToString(content string) {
	e.webview.vtbl.NavigateToString.Call(
		uintptr(unsafe.Pointer(e.webview)),
		uintptr(unsafe.Pointer(windows.StringToUTF16Ptr(content))),
	)
}

func (e *Chromium) Init(script string) {
	e.webview.vtbl.AddScriptToExecuteOnDocumentCreated.Call(
		uintptr(unsafe.Pointer(e.webview)),
		uintptr(unsafe.Pointer(windows.StringToUTF16Ptr(script))),
		0,
	)
}

func (e *Chromium) Eval(script string) {

	if e.webview == nil {
		return
	}

	_script, err := windows.UTF16PtrFromString(script)
	if err != nil {
		fmt.Println("Eval", err)
	} else {
		//fmt.Println("Eval", script)
	}
	hr, _, err := e.webview.vtbl.ExecuteScript.Call(
		uintptr(unsafe.Pointer(e.webview)),
		uintptr(unsafe.Pointer(_script)),
		0,
	)
	err = Error(hr, err)
	if err != nil {
		fmt.Println("ExecuteScript", err.Error())
	}
}

func (e *Chromium) Show() error {
	return e.controller.PutIsVisible(true)
}

func (e *Chromium) Hide() error {
	return e.controller.PutIsVisible(false)
}

func (e *Chromium) QueryInterface(_, _ uintptr) uintptr {
	return 0
}

func (e *Chromium) AddRef() uintptr {
	return 1
}

func (e *Chromium) Release() uintptr {
	return 1
}

func (e *Chromium) EnvironmentCompleted(res uintptr, env *ICoreWebView2Environment) uintptr {
	//fmt.Println("EnvironmentCompleted", time.Now().Sub(e.start))
	if int32(res) < 0 {
		fmt.Println("EnvironmentCompleted", "Creating environment failed with %08x: %s", res, syscall.Errno(res))
	}
	env.vtbl.AddRef.Call(uintptr(unsafe.Pointer(env)))
	e.environment = env
	e.start = time.Now()
	env.vtbl.CreateCoreWebView2Controller.Call(
		uintptr(unsafe.Pointer(env)),
		e.hwnd,
		uintptr(unsafe.Pointer(e.controllerCompleted)),
	)
	//if e.webview2RuntimeVersion
	//err := e.webview.GetICoreWebView2Environment3().CreateCoreWebView2CompositionController(HWND(e.hwnd), e.compositionCompleted)
	//if err != nil {
	//	fmt.Println("CreateCoreWebView2CompositionController", err.Error())
	//}
	return 0
}

func (e *Chromium) CreateCoreWebView2CompositionControllerCompleted(errorCode uintptr, webView *ICoreWebView2CompositionController) uintptr {
	return 0
}
func (e *Chromium) CreateCoreWebView2ControllerCompleted(res uintptr, controller *ICoreWebView2Controller) uintptr {
	//fmt.Println("CreateCoreWebView2ControllerCompleted", time.Now().Sub(e.start))
	if int32(res) < 0 {
		fmt.Printf("CreateCoreWebView2ControllerCompleted, Creating controller failed with %08x: %s\n", res, syscall.Errno(res))
		return 1
	}
	controller.vtbl.AddRef.Call(uintptr(unsafe.Pointer(controller)))
	e.controller = controller

	var token _EventRegistrationToken
	controller.vtbl.GetCoreWebView2.Call(
		uintptr(unsafe.Pointer(controller)),
		uintptr(unsafe.Pointer(&e.webview)),
	)
	e.webview.vtbl.AddRef.Call(
		uintptr(unsafe.Pointer(e.webview)),
	)
	e.webview.vtbl.AddWebMessageReceived.Call(
		uintptr(unsafe.Pointer(e.webview)),
		uintptr(unsafe.Pointer(e.webMessageReceived)),
		uintptr(unsafe.Pointer(&token)),
	)
	e.webview.vtbl.AddPermissionRequested.Call(
		uintptr(unsafe.Pointer(e.webview)),
		uintptr(unsafe.Pointer(e.permissionRequested)),
		uintptr(unsafe.Pointer(&token)),
	)
	e.webview.vtbl.AddWebResourceRequested.Call(
		uintptr(unsafe.Pointer(e.webview)),
		uintptr(unsafe.Pointer(e.webResourceRequested)),
		uintptr(unsafe.Pointer(&token)),
	)
	e.webview.vtbl.AddDocumentTitleChanged.Call(
		uintptr(unsafe.Pointer(e.webview)),
		uintptr(unsafe.Pointer(e.documentTitleChanged)),
		uintptr(unsafe.Pointer(&token)),
	)
	e.GetICoreWebView2_4().AddDownloadStarting(e.downloadStart)
	e.webview.GetICoreWebView2_2().AddWebResourceResponseReceived(e.responseReceived)
	e.webview.vtbl.AddSourceChanged.Call(
		uintptr(unsafe.Pointer(e.webview)),
		uintptr(unsafe.Pointer(e.sourceChanged)),
		uintptr(unsafe.Pointer(&token)),
	)
	e.webview.vtbl.AddNewWindowRequested.Call(
		uintptr(unsafe.Pointer(e.webview)),
		uintptr(unsafe.Pointer(e.newWindow)),
		uintptr(unsafe.Pointer(&token)),
	)
	e.webview.vtbl.AddNavigationCompleted.Call(
		uintptr(unsafe.Pointer(e.webview)),
		uintptr(unsafe.Pointer(e.navigationCompleted)),
		uintptr(unsafe.Pointer(&token)),
	)
	e.webview.vtbl.AddProcessFailed.Call(
		uintptr(unsafe.Pointer(e.webview)),
		uintptr(unsafe.Pointer(e.processFailed)),
		uintptr(unsafe.Pointer(&token)),
	)
	e.webview.vtbl.AddContainsFullScreenElementChanged.Call(
		uintptr(unsafe.Pointer(e.webview)),
		uintptr(unsafe.Pointer(e.containsFullScreenElementChanged)),
		uintptr(unsafe.Pointer(&token)),
	)
	e.webview.GetICoreWebView2_14().AddServerCertificateErrorDetected(e.certificateError)
	e.controller.AddZoomFactorChanged(e.zoomFactorChanged)
	//e.controller.AddMoveFocusRequested(e.moveFocusRequested)

	e.controller.AddAcceleratorKeyPressed(e.acceleratorKeyPressed, &token)
	e.webview.GetICoreWebView2_11().AddContextMenuRequested(e.contextMenuRequested)
	e.webview.AddWindowCloseRequested(e.windowCloseRequested)

	atomic.StoreUintptr(&e.inited, 1)
	e.Init("window.external={invoke:s=>window.chrome.webview.postMessage(s)}")
	if e.InitCompletedCallback != nil {
		e.InitCompletedCallback()
	}
	return 0
}

func (e *Chromium) ContainsFullScreenElementChanged(sender *ICoreWebView2, args *ICoreWebView2ContainsFullScreenElementChangedEventArgs) uintptr {
	if e.ContainsFullScreenElementChangedCallback != nil {
		e.ContainsFullScreenElementChangedCallback(sender, args)
	}
	return 0
}

func (e *Chromium) MessageReceived(sender *ICoreWebView2, args *ICoreWebView2WebMessageReceivedEventArgs) uintptr {
	var _message *uint16
	args.vtbl.TryGetWebMessageAsString.Call(
		uintptr(unsafe.Pointer(args)),
		uintptr(unsafe.Pointer(&_message)),
	)

	message := w32.Utf16PtrToString(_message)

	if hasCapability(e.webview2RuntimeVersion, GetAdditionalObjects) {
		obj, err := args.GetAdditionalObjects()
		if err != nil {
			fmt.Println("MessageReceived", err)
		}

		if obj != nil && e.MessageWithAdditionalObjectsCallback != nil {
			defer obj.Release()
			e.MessageWithAdditionalObjectsCallback(message, sender, args)
		} else if e.MessageCallback != nil {
			e.MessageCallback(message)
		}
	} else if e.MessageCallback != nil {
		e.MessageCallback(message)
	}

	sender.vtbl.PostWebMessageAsString.Call(
		uintptr(unsafe.Pointer(sender)),
		uintptr(unsafe.Pointer(_message)),
	)
	windows.CoTaskMemFree(unsafe.Pointer(_message))
	return 0
}

func (e *Chromium) SetPermission(kind CoreWebView2PermissionKind, state CoreWebView2PermissionState) {
	e.permissions[kind] = state
}

func (e *Chromium) SetBackgroundColour(R, G, B, A uint8) {
	controller := e.GetController()
	controller2 := controller.GetICoreWebView2Controller2()

	backgroundCol := COREWEBVIEW2_COLOR{
		A: A,
		R: R,
		G: G,
		B: B,
	}

	// WebView2 only has 0 and 255 as valid values.
	if backgroundCol.A > 0 && backgroundCol.A < 255 {
		backgroundCol.A = 255
	}

	err := controller2.PutDefaultBackgroundColor(backgroundCol)
	if err != nil {
		fmt.Println("SetBackgroundColour", err)
	}
}

func (e *Chromium) SetGlobalPermission(state CoreWebView2PermissionState) {
	e.globalPermission = &state
}

func (e *Chromium) PermissionRequested(_ *ICoreWebView2, args *iCoreWebView2PermissionRequestedEventArgs) uintptr {
	var kind CoreWebView2PermissionKind
	args.vtbl.GetPermissionKind.Call(
		uintptr(unsafe.Pointer(args)),
		uintptr(kind),
	)
	var result CoreWebView2PermissionState
	if e.globalPermission != nil {
		result = *e.globalPermission
	} else {
		var ok bool
		result, ok = e.permissions[kind]
		if !ok {
			result = CoreWebView2PermissionStateDefault
		}
	}
	args.vtbl.PutState.Call(
		uintptr(unsafe.Pointer(args)),
		uintptr(result),
	)
	return 0
}

func (e *Chromium) WebResourceRequested(sender *ICoreWebView2, args *ICoreWebView2WebResourceRequestedEventArgs) uintptr {
	req, err := args.GetRequest()
	if err != nil {
		fmt.Println("WebResourceRequested", err)
		return 0
	}
	defer req.Release()

	if e.WebResourceRequestedCallback != nil {
		e.WebResourceRequestedCallback(req, args)
	}
	return 0
}

func (e *Chromium) AddWebResourceRequestedFilter(filter string, ctx COREWEBVIEW2_WEB_RESOURCE_CONTEXT) {
	err := e.webview.AddWebResourceRequestedFilter(filter, ctx)
	if err != nil {
		fmt.Println("AddWebResourceRequestedFilter", err)
	}
}

func (e *Chromium) Environment() *ICoreWebView2Environment {
	return e.environment
}
func (e *Chromium) ClearServerCertificateErrorActionsCompleted(errorCode uintptr) uintptr {
	if e.ClearCertificate != nil {
		e.ClearCertificate(errorCode)
	}
	return 0
}
func (e *Chromium) PrintToPdfCompleted(errorCode uintptr, isSuccessful bool) uintptr {
	if e.PrintPdfCompleted != nil {
		e.PrintPdfCompleted(errorCode, isSuccessful)
	}
	return 0
}
func (e *Chromium) ServerCertificateErrorDetected(sender *ICoreWebView2, args *ICoreWebView2ServerCertificateErrorDetectedEventArgs) uintptr {
	if e.CertificateError != nil {
		e.CertificateError(sender, args)
	}
	return 0
}

// AcceleratorKeyPressed is called when an accelerator key is pressed.
// If the AcceleratorKeyCallback method has been set, it will defer handling of the keypress
// to the callback. That callback returns a bool indicating if the event was handled.
func (e *Chromium) AcceleratorKeyPressed(sender *ICoreWebView2Controller, args *ICoreWebView2AcceleratorKeyPressedEventArgs) uintptr {
	if e.AcceleratorKeyCallback == nil {
		return 0
	}
	eventKind, _ := args.GetKeyEventKind()
	if eventKind == COREWEBVIEW2_KEY_EVENT_KIND_KEY_DOWN ||
		eventKind == COREWEBVIEW2_KEY_EVENT_KIND_SYSTEM_KEY_DOWN {
		virtualKey, _ := args.GetVirtualKey()
		lParam, _ := args.GetKeyEventLParam()
		status, _ := args.GetPhysicalKeyStatus()
		if !status.WasKeyDown {
			args.PutHandled(e.AcceleratorKeyCallback(virtualKey, lParam))
			return 0
		}
	}
	args.PutHandled(false)
	return 0
}
func (e *Chromium) FocusChanged(sender *ICoreWebView2Controller, args *IUnknown) uintptr {
	e.focus = !e.focus
	if e.FocusChangedCallback != nil {
		e.FocusChangedCallback(sender, args)
	}
	return 0
}
func (e *Chromium) GetSettings() (*ICoreWebView2Settings, error) {
	return e.webview.GetSettings()
}

func (e *Chromium) GetController() *ICoreWebView2Controller {
	return e.controller
}

func boolToInt(input bool) int {
	if input {
		return 1
	}
	return 0
}

func (e *Chromium) NavigationCompleted(sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) uintptr {
	if e.NavigationCompletedCallback != nil {
		e.NavigationCompletedCallback(sender, args)
	}
	return 0
}

func (e *Chromium) DocumentTitleChanged(sender *ICoreWebView2, args *IUnknown) uintptr {
	if e.DocumentTitleChangedCallback != nil {
		e.DocumentTitleChangedCallback(sender, args)
	}
	return 0
}
func (e *Chromium) FaviconChanged(sender *ICoreWebView2, args *IUnknown) uintptr {
	if e.FaviconChangedCallback != nil {
		e.FaviconChangedCallback(sender, args)
	}
	return 0
}
func (e *Chromium) GetFaviconCompleted(errorCode uintptr, faviconStream *IStream) uintptr {
	if e.GetFaviconCompletedCallback != nil {
		e.GetFaviconCompletedCallback(errorCode, faviconStream)
	}
	return 0
}
func (e *Chromium) SourceChanged(sender *ICoreWebView2, args *ICoreWebView2SourceChangedEventArgs) uintptr {
	if e.SourceChangedCallback != nil {
		e.SourceChangedCallback(sender, args)
	}
	return 0
}
func (e *Chromium) DownloadStarting(sender *ICoreWebView2, args *ICoreWebView2DownloadStartingEventArgs) uintptr {
	if e.DownloadStartCallback != nil {
		e.DownloadStartCallback(sender, args)
	}
	return 0
}
func (e *Chromium) StateChanged(sender *ICoreWebView2DownloadOperation, args *IUnknown) uintptr {
	if e.DownloadStateChangedCallback != nil {
		e.DownloadStateChangedCallback(sender, args)
	}
	return 0
}
func (e *Chromium) BytesReceivedChanged(sender *ICoreWebView2DownloadOperation, args *IUnknown) uintptr {
	if e.DownloadReceivedChangedCallback != nil {
		e.DownloadReceivedChangedCallback(sender, args)
	}
	return 0
}

func (e *Chromium) HistoryChanged(sender *ICoreWebView2, args *IUnknown) uintptr {
	if e.HistoryChangedCallback != nil {
		e.HistoryChangedCallback(sender, args)
	}
	return 0
}

func (e *Chromium) ProcessFailed(sender *ICoreWebView2, args *ICoreWebView2ProcessFailedEventArgs) uintptr {
	if e.ProcessFailedCallback != nil {
		e.ProcessFailedCallback(sender, args)
	}
	return 0
}

func (e *Chromium) NotifyParentWindowPositionChanged() error {
	//It looks like the wndproc function is called before the controller initialization is complete.
	//Because of this the controller is nil
	if e.controller == nil {
		return nil
	}
	return e.controller.NotifyParentWindowPositionChanged()
}

func (e *Chromium) Focus() {
	err := e.controller.MoveFocus(COREWEBVIEW2_MOVE_FOCUS_REASON_PROGRAMMATIC)
	if err != nil {
		fmt.Println("Focus", err)
	}
}

func (e *Chromium) GetZoomFactor() float64 {
	val, err := e.controller.GetZoomFactor()
	if err != nil {
		fmt.Println("GetZoomFactor", err)
	}
	return val
}
func (e *Chromium) PutZoomFactor(zoomFactor float64) {
	err := e.controller.PutZoomFactor(zoomFactor)
	if err != nil {
		fmt.Println("PutZoomFactor", err)
	}
}

func (e *Chromium) OpenDevToolsWindow() {
	e.webview.OpenDevToolsWindow()
}
func (e *Chromium) Title() (string, error) {
	return e.webview.Title()
}
func (e *Chromium) Url() (string, error) {
	return e.webview.Url()
}

func (e *Chromium) Reload() {
	e.webview.Reload()
}

func (e *Chromium) Stop() {
	e.webview.Stop()
}

func (e *Chromium) GoBack() {
	e.webview.GoBack()
}
func (e *Chromium) GoForward() {
	e.webview.GoForward()
}
func (e *Chromium) CallDevToolsProtocolMethod(methodName, parametersAsJson string) {
	e.webview.CallDevToolsProtocolMethod(methodName, parametersAsJson, e.callDelMethodCompleted)
}

func (e *Chromium) HasCapability(c Capability) bool {
	return hasCapability(e.webview2RuntimeVersion, c)
}

func (e *Chromium) GetIsSwipeNavigationEnabled() (bool, error) {
	if !hasCapability(e.webview2RuntimeVersion, SwipeNavigation) {
		return false, UnsupportedCapabilityError
	}
	webview2Settings, err := e.webview.GetSettings()
	if err != nil {
		return false, err
	}
	webview2Settings6 := webview2Settings.GetICoreWebView2Settings6()
	var result bool
	result, err = webview2Settings6.GetIsSwipeNavigationEnabled()
	if err != windows.DS_S_SUCCESS {
		return false, err
	}
	return result, nil
}

func (e *Chromium) PutIsSwipeNavigationEnabled(enabled bool) error {
	if !hasCapability(e.webview2RuntimeVersion, SwipeNavigation) {
		return UnsupportedCapabilityError
	}
	webview2Settings, err := e.webview.GetSettings()
	if err != nil {
		return err
	}
	webview2Settings6 := webview2Settings.GetICoreWebView2Settings6()
	err = webview2Settings6.PutIsSwipeNavigationEnabled(enabled)
	if err != windows.DS_S_SUCCESS {
		return err
	}
	return nil
}

func (e *Chromium) AllowExternalDrag(allow bool) error {
	if !hasCapability(e.webview2RuntimeVersion, AllowExternalDrop) {
		return UnsupportedCapabilityError
	}
	controller := e.GetController()
	controller4 := controller.GetICoreWebView2Controller4()
	err := controller4.PutAllowExternalDrop(allow)
	if err != windows.DS_S_SUCCESS {
		return err
	}
	return nil
}

func (e *Chromium) SetMuted(allow bool) error {
	return e.webview.SetMuted(allow)
}

func (e *Chromium) NewWindowRequested(sender *ICoreWebView2, args *ICoreWebView2NewWindowRequestedEventArgs) uintptr {
	if e.NewWindowCallback != nil {
		e.NewWindowCallback(sender, args)
	}
	return 0
}

func (e *Chromium) ClearCertificateCompleted() *ICoreWebView2ClearServerCertificateErrorActionsCompletedHandler {
	return e.clearCertificate
}
func (e *Chromium) DownloadStateChanged() *ICoreWebView2StateChangedEventHandler {
	return e.downloadStateChanged

}
func (e *Chromium) DownloadReceivedChanged() *ICoreWebView2BytesReceivedChangedEventHandler {
	return e.downloadReceivedChanged
}

func (e *Chromium) GetAllowExternalDrag() (bool, error) {
	if !hasCapability(e.webview2RuntimeVersion, AllowExternalDrop) {
		return false, UnsupportedCapabilityError
	}
	controller := e.GetController()
	controller4 := controller.GetICoreWebView2Controller4()
	result, err := controller4.GetAllowExternalDrop()
	if err != windows.DS_S_SUCCESS {
		return false, err
	}
	return result, nil
}
func (e *Chromium) ZoomFactorChanged(sender *ICoreWebView2Controller, args *IUnknown) uintptr {
	if e.ZoomFactorChangedCallback != nil {
		e.ZoomFactorChangedCallback(sender, args)
	}
	return 0
}
func (e *Chromium) CallDevToolsProtocolMethodCompleted(errorCode uintptr, returnObjectAsJson uintptr) uintptr {
	if e.CallDevToolsProtocolMethodCompletedCallback != nil {
		e.CallDevToolsProtocolMethodCompletedCallback(errorCode, returnObjectAsJson)
	}
	return 0
}
func (e *Chromium) ContextMenuRequested(sender *ICoreWebView2, args *ICoreWebView2ContextMenuRequestedEventArgs) uintptr {
	if e.ContextMenuRequestedCallback != nil {
		e.ContextMenuRequestedCallback(sender, args)
	}
	return 0
}
func (e *Chromium) Focused() bool {
	return e.focus
}
func (e *Chromium) Close() {
	if e.controller != nil {
		e.controller.Close()
	}
}
func (e *Chromium) WindowCloseRequested(sender *ICoreWebView2, args *IUnknown) uintptr {
	url, err := sender.Url()
	fmt.Println("WindowCloseRequested", url, err)
	return 0
}
func (e *Chromium) WebResourceResponseReceived(sender *ICoreWebView2, args *ICoreWebView2WebResourceResponseReceivedEventArgs) uintptr {
	if e.WebResourceResponseReceivedCallback != nil {
		e.WebResourceResponseReceivedCallback(sender, args)
	}
	return 0
}
func (e *Chromium) TrySuspendCompleted(errorCode uintptr, isSuccessful bool) uintptr {
	fmt.Println("TrySuspendCompleted", errorCode, isSuccessful)
	return 0
}
