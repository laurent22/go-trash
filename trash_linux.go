// +build linux

package trash

import (
	"os"
)

func IsAvailable() bool {
	return false
}

func MoveToTrash(filePath string) (string, error) {
	os.Remove(filePath)
	return "", nil
}
