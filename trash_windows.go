// +build windows

package trash

import (
	"errors"
	"strconv"
	"unicode/utf16"
)

// Tells whether it is possible to move a file to the trash
func IsAvailable() bool {
	return true
}

func DoubleNullTerminatedUTF16PtrFromString(s string) *uint16 {
	return &(utf16.Encode([]rune(s + "\x00\x00"))[0])
}

func MoveToTrash(filePath string) (string, error) {
	if result := SHFileOperation(&SHFILEOPSTRUCT{
		Hwnd:                  HWND(0),
		WFunc:                 FO_DELETE,
		PFrom:                 DoubleNullTerminatedUTF16PtrFromString(filePath),
		PTo:                   nil,
		FFlags:                FOF_ALLOWUNDO | FOF_NOCONFIRMATION | FOF_NOERRORUI | FOF_SILENT,
		FAnyOperationsAborted: BOOL(0),
		HNameMappings:         0,
		LpszProgressTitle:     DoubleNullTerminatedUTF16PtrFromString(""), // Note: double-null termination not required
	}); result != 0 {
		return "", errors.New("File operation returned code " + strconv.Itoa(int(result)))
	}

	return "", nil
}
