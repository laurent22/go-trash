// +build windows

package trash

import (
	"errors"
	"strconv"
	"syscall"

	"github.com/lxn/win"
)

// Tells whether it is possible to move a file to the trash
func IsAvailable() bool {
	return true
}

func getShortPathName(path string) (string, error) {
	p, err := syscall.UTF16FromString(path)
	if err != nil {
		return "", err
	}
	b := p // GetShortPathName says we can reuse buffer
	n := uint32(len(b))
	for {
		n, err = syscall.GetShortPathName(&p[0], &b[0], uint32(len(b)))
		if err != nil {
			return "", err
		}
		if n <= uint32(len(b)) {
			return syscall.UTF16ToString(b[:n]), nil
		}
		b = make([]uint16, n)
	}
}

func MoveToTrash(filePath string) (string, error) {

	filePath, err := getShortPathName(filePath)
	if err != nil {
		return "", err
	}

	fileop := win.SHFILEOPSTRUCT{
		Hwnd:                  win.HWND(0),
		WFunc:                 win.FO_DELETE,
		PFrom:                 syscall.StringToUTF16Ptr(filePath),
		PTo:                   nil,
		FFlags:                win.FOF_ALLOWUNDO | win.FOF_NOCONFIRMATION | win.FOF_NOERRORUI | win.FOF_SILENT,
		FAnyOperationsAborted: win.BOOL(0),
		HNameMappings:         0,
		LpszProgressTitle:     syscall.StringToUTF16Ptr(""),
	}

	result := win.SHFileOperation(&fileop)
	if result != 0 {
		return "", errors.New("File operation returned code " + strconv.Itoa(int(result)))
	}

	return "", nil
}
