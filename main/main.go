package main

import (
	"flag"
	"fmt"
	"github.com/Mengdch/go-webview2/pkg/edge"
	"github.com/Mengdch/go-webview2/webviewloader"
	"github.com/Mengdch/win"
	"runtime"
	"syscall"
	"unsafe"
)

var procMap = make(map[win.HWND]uintptr)

func main() {
	du := flag.String("du", "", "本地库")
	flag.Parse()

	fp, err := webviewloader.GetAvailableCoreWebView2BrowserVersionString(*du)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(fp)
	}
	runtime.LockOSThread()
	className := "testClass"
	classNamePtr, err := syscall.UTF16PtrFromString(className)
	if err != nil {
		fmt.Println(err)
		return
	}
	windowName := "testWindow"
	windowNamePtr, err := syscall.UTF16PtrFromString(windowName)
	if err != nil {
		fmt.Println(err)
		return
	}
	classViewName := "testViewClass"
	windowViewName := "testViewWindow"
	classViewNamePtr, err := syscall.UTF16PtrFromString(classViewName)
	if err != nil {
		fmt.Println(err)
		return
	}
	windowViewNamePtr, err := syscall.UTF16PtrFromString(windowViewName)
	if err != nil {
		fmt.Println(err)
		return
	}
	hInst := win.GetModuleHandle(nil)
	wndClass := win.WNDCLASSEX{
		Style:         win.CS_HREDRAW | win.CS_VREDRAW,
		LpfnWndProc:   syscall.NewCallbackCDecl(classMsgProc),
		HInstance:     hInst,
		LpszClassName: classNamePtr,
		HCursor:       win.LoadCursor(0, win.MAKEINTRESOURCE(win.IDC_ARROW)),
		HbrBackground: win.GetSysColorBrush(win.COLOR_INACTIVEBORDER),
	}
	wndClass.CbSize = uint32(unsafe.Sizeof(wndClass))
	win.RegisterClassEx(&wndClass)
	wndClass = win.WNDCLASSEX{
		Style:         win.CS_DBLCLKS,
		LpfnWndProc:   syscall.NewCallbackCDecl(classMsgProc),
		HInstance:     hInst,
		LpszClassName: classViewNamePtr,
		HbrBackground: win.GetSysColorBrush(win.COLOR_WINDOW),
		HCursor:       win.LoadCursor(0, win.MAKEINTRESOURCE(win.IDC_ARROW)),
	}
	wndClass.CbSize = uint32(unsafe.Sizeof(wndClass))
	win.RegisterClassEx(&wndClass)

	hWnd := win.CreateWindowEx(0, classNamePtr, windowNamePtr, win.WS_OVERLAPPEDWINDOW, 0, 0, 1600, 900,
		0, 0, hInst, unsafe.Pointer(nil))
	if hWnd == 0 {
		fmt.Println("CreateWindow")
		return
	}
	ps := func(hWnd win.HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr {
		return win.DefWindowProc(hWnd, msg, wParam, lParam)
	}
	procMap[hWnd] = syscall.NewCallbackCDecl(ps)
	win.ShowWindow(hWnd, win.SW_SHOW)
	var parentRect win.RECT
	win.GetClientRect(hWnd, &parentRect)
	vWnd := win.CreateWindowEx(0, classViewNamePtr, windowViewNamePtr,
		win.WS_CHILD|win.WS_CLIPSIBLINGS|win.WS_CLIPCHILDREN|win.WS_VISIBLE, 0, 0, parentRect.Width(),
		parentRect.Height(), hWnd, 0, hInst, unsafe.Pointer(nil))
	ps = func(hWnd win.HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr {
		return win.DefWindowProc(hWnd, msg, wParam, lParam)
	}
	procMap[hWnd] = syscall.NewCallbackCDecl(ps)
	chromium := edge.NewChromium()
	chromium.InitCompletedCallback = func() {
		fmt.Println("InitCompletedCallback")
		chromium.OpenDevToolsWindow()
		chromium.Resize()
		chromium.Navigate("https://www.google.com")
	}
	if !chromium.Embed(uintptr(vWnd)) {
		fmt.Println("Embed error")
		return
	}
	fmt.Println("GetMessage")
	msg := (*win.MSG)(unsafe.Pointer(win.GlobalAlloc(0, unsafe.Sizeof(win.MSG{}))))
	defer win.GlobalFree(win.HGLOBAL(unsafe.Pointer(msg)))
	for win.GetMessage(msg, 0, 0, 0) > 0 {
		win.TranslateMessage(msg)
		win.DispatchMessage(msg)
	}
}
func classMsgProc(hWnd win.HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr {
	if v, e := procMap[hWnd]; e {
		return win.CallWindowProc(v, hWnd, msg, wParam, lParam)
	}
	return win.DefWindowProc(hWnd, msg, wParam, lParam)
}
