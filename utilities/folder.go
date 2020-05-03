package utilities

import (
	"log"
	"syscall"
)

func HideFolderForWindows(path string) {
	pathW, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		log.Fatal(err)
	}
	err = syscall.SetFileAttributes(pathW, syscall.FILE_ATTRIBUTE_HIDDEN)
	if err != nil {
		log.Fatal(err)
	}
}
