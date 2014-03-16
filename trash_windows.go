// +build windows

package trash

/*
#include "recycle.h"
*/
import "C"

import (
	"errors"
)

func IsAvailable() bool {
	return true
}

func MoveToTrash(filePath string) (string, error) {
	files := []string{filePath}
	C_files := C.makeCharArray(C.int(len(files)))
	defer C.freeCharArray(C_files, C.int(len(files)))
	for i, s := range files {
		C.setArrayString(C_files, C.CString(s), C.int(i))
	}

	success := C.RecycleFiles(C_files, C.int(len(files)), C.int(0))
	if success != 1 {
		return "", errors.New("file could not be recycled")
	}
	return "", nil
}
