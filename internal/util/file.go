package util

import "os"

func IsFileExist(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		return false
	}
	return true
}
