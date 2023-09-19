package main

//Ok so uhhh this is progress as I need to append a null terminator so the text can be fully outputted instead of the first letter of the string.

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

func main() {
	var user32 = windows.NewLazyDLL("user32.dll")
	var messageBox = user32.NewProc("MessageBoxA")

	textPtr, _ := windows.UTF16PtrFromString("Hello")
	captionPtr, _ := windows.UTF16PtrFromString("Hello")

	result, _, _ := messageBox.Call(
        uintptr(0),
        uintptr(unsafe.Pointer(textPtr)),
        uintptr(unsafe.Pointer(captionPtr)),
        uintptr(0),
    )

    if result == 0 {
        // Handle error if necessary
    }
}
