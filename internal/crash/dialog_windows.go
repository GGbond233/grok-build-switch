//go:build windows

package crash

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	mbOK            = 0x00000000
	mbIconError     = 0x00000010
	mbIconInfo      = 0x00000040
	mbSetForeground = 0x00010000
)

var (
	user32           = windows.NewLazySystemDLL("user32.dll")
	kernel32         = windows.NewLazySystemDLL("kernel32.dll")
	messageBoxW      = user32.NewProc("MessageBoxW")
	getConsoleWindow = kernel32.NewProc("GetConsoleWindow")
)

func showErrorDialog(title, message string) {
	showDialog(title, message, mbOK|mbIconError|mbSetForeground)
}

func showInfoDialog(title, message string) {
	showDialog(title, message, mbOK|mbIconInfo|mbSetForeground)
}

func showDialog(title, message string, flags uintptr) {
	if window, _, _ := getConsoleWindow.Call(); window != 0 {
		return
	}
	titlePtr, err := windows.UTF16PtrFromString(title)
	if err != nil {
		return
	}
	messagePtr, err := windows.UTF16PtrFromString(message)
	if err != nil {
		return
	}
	_, _, _ = messageBoxW.Call(
		0,
		uintptr(unsafe.Pointer(messagePtr)),
		uintptr(unsafe.Pointer(titlePtr)),
		flags,
	)
}
