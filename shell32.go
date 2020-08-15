// +build windows

package trash

import (
	"syscall"
	"unsafe"

	"github.com/lxn/win"
	"golang.org/x/sys/windows"
)

const (
	FO_MOVE   = 1
	FO_COPY   = 2
	FO_DELETE = 3
	FO_RENAME = 4
)

const (
	FOF_MULTIDESTFILES = 1
	FOF_SILENT         = 4
	FOF_NOCONFIRMATION = 16
	FOF_ALLOWUNDO      = 64
	FOF_NOERRORUI      = 1024
)

const (
	DE_SAMEFILE         = 0x71 // The source and destination files are the same file.
	DE_MANYSRC1DEST     = 0x72 // Multiple file paths were specified in the source buffer, but only one destination file path.
	DE_DIFFDIR          = 0x73 // Rename operation was specified but the destination path is a different directory. Use the move operation instead.
	DE_ROOTDIR          = 0x74 // The source is a root directory, which cannot be moved or renamed.
	DE_OPCANCELLED      = 0x75 // The operation was canceled by the user, or silently canceled if the appropriate flags were supplied to SHFileOperation.
	DE_DESTSUBTREE      = 0x76 // The destination is a subtree of the source.
	DE_ACCESSDENIEDSRC  = 0x78 // Security settings denied access to the source.
	DE_PATHTOODEEP      = 0x79 // The source or destination path exceeded or would exceed MAX_PATH.
	DE_MANYDEST         = 0x7A // The operation involved multiple destination paths, which can fail in the case of a move operation.
	DE_INVALIDFILES     = 0x7C // The path in the source or destination or both was invalid.
	DE_DESTSAMETREE     = 0x7D // The source and destination have the same parent folder.
	DE_FLDDESTISFILE    = 0x7E // The destination path is an existing file.
	DE_FILEDESTISFLD    = 0x80 // The destination path is an existing folder.
	DE_FILENAMETOOLONG  = 0x81 // The name of the file exceeds MAX_PATH.
	DE_DEST_IS_CDROM    = 0x82 // The destination is a read-only CD-ROM, possibly unformatted.
	DE_DEST_IS_DVD      = 0x83 // The destination is a read-only DVD, possibly unformatted.
	DE_DEST_IS_CDRECORD = 0x84 // The destination is a writable CD-ROM, possibly unformatted.
	DE_FILE_TOO_LARGE   = 0x85 // The file involved in the operation is too large for the destination media or file system.
	DE_SRC_IS_CDROM     = 0x86 // The source is a read-only CD-ROM, possibly unformatted.
	DE_SRC_IS_DVD       = 0x87 // The source is a read-only DVD, possibly unformatted.
	DE_SRC_IS_CDRECORD  = 0x88 // The source is a writable CD-ROM, possibly unformatted.
	DE_ERROR_MAX        = 0xB7 // MAX_PATH was exceeded during the operation.
)

type SHFILEOPSTRUCT struct {
	Hwnd                  win.HWND
	WFunc                 uint32
	PFrom                 *uint16
	PTo                   *uint16
	FFlags                uint32
	FAnyOperationsAborted win.BOOL
	HNameMappings         uint32
	LpszProgressTitle     *uint16
}

var (
	// Library
	libshell32 *windows.LazyDLL

	// Functions
	shFileOperation *windows.LazyProc
)

func init() {
	// Library
	libshell32 = windows.NewLazySystemDLL("shell32.dll")

	shFileOperation = libshell32.NewProc("SHFileOperationW")
}

func SHFileOperation(lpFileOp *SHFILEOPSTRUCT) win.HRESULT {
	ret, _, _ := syscall.Syscall(shFileOperation.Addr(), 1,
		uintptr(unsafe.Pointer(lpFileOp)),
		0,
		0)

	return win.HRESULT(ret)
}
