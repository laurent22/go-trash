// +build windows

package trash

import (
	"errors"
	"strconv"
	"unicode/utf16"

	"github.com/lxn/win"
)

// Tells whether it is possible to move a file to the trash
func IsAvailable() bool {
	return true
}

func DoubleNullTerminatedUTF16PtrFromString(s string) *uint16 {
	return &(utf16.Encode([]rune(s + "\x00\x00"))[0])
}

func MoveToTrash(filePath string) (string, error) {
	if result := win.SHFileOperation(&win.SHFILEOPSTRUCT{
		Hwnd:                  win.HWND(0),
		WFunc:                 win.FO_DELETE,
		PFrom:                 DoubleNullTerminatedUTF16PtrFromString(filePath),
		PTo:                   nil,
		FFlags:                win.FOF_ALLOWUNDO | win.FOF_NOCONFIRMATION | win.FOF_NOERRORUI | win.FOF_SILENT,
		FAnyOperationsAborted: win.BOOL(0),
		HNameMappings:         0,
		LpszProgressTitle:     DoubleNullTerminatedUTF16PtrFromString(""), // Note: double-null termination not required
	}); result != 0 {
		return "", errors.New("File operation returned code " + strconv.Itoa(int(result)))
	}

	return "", nil
}
